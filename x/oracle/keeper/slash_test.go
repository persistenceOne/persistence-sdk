package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkstaking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/testutil"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

// TestSlashAndResetMissCounters added more test cases for slashing, incl. unbonded validators.
func (s *KeeperTestSuite) TestSlashAndResetMissCounters() {
	app, ctx := s.app, s.ctx
	valAddr := s.valAddresses[0]

	amt := testutil.ValidatorAmountBonded
	s.Require().Equal(amt, s.app.StakingKeeper.Validator(s.ctx, valAddr).GetBondedTokens())

	// override the params with values that are easy for testing
	params := types.DefaultParams()
	params.VotePeriod = 10                                  // 10 blocks
	params.SlashFraction = sdk.NewDecWithPrec(5, 1)         // 50%
	params.SlashWindow = 100                                // 100 blocks (10 vote periods)
	params.MinValidPerWindow = sdk.MustNewDecFromStr("0.5") // 50%
	app.OracleKeeper.SetParams(ctx, params)

	// missCounterSlash is a special value that:
	// 	missCounterSlash / votePeriodsPerWindow < minValidPerWindow
	missCounterSlash := uint64(6)

	// missCounterNoSlash is a special value that:
	// 	missCounterNoSlash / votePeriodsPerWindow >= minValidPerWindow
	missCounterNoSlash := uint64(5)

	// Case 1, no slash
	s.app.OracleKeeper.SetMissCounter(s.ctx, valAddr, missCounterNoSlash)
	s.app.OracleKeeper.SlashAndResetMissCounters(s.ctx)
	staking.EndBlocker(s.ctx, s.app.StakingKeeper)

	validator, _ := s.app.StakingKeeper.GetLiquidValidator(s.ctx, valAddr)
	s.Require().Equal(amt, validator.GetBondedTokens())

	// Case 2, slash
	s.app.OracleKeeper.SetMissCounter(s.ctx, valAddr, missCounterSlash)
	s.app.OracleKeeper.SlashAndResetMissCounters(s.ctx)

	validator, _ = s.app.StakingKeeper.GetLiquidValidator(s.ctx, valAddr)
	s.Require().Equal(amt.Sub(params.SlashFraction.MulInt(amt).TruncateInt()), validator.GetBondedTokens())
	s.Require().True(validator.Jailed)

	missCounter := app.OracleKeeper.GetMissCounter(ctx, valAddr)
	s.Require().Zero(missCounter)

	// Case 3, slash unbonded validator
	validator, _ = s.app.StakingKeeper.GetLiquidValidator(s.ctx, valAddr)
	validator.Status = sdkstaking.Unbonded
	validator.Jailed = false
	validator.Tokens = amt
	s.app.StakingKeeper.SetValidator(s.ctx, validator)

	s.app.OracleKeeper.SetMissCounter(s.ctx, valAddr, missCounterSlash)
	s.app.OracleKeeper.SlashAndResetMissCounters(s.ctx)
	validator, _ = s.app.StakingKeeper.GetLiquidValidator(s.ctx, valAddr)
	s.Require().Equal(amt, validator.Tokens)
	s.Require().False(validator.Jailed)

	// Case 4, slash jailed validator
	validator, _ = s.app.StakingKeeper.GetLiquidValidator(s.ctx, valAddr)
	validator.Status = sdkstaking.Bonded
	validator.Jailed = true
	validator.Tokens = amt
	s.app.StakingKeeper.SetValidator(s.ctx, validator)

	s.app.OracleKeeper.SetMissCounter(s.ctx, valAddr, missCounterSlash)
	s.app.OracleKeeper.SlashAndResetMissCounters(s.ctx)
	validator, _ = s.app.StakingKeeper.GetLiquidValidator(s.ctx, valAddr)
	s.Require().Equal(amt, validator.Tokens)
}
