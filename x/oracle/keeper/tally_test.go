package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/keeper"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/testutil"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

// TestTally is a test for the tallying logic that calculates the median.
func (s *KeeperTestSuite) TestTally() {
	testCases := []tallyTestCase{
		{
			Description: "basic test for picking a winner",
			// Tally 2 votes
			// rewardBand: 1.000000000000000000
			// weightedMedian: 1.000000000000000000
			// standardDeviation: 0.707106781186547525
			// rewardSpread: 0.707106781186547525

			Ballot: types.ExchangeRateBallot{
				types.NewVoteForTally(sdk.MustNewDecFromStr("1.0"), "AAA", val(1), 10000),
				types.NewVoteForTally(sdk.MustNewDecFromStr("2.0"), "AAA", val(2), 10000),
			},
			RewardBand: sdk.MustNewDecFromStr("1.0"),
			ValidatorClaimMap: map[string]types.Claim{
				val(1).String(): types.NewClaim(10000, 0, 0, val(1)),
				val(2).String(): types.NewClaim(10000, 0, 0, val(2)),
			},
			ExpectedWeightedMedian: sdk.MustNewDecFromStr("1.0"),
			ExpectedValidatorClaimMap: map[string]types.Claim{
				val(1).String(): types.NewClaim(10000, 10000, 1, val(1)),
				val(2).String(): types.NewClaim(10000, 0, 0, val(2)),
			},
		},
		{
			Description: "basic test for picking all winners in reward spread",
			// Tally 2 votes
			// rewardBand: 1.000000000000000000
			// weightedMedian: 1.000000000000000000
			// standardDeviation: 0.353553390593273762
			// rewardSpread: 0.500000000000000000

			Ballot: types.ExchangeRateBallot{
				types.NewVoteForTally(sdk.MustNewDecFromStr("1.0"), "AAA", val(1), 10000),
				types.NewVoteForTally(sdk.MustNewDecFromStr("1.5"), "AAA", val(2), 10000),
			},
			RewardBand: sdk.MustNewDecFromStr("1.0"),
			ValidatorClaimMap: map[string]types.Claim{
				val(1).String(): types.NewClaim(10000, 0, 0, val(1)),
				val(2).String(): types.NewClaim(10000, 0, 0, val(2)),
			},
			ExpectedWeightedMedian: sdk.MustNewDecFromStr("1.0"),
			ExpectedValidatorClaimMap: map[string]types.Claim{
				val(1).String(): types.NewClaim(10000, 10000, 1, val(1)),
				val(2).String(): types.NewClaim(10000, 10000, 1, val(2)),
			},
		},
		{
			Description: "basic test for >50% attack",
			// Tally 5 votes
			// rewardBand: 1.000000000000000000
			// weightedMedian: 999.000000000000000000
			// standardDeviation: 892.638336617916046807
			// rewardSpread: 892.638336617916046807

			Ballot: types.ExchangeRateBallot{
				types.NewVoteForTally(sdk.MustNewDecFromStr("1.0"), "AAA", val(1), 10000),
				types.NewVoteForTally(sdk.MustNewDecFromStr("1.0"), "AAA", val(2), 10000),
				types.NewVoteForTally(sdk.MustNewDecFromStr("1.0"), "AAA", val(3), 10000),
				types.NewVoteForTally(sdk.MustNewDecFromStr("1.0"), "AAA", val(4), 10000),
				types.NewVoteForTally(sdk.MustNewDecFromStr("999.0"), "AAA", val(5), 40002),
			},
			RewardBand: sdk.MustNewDecFromStr("1.0"),
			ValidatorClaimMap: map[string]types.Claim{
				val(1).String(): types.NewClaim(10000, 0, 0, val(1)),
				val(2).String(): types.NewClaim(10000, 0, 0, val(2)),
				val(3).String(): types.NewClaim(10000, 0, 0, val(3)),
				val(4).String(): types.NewClaim(10000, 0, 0, val(4)),
				val(5).String(): types.NewClaim(40002, 0, 0, val(5)),
			},
			ExpectedWeightedMedian: sdk.MustNewDecFromStr("999.0"),
			ExpectedValidatorClaimMap: map[string]types.Claim{
				val(1).String(): types.NewClaim(10000, 0, 0, val(1)),
				val(2).String(): types.NewClaim(10000, 0, 0, val(2)),
				val(3).String(): types.NewClaim(10000, 0, 0, val(3)),
				val(4).String(): types.NewClaim(10000, 0, 0, val(4)),
				val(5).String(): types.NewClaim(40002, 40002, 1, val(5)),
			},
		},
	}

	for _, testCase := range testCases {
		s.T().Log("TestTally Case:", testCase.Description)

		resultValidatorClaimMap := copyValidatorClaimMap(testCase.ValidatorClaimMap)
		weightedMedian, err := keeper.Tally(
			testCase.Ballot,
			testCase.RewardBand,
			resultValidatorClaimMap,
		)

		s.Require().NoError(err)
		s.Require().EqualValues(testCase.ExpectedWeightedMedian, weightedMedian)
		s.Require().EqualValues(testCase.ExpectedValidatorClaimMap, resultValidatorClaimMap)
	}
}

