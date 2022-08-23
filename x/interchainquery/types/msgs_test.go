package types_test

import (
	"testing"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/stretchr/testify/suite"

	"github.com/persistenceOne/persistence-sdk/simapp"
	"github.com/persistenceOne/persistence-sdk/x/interchainquery/keeper"
	"github.com/persistenceOne/persistence-sdk/x/interchainquery/types"
)

const TestOwnerAddress = "cosmos17dtl0mjt3t77kpuhg2edqzjpszulwhgzuj9ljs"

func init() {
	ibctesting.DefaultTestingAppInit = simapp.SetupTestingApp
}

func TestTypesTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}

type TypesTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
	path   *ibctesting.Path
}

func (s *TypesTestSuite) GetSimApp(chain *ibctesting.TestChain) *simapp.SimApp {
	app, ok := chain.App.(*simapp.SimApp)
	if !ok {
		panic("not sim app")
	}

	return app
}

func (s *TypesTestSuite) SetupTest() {
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

func (s *TypesTestSuite) TestMsgSubmitQueryResponse() {
	bondedQuery := stakingtypes.QueryValidatorsRequest{Status: stakingtypes.BondStatusBonded}
	bz1, err := bondedQuery.Marshal()
	s.NoError(err)

	qvr := stakingtypes.QueryValidatorsResponse{
		Validators: s.GetSimApp(s.chainB).StakingKeeper.GetBondedValidatorsByPower(s.chainB.GetContext()),
	}

	msg := types.MsgSubmitQueryResponse{
		ChainId:     s.chainB.ChainID + "-N",
		QueryId:     keeper.GenerateQueryHash(s.path.EndpointB.ConnectionID, s.chainB.ChainID, "cosmos.staking.v1beta1.Query/Validators", bz1, ""),
		Result:      s.GetSimApp(s.chainB).AppCodec().MustMarshalJSON(&qvr),
		Height:      s.chainB.CurrentHeader.Height,
		FromAddress: TestOwnerAddress,
	}

	s.NoError(msg.ValidateBasic())
	s.Equal(types.RouterKey, msg.Route())
	s.Equal(types.TypeMsgSubmitQueryResponse, msg.Type())
	s.Equal(TestOwnerAddress, msg.GetSigners()[0].String())
}
