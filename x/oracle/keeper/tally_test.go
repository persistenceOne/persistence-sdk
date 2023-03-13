package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

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

// TestBuildClaimsMapAndTallyBelowThreshold is a test for collection claims map and tallying for
// validators votes amount below VoteThreshold.
func (s *KeeperTestSuite) TestBuildClaimsMapAndTallyBelowThreshold() {
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
	s.Require().Equal(100, len(app.StakingKeeper.GetBondedValidatorsByPower(ctx)))

	// override the params with values that are easy for testing
	params := types.DefaultParams()
	params.VotePeriod = 1                            // 10 block
	params.VoteThreshold = sdk.NewDecWithPrec(50, 2) // 50%
	params.RewardBand = sdk.NewDecWithPrec(50, 2)    // 50%
	params.AcceptList = types.DenomList{{
		BaseDenom:   types.AtomSymbol,
		SymbolDenom: types.AtomSymbol,
		Exponent:    6,
	}}
	app.OracleKeeper.SetParams(ctx, params)

	// initial exchange rate
	app.OracleKeeper.SetExchangeRate(ctx, types.AtomSymbol, sdk.MustNewDecFromStr("1.0"))

	s.T().Log("TestBuildClaimsMapAndTally: 1 vote out of 100 doesn't count (below the threshold of 50%)")

	app.OracleKeeper.SetAggregateExchangeRateVote(ctx, valAddresses[0], types.NewAggregateExchangeRateVote([]types.ExchangeRateTuple{{
		Denom: types.AtomSymbol, ExchangeRate: sdk.MustNewDecFromStr("999.0"),
	}}, valAddresses[0]))

	err = app.OracleKeeper.BuildClaimsMapAndTally(ctx, params)
	s.Require().NoError(err)

	// we haven't reached the vote threshold yet, so the exchange rate not updated and is reset
	_, err = app.OracleKeeper.GetExchangeRate(ctx, types.AtomSymbol)
	s.Require().ErrorContains(err, types.ErrUnknownDenom.Error())

	// rest of validators marked with misses
	for valN := 1; valN < len(valAddresses); valN++ {
		s.Require().Equal(
			uint64(1),
			app.OracleKeeper.GetMissCounter(ctx, valAddresses[valN]),
			fmt.Sprintf("validator: %d", valN),
		)
	}
}

// TestBuildClaimsMapAndTallyAboveThreshold is a test for collection claims map and tallying for
// validators votes amount up to or above VoteThreshold.
func (s *KeeperTestSuite) TestBuildClaimsMapAndTallyAboveThreshold() {
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
	s.Require().Equal(100, len(app.StakingKeeper.GetBondedValidatorsByPower(ctx)))

	// override the params with values that are easy for testing
	params := types.DefaultParams()
	params.VotePeriod = 1                            // 10 block
	params.VoteThreshold = sdk.NewDecWithPrec(50, 2) // 50%
	params.RewardBand = sdk.NewDecWithPrec(50, 2)    // 50%
	params.AcceptList = types.DenomList{{
		BaseDenom:   types.AtomDenom,
		SymbolDenom: types.AtomSymbol,
		Exponent:    6,
	}}
	app.OracleKeeper.SetParams(ctx, params)

	// initial exchange rate
	app.OracleKeeper.SetExchangeRate(ctx, types.AtomSymbol, sdk.MustNewDecFromStr("1.0"))

	s.T().Log("TestBuildClaimsMapAndTally: >=50 votes out of 100 (above the threshold of 50%)")

	const halfOfBondedValidatorsCount = 50

	for valN := 0; valN < halfOfBondedValidatorsCount; valN++ {
		app.OracleKeeper.SetAggregateExchangeRateVote(ctx, valAddresses[valN], types.NewAggregateExchangeRateVote([]types.ExchangeRateTuple{{
			Denom: types.AtomSymbol, ExchangeRate: sdk.MustNewDecFromStr("999.0"),
		}}, valAddresses[valN]))
	}

	err = app.OracleKeeper.BuildClaimsMapAndTally(ctx, params)
	s.Require().NoError(err)

	// have reached the vote threshold yet, so the exchange rate is updated
	newRate, err := app.OracleKeeper.GetExchangeRate(ctx, types.AtomSymbol)
	s.Require().NoError(err)
	s.Require().EqualValues(sdk.MustNewDecFromStr("999.0"), newRate)

	// first half of validators doesn't have a miss
	for valN := 0; valN < halfOfBondedValidatorsCount; valN++ {
		s.Require().Zero(
			app.OracleKeeper.GetMissCounter(ctx, valAddresses[valN]),
			fmt.Sprintf("validator: %d", valN),
		)
	}

	// rest of validators marked with misses
	for valN := halfOfBondedValidatorsCount; valN < len(valAddresses); valN++ {
		s.Require().Equal(
			uint64(1),
			app.OracleKeeper.GetMissCounter(ctx, valAddresses[valN]),
			fmt.Sprintf("validator: %d", valN),
		)
	}
}
