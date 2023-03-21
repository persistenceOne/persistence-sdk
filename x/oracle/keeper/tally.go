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

	powerReduction := k.StakingKeeper.PowerReduction(ctx)

	// Calculate total validator power
	var totalBondedValidatorPower int64

	for _, v := range k.StakingKeeper.GetBondedValidatorsByPower(ctx) {
		addr := v.GetOperator()
		valConsensusPower := v.GetConsensusPower(powerReduction)
		totalBondedValidatorPower += valConsensusPower
		validatorClaimMap[addr.String()] = types.NewClaim(valConsensusPower, 0, 0, addr)
	}

	var (
		// voteTargets defines the symbol (ticker) denoms that we require votes on
		voteTargets      = make([]string, 0, len(params.AcceptList))
		voteTargetDenoms = make([]string, 0, len(params.AcceptList))
	)

	for _, v := range params.AcceptList {
		voteTargets = append(voteTargets, v.SymbolDenom)
		voteTargetDenoms = append(voteTargetDenoms, v.SymbolDenom)
	}

	// Clear all exchange rates
	k.ClearExchangeRates(ctx)

	// Organize votes to ballot by denom
	// NOTE: **Filter out inactive or jailed validators**
	ballotDenomSlice := k.OrganizeBallotByDenom(ctx, validatorClaimMap)

	threshold := k.GetVoteThreshold(ctx).MulInt64(types.MaxVoteThresholdMultiplier).TruncateInt64()

	// Iterate through ballots and update exchange rates; drop if not enough votes have been achieved.
	for _, ballotDenom := range ballotDenomSlice {
		// Calculate the portion of votes received as an integer, scaled up using the
		// same multiplier as the `threshold` computed above
		support := ballotDenom.Ballot.Power() * types.MaxVoteThresholdMultiplier / totalBondedValidatorPower
		if support < threshold {
			ctx.Logger().Info("Ballot voting power is under vote threshold, dropping ballot", "denom", ballotDenom)
			continue
		}

		// Get weighted median of exchange rates
		exchangeRate, err := Tally(ballotDenom.Ballot, params.RewardBand, validatorClaimMap)
		if err != nil {
			return err
		}

		// Set the exchange rate, emit ABCI event
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
	k.ClearVotes(ctx, params.VotePeriod)

	return nil
}

// Tally calculates the median and returns it. It sets the set of voters to be
// rewarded, i.e. voted within a reasonable spread from the weighted median to
// the store. Note, the ballot is sorted by ExchangeRate.
// https://classic-docs.terra.money/docs/develop/module-specifications/spec-oracle.html#tally
func Tally(
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
