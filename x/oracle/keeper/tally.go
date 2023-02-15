package keeper

import (
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BuildClaimsMapAndTally builds a claim map over all validators in active set,
// marks misses and handles ballot tallying.
func (k Keeper) BuildClaimsMapAndTally(ctx sdk.Context, params types.Params) error {
	// Build claim map over all validators in active set
	validatorClaimMap := make(map[string]types.Claim)

	maxValidators := k.StakingKeeper.MaxValidators(ctx)
	iterator := k.StakingKeeper.ValidatorsPowerStoreIterator(ctx)

	defer iterator.Close()

	powerReduction := k.StakingKeeper.PowerReduction(ctx)

	addedValidators := 0
	for ; iterator.Valid() && addedValidators < int(maxValidators); iterator.Next() {
		validator := k.StakingKeeper.Validator(ctx, iterator.Value())

		// Exclude not bonded validator
		if validator.IsBonded() {
			valAddr := validator.GetOperator()
			validatorClaimMap[valAddr.String()] = types.Claim{
				Power:     validator.GetConsensusPower(powerReduction),
				Weight:    0,
				WinCount:  0,
				Recipient: valAddr,
			}

			addedValidators++
		}
	}

	var (
		// voteTargets defines the symbol (ticker) denoms that we require votes on
		voteTargets      []string
		voteTargetDenoms []string
	)

	for _, v := range params.AcceptList {
		voteTargets = append(voteTargets, v.SymbolDenom)
		voteTargetDenoms = append(voteTargetDenoms, v.BaseDenom)
	}

	// Clear all exchange rates
	k.IterateExchangeRates(ctx, func(denom string, _ sdk.Dec) (stop bool) {
		k.DeleteExchangeRate(ctx, denom)
		return false
	})

	// Organize votes to ballot by denom
	// NOTE: **Filter out inactive or jailed validators**
	voteMap := k.OrganizeBallotByDenom(ctx, validatorClaimMap)

	ballotDenomSlice := types.BallotMapToSlice(voteMap)

	// Iterate through ballots and update exchange rates; drop if not enough votes have been achieved.
	for _, ballotDenom := range ballotDenomSlice {
		// Get weighted median of exchange rates
		exchangeRate, err := Tally(ctx, ballotDenom.Ballot, params.RewardBand, validatorClaimMap)
		if err != nil {
			return err
		}

		// Set the exchange rate, emit event
		k.SetExchangeRateWithEvent(ctx, ballotDenom.Denom, exchangeRate)
	}

	// update miss counting & slashing
	voteTargetsLen := len(voteTargets)

	claimSlice := types.ClaimMapToSlice(validatorClaimMap)
	for _, claim := range claimSlice {
		// Skip valid voters
		if int(claim.WinCount) == voteTargetsLen {
			continue
		}

		// Increase miss counter
		k.SetMissCounter(ctx, claim.Recipient, k.GetMissCounter(ctx, claim.Recipient)+1)
	}

	// Distribute rewards to ballot winners
	k.RewardBallotWinners(
		ctx,
		int64(params.VotePeriod),
		int64(params.RewardDistributionWindow),
		voteTargetDenoms,
		claimSlice,
	)

	// Clear the ballot
	k.ClearBallots(ctx, params.VotePeriod)

	return nil
}

// Tally calculates the median and returns it. It sets the set of voters to be
// rewarded, i.e. voted within a reasonable spread from the weighted median to
// the store. Note, the ballot is sorted by ExchangeRate.
func Tally(
	ctx sdk.Context,
	ballot types.ExchangeRateBallot,
	rewardBand sdk.Dec,
	validatorClaimMap map[string]types.Claim,
) (sdk.Dec, error) {
	weightedMedian, err := ballot.WeightedMedian()

	if err != nil {
		return sdk.ZeroDec(), err
	}

	standardDeviation, err := ballot.StandardDeviation()
	if err != nil {
		return sdk.ZeroDec(), err
	}

	// rewardSpread is the MAX((weightedMedian * (rewardBand/2)), standardDeviation)
	rewardSpread := weightedMedian.Mul(rewardBand.QuoInt64(2))
	rewardSpread = sdk.MaxDec(rewardSpread, standardDeviation)

	for _, tallyVote := range ballot {
		// Filter ballot winners. For voters, we filter out the tally vote iff:
		// (weightedMedian - rewardSpread) <= ExchangeRate <= (weightedMedian + rewardSpread)
		if (tallyVote.ExchangeRate.GTE(weightedMedian.Sub(rewardSpread)) &&
			tallyVote.ExchangeRate.LTE(weightedMedian.Add(rewardSpread))) ||
			!tallyVote.ExchangeRate.IsPositive() {
			key := tallyVote.Voter.String()
			claim := validatorClaimMap[key]

			claim.Weight += tallyVote.Power
			claim.WinCount++
			validatorClaimMap[key] = claim
		}
	}

	return weightedMedian, nil
}
