package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

// Test the reward giving mechanism
func (s *IntegrationTestSuite) TestRewardBallotWinners() {
	// Add claim pools
	claims := []types.Claim{
		types.NewClaim(10, 10, 0, valAddr),
		types.NewClaim(20, 20, 0, valAddr2),
	}

	// Prepare reward pool
	givingAmt := sdk.NewCoins(sdk.NewInt64Coin(types.PersistenceDenom, 30000000))
	err := s.app.BankKeeper.MintCoins(s.ctx, minttypes.ModuleName, givingAmt)
	s.Require().NoError(err)
	err = s.app.BankKeeper.SendCoinsFromModuleToModule(s.ctx,  minttypes.ModuleName, types.ModuleName, givingAmt)
	s.Require().NoError(err)

	var voteTargets []string
	params := s.app.OracleKeeper.GetParams(s.ctx)
	for _, v := range params.AcceptList {
		voteTargets = append(voteTargets, v.SymbolDenom)
	}

	votePeriodsPerWindow := sdk.NewDec((int64)(s.app.OracleKeeper.RewardDistributionWindow(s.ctx))).
		QuoInt64((int64)(s.app.OracleKeeper.VotePeriod(s.ctx))).
		TruncateInt64()
	s.app.OracleKeeper.RewardBallotWinners(s.ctx, (int64)(s.app.OracleKeeper.VotePeriod(s.ctx)), (int64)(s.app.OracleKeeper.RewardDistributionWindow(s.ctx)), voteTargets, claims)
	outstandingRewardsDec := s.app.DistrKeeper.GetValidatorOutstandingRewardsCoins(s.ctx, valAddr)
	outstandingRewards, _ := outstandingRewardsDec.TruncateDecimal()
	s.Require().Equal(sdk.NewDecFromInt(givingAmt.AmountOf(types.PersistenceDenom)).QuoInt64(votePeriodsPerWindow).QuoInt64(3).TruncateInt(),
		outstandingRewards.AmountOf(types.PersistenceDenom))
}

func (s *IntegrationTestSuite) TestRewardBallotWinnersZeroPower() {
	zeroClaim := types.NewClaim(0, 0, 0, valAddr)
	s.app.OracleKeeper.RewardBallotWinners(s.ctx, 0, 0, []string{}, []types.Claim{zeroClaim})
	outstandingRewardsDec := s.app.DistrKeeper.GetValidatorOutstandingRewardsCoins(s.ctx, valAddr)
	s.Require().Equal("", outstandingRewardsDec.String())
}
