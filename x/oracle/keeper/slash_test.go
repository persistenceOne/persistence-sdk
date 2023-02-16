package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/testutil"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

// TestSlashAndResetMissCounters is a test for the slashing mechanism
func (s *KeeperTestSuite) TestSlashAndResetMissCounters() {
	app, ctx := s.app, s.ctx
	valAddr := s.valAddresses[0]

	// override the params with values that are easy for testing
	params := types.DefaultParams()
	params.VotePeriod = 10                                  // 10 blocks
	params.SlashFraction = sdk.NewDecWithPrec(5, 1)         // 50%
	params.SlashWindow = 100                                // 100 blocks (10 vote periods)
	params.MinValidPerWindow = sdk.MustNewDecFromStr("0.5") // 50%
	app.OracleKeeper.SetParams(ctx, params)

	// missCounter is a special value that:
	// 	missCounter / votePeriodsPerWindow < minValidPerWindow
	missCounter := uint64(6)

	app.OracleKeeper.SetMissCounter(ctx, valAddr, missCounter)
	app.OracleKeeper.SlashAndResetMissCounters(ctx)

	// ensure slashing effects are applied
	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	validator, found := app.StakingKeeper.GetValidator(ctx, valAddr)
	s.Require().True(found, "ensure that validator is found")
	s.Require().Equal(stakingtypes.Unbonding, validator.Status, "ensure its status is unbonding")
	s.Require().Equal(testutil.ValidatorAmountBonded.QuoRaw(2), validator.Tokens, "ensure its tokens are slashed")

	missCounter = app.OracleKeeper.GetMissCounter(ctx, valAddr)
	s.Require().Zero(missCounter, "ensure miss counter must be zero now")
}

func (s *KeeperTestSuite) TestSlashAndResetMissCounters2() {
	// initial setup
	addr, addr2 := valAddr, valAddr2
	amt := sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)

	s.Require().Equal(amt, s.app.StakingKeeper.Validator(s.ctx, addr).GetBondedTokens())
	s.Require().Equal(amt, s.app.StakingKeeper.Validator(s.ctx, addr2).GetBondedTokens())

	votePeriodsPerWindow := sdk.NewDec(int64(s.app.OracleKeeper.SlashWindow(s.ctx))).QuoInt64(int64(s.app.OracleKeeper.VotePeriod(s.ctx))).TruncateInt64()
	slashFraction := s.app.OracleKeeper.SlashFraction(s.ctx)
	minValidVotes := s.app.OracleKeeper.MinValidPerWindow(s.ctx).MulInt64(votePeriodsPerWindow).TruncateInt64()
	// Case 1, no slash
	s.app.OracleKeeper.SetMissCounter(s.ctx, valAddr, uint64(votePeriodsPerWindow-minValidVotes))
	s.app.OracleKeeper.SlashAndResetMissCounters(s.ctx)
	staking.EndBlocker(s.ctx, s.app.StakingKeeper)

	validator, _ := s.app.StakingKeeper.GetValidator(s.ctx, valAddr)
	s.Require().Equal(amt, validator.GetBondedTokens())

	// Case 2, slash
	s.app.OracleKeeper.SetMissCounter(s.ctx, valAddr, uint64(votePeriodsPerWindow-minValidVotes+1))
	s.app.OracleKeeper.SlashAndResetMissCounters(s.ctx)
	validator, _ = s.app.StakingKeeper.GetValidator(s.ctx, valAddr)
	s.Require().Equal(amt.Sub(slashFraction.MulInt(amt).TruncateInt()), validator.GetBondedTokens())
	s.Require().True(validator.Jailed)

	// Case 3, slash unbonded validator
	validator, _ = s.app.StakingKeeper.GetValidator(s.ctx, valAddr)
	validator.Status = stakingtypes.Unbonded
	validator.Jailed = false
	validator.Tokens = amt
	s.app.StakingKeeper.SetValidator(s.ctx, validator)

	s.app.OracleKeeper.SetMissCounter(s.ctx, valAddr, uint64(votePeriodsPerWindow-minValidVotes+1))
	s.app.OracleKeeper.SlashAndResetMissCounters(s.ctx)
	validator, _ = s.app.StakingKeeper.GetValidator(s.ctx, valAddr)
	s.Require().Equal(amt, validator.Tokens)
	s.Require().False(validator.Jailed)

	// Case 4, slash jailed validator
	validator, _ = s.app.StakingKeeper.GetValidator(s.ctx, valAddr)
	validator.Status = stakingtypes.Bonded
	validator.Jailed = true
	validator.Tokens = amt
	s.app.StakingKeeper.SetValidator(s.ctx, validator)

	s.app.OracleKeeper.SetMissCounter(s.ctx, valAddr, uint64(votePeriodsPerWindow-minValidVotes+1))
	s.app.OracleKeeper.SlashAndResetMissCounters(s.ctx)
	validator, _ = s.app.StakingKeeper.GetValidator(s.ctx, valAddr)
	s.Require().Equal(amt, validator.Tokens)
}
