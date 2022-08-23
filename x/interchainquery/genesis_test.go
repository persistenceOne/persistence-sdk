package interchainquery_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/stretchr/testify/suite"

	"github.com/persistenceOne/persistence-sdk/simapp"
	"github.com/persistenceOne/persistence-sdk/x/interchainquery"
	"github.com/persistenceOne/persistence-sdk/x/interchainquery/keeper"
	"github.com/persistenceOne/persistence-sdk/x/interchainquery/types"
)

func init() {
	ibctesting.DefaultTestingAppInit = simapp.SetupTestingApp
}

func TestInterChainQueryTestSuite(t *testing.T) {
	suite.Run(t, new(InterChainQueryTestSuite))
}

type InterChainQueryTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
	path   *ibctesting.Path
}

func (s *InterChainQueryTestSuite) GetSimApp(chain *ibctesting.TestChain) *simapp.SimApp {
	app, ok := chain.App.(*simapp.SimApp)
	if !ok {
		panic("not sim app")
	}

	return app
}

func (s *InterChainQueryTestSuite) SetupTest() {
	s.coordinator = ibctesting.NewCoordinator(s.T(), 2)
	s.chainA = s.coordinator.GetChain(ibctesting.GetChainID(1))
	s.chainB = s.coordinator.GetChain(ibctesting.GetChainID(2))

	s.path = newSimAppPath(s.chainA, s.chainB)
	s.coordinator.SetupConnections(s.path)
}

func newSimAppPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort

	return path
}

func (s *InterChainQueryTestSuite) TestInitGenesis() {
	bondedQuery := stakingtypes.QueryValidatorsRequest{Status: stakingtypes.BondStatusBonded}
	bz1, err := bondedQuery.Marshal()
	s.NoError(err)

	query := s.GetSimApp(s.chainA).InterchainQueryKeeper.NewQuery(
		s.chainA.GetContext(),
		"",
		s.path.EndpointB.ConnectionID,
		s.chainB.ChainID,
		"cosmos.staking.v1beta1.Query/Validators",
		bz1,
		sdk.NewInt(200),
		"",
		0,
	)

	interchainquery.InitGenesis(s.chainA.GetContext(), s.GetSimApp(s.chainA).InterchainQueryKeeper, types.GenesisState{Queries: []types.Query{*query}})

	id := keeper.GenerateQueryHash(s.path.EndpointB.ConnectionID, s.chainB.ChainID, "cosmos.staking.v1beta1.Query/Validators", bz1, "")
	queryResponse, found := s.GetSimApp(s.chainA).InterchainQueryKeeper.GetQuery(s.chainA.GetContext(), id)
	s.True(found)
	s.Equal(s.path.EndpointB.ConnectionID, queryResponse.ConnectionId)
	s.Equal(s.chainB.ChainID, queryResponse.ChainId)
	s.Equal("cosmos.staking.v1beta1.Query/Validators", queryResponse.QueryType)
	s.Equal(sdk.NewInt(200), queryResponse.Period)
	s.Equal(uint64(0), queryResponse.Ttl)
	s.Equal("", queryResponse.CallbackId)
}