type tallyTestCase struct {
	Description string

	Ballot            types.ExchangeRateBallot
	RewardBand        sdk.Dec
	ValidatorClaimMap map[string]types.Claim

	ExpectedValidatorClaimMap map[string]types.Claim
	ExpectedWeightedMedian    sdk.Dec
}

func copyValidatorClaimMap(m map[string]types.Claim) map[string]types.Claim {
	mm := make(map[string]types.Claim, len(m))
	for k, v := range m {
		mm[k] = v
	}

	return mm
}

func val(n int) sdk.ValAddress {
	return sdk.ValAddress(fmt.Sprintf("val%08d__________", n))
}

// TestBuildClaimsMapAndTally is a test for collection clams map and tallying.
func (s *KeeperTestSuite) TestBuildClaimsMapAndTally() {
	// custom app and context for this test
	app, ctx := s.initAppAndContext()

	// generate 100 equal validators in consensus
	_, valAddresses, err := testutil.StakingAddValidators(
		app.BankKeeper,
		app.StakingKeeper,
		ctx,
		100,
	)
	s.Require().NoError(err)
	s.Require().Len(valAddresses, 100)
	s.Require().Equal(100, countActiveValidators(ctx, app.StakingKeeper))

	// override the params with values that are easy for testing
	params := types.DefaultParams()
	params.VotePeriod = 1                            // 10 block
	params.VoteThreshold = sdk.NewDecWithPrec(50, 2) // 50%
	params.RewardBand = sdk.NewDecWithPrec(50, 2)    // 50%
	app.OracleKeeper.SetParams(ctx, params)

	// initial exchange rate
	app.OracleKeeper.SetExchangeRate(ctx, "ATOM", sdk.MustNewDecFromStr("1.0"))

	{
		s.T().Log("TestBuildClaimsMapAndTally: 1 vote out of 100 counts (?)")
		app.OracleKeeper.SetAggregateExchangeRateVote(ctx, valAddresses[0], types.NewAggregateExchangeRateVote([]types.ExchangeRateTuple{{
			Denom: "ATOM", ExchangeRate: sdk.MustNewDecFromStr("999.0"),
		}}, valAddresses[0]))

		err = app.OracleKeeper.BuildClaimsMapAndTally(ctx, params)
		s.Require().NoError(err)

		finalRate, err := app.OracleKeeper.GetExchangeRate(ctx, "ATOM")
		s.Require().NoError(err)
		s.Require().EqualValues(sdk.MustNewDecFromStr("999.0"), finalRate)

		// rest of validators marked with misses
		for i := 1; i < len(valAddresses); i++ {
			s.Require().Equal(uint64(1), app.OracleKeeper.GetMissCounter(ctx, valAddresses[i]))
		}
	}
}

func countActiveValidators(ctx sdk.Context, k stakingkeeper.Keeper) int {
	uniqueValidators := make(map[string]struct{})

	maxValidators := k.MaxValidators(ctx)
	iterator := k.ValidatorsPowerStoreIterator(ctx)

	for ; iterator.Valid() && len(uniqueValidators) < int(maxValidators); iterator.Next() {
		validator := k.Validator(ctx, iterator.Value())

		if validator.IsBonded() {
			valOper := validator.GetOperator()
			uniqueValidators[valOper.String()] = struct{}{}
		}
	}

	iterator.Close()

	return len(uniqueValidators)
}
