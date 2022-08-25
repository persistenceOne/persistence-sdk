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

func (suite *TypesTestSuite) GetSimApp(chain *ibctesting.TestChain) *simapp.SimApp {
	app, ok := chain.App.(*simapp.SimApp)
	if !ok {
		panic("not sim app")
	}

	return app
}

func (suite *TypesTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))

	suite.path = newSimAppPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(suite.path)
}

func (suite *TypesTestSuite) TestMsgSubmitQueryResponse() {
	bondedQuery := stakingtypes.QueryValidatorsRequest{Status: stakingtypes.BondStatusBonded}
	bz, err := bondedQuery.Marshal()
	suite.NoError(err)

	qvr := stakingtypes.QueryValidatorsResponse{
		Validators: suite.GetSimApp(suite.chainB).StakingKeeper.GetBondedValidatorsByPower(suite.chainB.GetContext()),
	}

	msg := types.MsgSubmitQueryResponse{
		ChainId:     suite.chainB.ChainID + "-N",
		QueryId:     keeper.GenerateQueryHash(suite.path.EndpointB.ConnectionID, suite.chainB.ChainID, "cosmos.staking.v1beta1.Query/Validators", bz, ""),
		Result:      suite.GetSimApp(suite.chainB).AppCodec().MustMarshalJSON(&qvr),
		Height:      suite.chainB.CurrentHeader.Height,
		FromAddress: TestOwnerAddress,
	}

	suite.NoError(msg.ValidateBasic())
	suite.Equal(types.RouterKey, msg.Route())
	suite.Equal(types.TypeMsgSubmitQueryResponse, msg.Type())
	suite.Equal(TestOwnerAddress, msg.GetSigners()[0].String())
}

func newSimAppPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort

	return path
}
