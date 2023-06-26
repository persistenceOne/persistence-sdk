package ibchooks_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
	"github.com/persistenceOne/persistence-sdk/v2/simapp"
	"github.com/persistenceOne/persistence-sdk/v2/tests/ibc-hooks/testutils"
	ibctestingWrapper "github.com/persistenceOne/persistence-sdk/v2/tests/ibctesting"
	"github.com/stretchr/testify/suite"
)

type HooksTestSuite struct {
	simapp.KeeperTestHelper

	coordinator *ibctesting.Coordinator

	chainA *ibctestingWrapper.TestChain
	chainB *ibctestingWrapper.TestChain
	chainC *ibctestingWrapper.TestChain

	pathAB *ibctesting.Path
	pathAC *ibctesting.Path
	pathBC *ibctesting.Path
}

// TODO: This needs to get removed. Waiting on https://github.com/cosmos/ibc-go/issues/3123
func (suite *HooksTestSuite) TearDownSuite() {
}

func TestIBCHooksTestSuite(t *testing.T) {
	suite.Run(t, new(HooksTestSuite))
}

func (suite *HooksTestSuite) SetupTest() {
	suite.Setup()

	ibctesting.DefaultTestingAppInit = ibctestingWrapper.InitSetupTestingApp(suite.T(), nil)

	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = &ibctestingWrapper.TestChain{
		TestChain: suite.coordinator.GetChain(ibctesting.GetChainID(1)),
	}
	suite.chainB = &ibctestingWrapper.TestChain{
		TestChain: suite.coordinator.GetChain(ibctesting.GetChainID(2)),
	}
	suite.chainC = &ibctestingWrapper.TestChain{
		TestChain: suite.coordinator.GetChain(ibctesting.GetChainID(3)),
	}

	err := suite.chainA.MoveEpochsToTheFuture()
	suite.Require().NoError(err)
	err = suite.chainB.MoveEpochsToTheFuture()
	suite.Require().NoError(err)
	err = suite.chainC.MoveEpochsToTheFuture()
	suite.Require().NoError(err)
	suite.pathAB = NewTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.Setup(suite.pathAB)
	suite.pathBC = NewTransferPath(suite.chainB, suite.chainC)
	suite.coordinator.Setup(suite.pathBC)
	suite.pathAC = NewTransferPath(suite.chainA, suite.chainC)
	suite.coordinator.Setup(suite.pathAC)
}

func NewTransferPath(chainA, chainB *ibctestingWrapper.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA.TestChain, chainB.TestChain)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointA.ChannelConfig.Version = transfertypes.Version
	path.EndpointB.ChannelConfig.Version = transfertypes.Version

	return path
}

type Chain int64

const (
	ChainA Chain = iota
	ChainB
	ChainC
)

func (suite *HooksTestSuite) GetChain(name Chain) *ibctestingWrapper.TestChain {
	switch name {
	case ChainA:
		return suite.chainA
	case ChainB:
		return suite.chainB
	case ChainC:
		return suite.chainC
	}
	return nil
}

func (suite *HooksTestSuite) TestOnRecvPacketHooks() {
	var (
		trace    transfertypes.DenomTrace
		amount   sdk.Int
		receiver string
		status   testutils.Status
	)

	testCases := []struct {
		msg      string
		malleate func(*testutils.Status)
		expPass  bool
	}{
		{"override", func(status *testutils.Status) {
			suite.chainB.GetSimApp().TransferStack.
				ICS4Middleware.Hooks = testutils.TestRecvOverrideHooks{Status: status}
		}, true},
		{"before and after", func(status *testutils.Status) {
			suite.chainB.GetSimApp().TransferStack.
				ICS4Middleware.Hooks = testutils.TestRecvBeforeAfterHooks{Status: status}
		}, true},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.msg, func() {
			suite.SetupTest() // reset

			path := NewTransferPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)
			receiver = suite.chainB.SenderAccount.GetAddress().String() // must be explicitly changed in malleate
			status = testutils.Status{}

			amount = sdk.NewInt(100) // must be explicitly changed in malleate
			seq := uint64(1)

			trace = transfertypes.ParseDenomTrace(sdk.DefaultBondDenom)

			// send coin from chainA to chainB
			transferMsg := transfertypes.NewMsgTransfer(
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				sdk.NewCoin(trace.IBCDenom(), amount),
				suite.chainA.SenderAccount.GetAddress().String(),
				receiver,
				clienttypes.NewHeight(1, 110),
				0,
				"",
			)
			_, err := suite.chainA.SendMsgs(transferMsg)
			suite.Require().NoError(err) // message committed

			tc.malleate(&status)

			data := transfertypes.NewFungibleTokenPacketData(
				trace.GetFullDenomPath(),
				amount.String(),
				suite.chainA.SenderAccount.GetAddress().String(),
				receiver,
				"",
			)

			packet := channeltypes.NewPacket(
				data.GetBytes(),
				seq,
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				path.EndpointB.ChannelConfig.PortID,
				path.EndpointB.ChannelID,
				clienttypes.NewHeight(1, 100),
				0,
			)

			ack := suite.chainB.GetSimApp().TransferStack.
				OnRecvPacket(suite.chainB.GetContext(), packet, suite.chainA.SenderAccount.GetAddress())

			if tc.expPass {
				suite.Require().True(ack.Success())
			} else {
				suite.Require().False(ack.Success())
			}

			if _, ok := suite.chainB.GetSimApp().TransferStack.
				ICS4Middleware.Hooks.(testutils.TestRecvOverrideHooks); ok {
				suite.Require().True(status.OverrideRan)
				suite.Require().False(status.BeforeRan)
				suite.Require().False(status.AfterRan)
			}

			if _, ok := suite.chainB.GetSimApp().TransferStack.
				ICS4Middleware.Hooks.(testutils.TestRecvBeforeAfterHooks); ok {
				suite.Require().False(status.OverrideRan)
				suite.Require().True(status.BeforeRan)
				suite.Require().True(status.AfterRan)
			}
		})
	}
}
