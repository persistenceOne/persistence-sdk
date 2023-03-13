package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

func (s *KeeperTestSuite) TestVoteThreshold() {
	app, ctx := s.app, s.ctx

	voteDec := app.OracleKeeper.GetVoteThreshold(ctx)
	s.Require().Equal(sdk.MustNewDecFromStr("0.5"), voteDec)

	newVoteTreshold := sdk.MustNewDecFromStr("0.6")
	defaultParams := types.DefaultParams()
	defaultParams.VoteThreshold = newVoteTreshold
	app.OracleKeeper.SetParams(ctx, defaultParams)

	voteThresholdDec := app.OracleKeeper.GetVoteThreshold(ctx)
	s.Require().Equal(newVoteTreshold, voteThresholdDec)
}
