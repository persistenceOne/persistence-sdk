package cli_test

import (
	"context"
	"testing"
	"time"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	cdcutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	epochskeeper "github.com/persistenceOne/persistence-sdk/v6/x/epochs/keeper"
	"github.com/persistenceOne/persistence-sdk/v6/x/epochs/types"
)

type QueryTestSuite struct {
	suite.Suite

	ctx          sdk.Context
	epochsKeeper *epochskeeper.Keeper
	queryHelper  *baseapp.QueryServiceTestHelper
	queryClient  types.QueryClient
}

func (s *QueryTestSuite) SetupSuite() {
	epochsStoreKey := storetypes.NewKVStoreKey(types.StoreKey)
	s.ctx = testutil.DefaultContext(epochsStoreKey, storetypes.NewTransientStoreKey("transient_test"))
	s.epochsKeeper = epochskeeper.NewKeeper(epochsStoreKey)
	s.epochsKeeper = s.epochsKeeper.SetHooks(types.NewMultiEpochHooks())
	s.ctx = s.ctx.WithBlockHeight(1).WithChainID("persistence-1").WithBlockTime(time.Now().UTC())
	s.epochsKeeper.InitGenesis(s.ctx, *types.DefaultGenesis())

	queryRouter := baseapp.NewGRPCQueryRouter()
	cfg := module.NewConfigurator(nil, nil, queryRouter)
	types.RegisterQueryServer(cfg.QueryServer(), epochskeeper.NewQuerier(*s.epochsKeeper))
	s.queryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: queryRouter,
		Ctx:             s.ctx,
	}
	interfaceRegistry := cdcutil.CodecOptions{AccAddressPrefix: "persistence", ValAddressPrefix: "persistencevaloper"}.NewInterfaceRegistry()
	s.queryHelper.SetInterfaceRegistry(interfaceRegistry)
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

	err := s.epochsKeeper.AddEpochInfo(s.ctx, epoch)
	require.NoError(s.T(), err)
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
