package cli_test

import (
	"context"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/persistenceOne/persistence-sdk/v2/simapp"
	"github.com/persistenceOne/persistence-sdk/v2/x/epochs/types"
)

type QueryTestSuite struct {
	suite.Suite

	app *simapp.SimApp
	ctx sdk.Context

	queryHelper *baseapp.QueryServiceTestHelper
	queryClient types.QueryClient
}

func (s *QueryTestSuite) SetupSuite() {
	s.app = simapp.Setup(s.T(), false)
	s.ctx = s.app.NewContext(false, tmproto.Header{})
	s.queryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.app.GRPCQueryRouter(),
		Ctx:             s.ctx,
	}
	s.queryClient = types.NewQueryClient(s.queryHelper)

	// add new epoch
	epoch := types.EpochInfo{
		Identifier:              "weekly",
		StartTime:               time.Time{},
		Duration:                time.Hour,
		CurrentEpoch:            0,
		CurrentEpochStartHeight: 0,
		CurrentEpochStartTime:   time.Time{},
		EpochCountingStarted:    false,
	}

	err := s.app.EpochsKeeper.AddEpochInfo(s.ctx, epoch)
	require.NoError(s.T(), err)

	s.app.Commit()
}

func (s *QueryTestSuite) TestQueriesNeverAlterState() {
	testCases := []struct {
		name   string
		query  string
		input  interface{}
		output interface{}
	}{
		{
			"Query current epoch",
			"/persistence.epochs.v1beta1.Query/CurrentEpoch",
			&types.QueryCurrentEpochRequest{Identifier: "weekly"},
			&types.QueryCurrentEpochResponse{},
		},
		{
			"Query epochs info",
			"/persistence.epochs.v1beta1.Query/EpochInfos",
			&types.QueryEpochsInfoRequest{},
			&types.QueryEpochsInfoResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.SetupSuite()
			err := s.queryHelper.Invoke(context.Background(), tc.query, tc.input, tc.output)
			s.Require().NoError(err)
		})
	}
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
