package ibc_hooks_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
	"github.com/persistenceOne/persistence-sdk/v2/simapp"
	"github.com/persistenceOne/persistence-sdk/v2/tests/ibc-hooks/testutils"
	ibctestingWrapper "github.com/persistenceOne/persistence-sdk/v2/tests/ibctesting"
	"github.com/persistenceOne/persistence-sdk/v2/utils"
	ibchookskeeper "github.com/persistenceOne/persistence-sdk/v2/x/ibc-hooks/keeper"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
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
	// This is used to test cw20s. It will only get assigned in the cw20 test
	pathCW20 *ibctesting.Path
}

const defaultPoolAmount int64 = 100000

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

type Direction int64

const (
	AtoB Direction = iota
	BtoA
	AtoC
	CtoA
	BtoC
	CtoB
	CW20toA
	AtoCW20
)

func (suite *HooksTestSuite) GetEndpoints(direction Direction) (sender *ibctesting.Endpoint, receiver *ibctesting.Endpoint) {
	switch direction {
	case AtoB:
		sender = suite.pathAB.EndpointA
		receiver = suite.pathAB.EndpointB
	case BtoA:
		sender = suite.pathAB.EndpointB
		receiver = suite.pathAB.EndpointA
	case AtoC:
		sender = suite.pathAC.EndpointA
		receiver = suite.pathAC.EndpointB
	case CtoA:
		sender = suite.pathAC.EndpointB
		receiver = suite.pathAC.EndpointA
	case BtoC:
		sender = suite.pathBC.EndpointA
		receiver = suite.pathBC.EndpointB
	case CtoB:
		sender = suite.pathBC.EndpointB
		receiver = suite.pathBC.EndpointA
	case CW20toA:
		sender = suite.pathCW20.EndpointB
		receiver = suite.pathCW20.EndpointA
	case AtoCW20:
		sender = suite.pathCW20.EndpointA
		receiver = suite.pathCW20.EndpointB
	default:
		panic("invalid direction")
	}
	return sender, receiver
}

// Get direction from chain pair
func (suite *HooksTestSuite) GetDirection(chainA, chainB Chain) Direction {
	switch {
	case chainA == ChainA && chainB == ChainB:
		return AtoB
	case chainA == ChainB && chainB == ChainA:
		return BtoA
	case chainA == ChainA && chainB == ChainC:
		return AtoC
	case chainA == ChainC && chainB == ChainA:
		return CtoA
	case chainA == ChainB && chainB == ChainC:
		return BtoC
	case chainA == ChainC && chainB == ChainB:
		return CtoB
	default:
		fmt.Println(chainA, chainB)
		panic("invalid chain pair")
	}
}

func (suite *HooksTestSuite) GetSenderChannel(chainA, chainB Chain) string {
	sender, _ := suite.GetEndpoints(suite.GetDirection(chainA, chainB))
	return sender.ChannelID
}

func (suite *HooksTestSuite) GetReceiverChannel(chainA, chainB Chain) string {
	_, receiver := suite.GetEndpoints(suite.GetDirection(chainA, chainB))
	return receiver.ChannelID
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
			transferMsg := transfertypes.NewMsgTransfer(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.NewCoin(trace.IBCDenom(), amount), suite.chainA.SenderAccount.GetAddress().String(), receiver, clienttypes.NewHeight(1, 110), 0)
			_, err := suite.chainA.SendMsgs(transferMsg)
			suite.Require().NoError(err) // message committed

			tc.malleate(&status)

			data := transfertypes.NewFungibleTokenPacketData(trace.GetFullDenomPath(), amount.String(), suite.chainA.SenderAccount.GetAddress().String(), receiver)
			packet := channeltypes.NewPacket(data.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(1, 100), 0)

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

func (suite *HooksTestSuite) makeMockPacket(receiver, memo string, prevSequence uint64) channeltypes.Packet {
	packetData := transfertypes.FungibleTokenPacketData{
		Denom:    sdk.DefaultBondDenom,
		Amount:   "1",
		Sender:   suite.chainB.SenderAccount.GetAddress().String(),
		Receiver: receiver,
		Memo:     memo,
	}

	return channeltypes.NewPacket(
		packetData.GetBytes(),
		prevSequence+1,
		suite.pathAB.EndpointB.ChannelConfig.PortID,
		suite.pathAB.EndpointB.ChannelID,
		suite.pathAB.EndpointA.ChannelConfig.PortID,
		suite.pathAB.EndpointA.ChannelID,
		clienttypes.NewHeight(0, 100),
		0,
	)
}

func (suite *HooksTestSuite) receivePacket(receiver, memo string) []byte {
	return suite.receivePacketWithSequence(receiver, memo, 0)
}

func (suite *HooksTestSuite) receivePacketWithSequence(receiver, memo string, prevSequence uint64) []byte {
	channelCap := suite.chainB.GetChannelCapability(
		suite.pathAB.EndpointB.ChannelConfig.PortID,
		suite.pathAB.EndpointB.ChannelID)

	packet := suite.makeMockPacket(receiver, memo, prevSequence)

	err := suite.chainB.GetSimApp().HooksICS4Wrapper.SendPacket(
		suite.chainB.GetContext(), channelCap, packet)
	suite.Require().NoError(err, "IBC send failed. Expected success. %s", err)

	// Update both clients
	err = suite.pathAB.EndpointB.UpdateClient()
	suite.Require().NoError(err)
	err = suite.pathAB.EndpointA.UpdateClient()
	suite.Require().NoError(err)

	// recv in chain a
	res, err := suite.pathAB.EndpointA.RecvPacketWithResult(packet)
	suite.Require().NoError(err)

	// get the ack from the chain a's response
	ack, err := ibctesting.ParseAckFromEvents(res.GetEvents())
	suite.Require().NoError(err)

	// manually send the acknowledgement to chain b
	err = suite.pathAB.EndpointA.AcknowledgePacket(packet, ack)
	suite.Require().NoError(err)
	return ack
}

func (suite *HooksTestSuite) TestRecvTransferWithMetadata() {
	// Setup contract
	suite.chainA.StoreContractCode(&suite.Suite, "./bytecode/echo.wasm")
	addr := suite.chainA.InstantiateContract(&suite.Suite, "{}", 1)

	ackBytes := suite.receivePacket(addr.String(), fmt.Sprintf(`{"wasm": {"contract": "%s", "msg": {"echo": {"msg": "test"} } } }`, addr))
	ackStr := string(ackBytes)
	fmt.Println(ackStr)
	var ack map[string]string // This can't be unmarshalled to Acknowledgement because it's fetched from the events
	err := json.Unmarshal(ackBytes, &ack)
	suite.Require().NoError(err)
	suite.Require().NotContains(ack, "error")
	suite.Require().Equal(ack["result"], "eyJjb250cmFjdF9yZXN1bHQiOiJkR2hwY3lCemFHOTFiR1FnWldOb2J3PT0iLCJpYmNfYWNrIjoiZXlKeVpYTjFiSFFpT2lKQlVUMDlJbjA9In0=")
}

// After successfully executing a wasm call, the contract should have the funds sent via IBC
func (suite *HooksTestSuite) TestFundsAreTransferredToTheContract() {
	// Setup contract
	suite.chainA.StoreContractCode(&suite.Suite, "./bytecode/echo.wasm")
	addr := suite.chainA.InstantiateContract(&suite.Suite, "{}", 1)

	// Check that the contract has no funds
	localDenom := utils.MustExtractDenomFromPacketOnRecv(suite.makeMockPacket("", "", 0))
	balance := suite.chainA.GetSimApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(0), balance.Amount)

	// Execute the contract via IBC
	ackBytes := suite.receivePacket(addr.String(), fmt.Sprintf(`{"wasm": {"contract": "%s", "msg": {"echo": {"msg": "test"} } } }`, addr))
	ackStr := string(ackBytes)
	fmt.Println(ackStr)
	var ack map[string]string // This can't be unmarshalled to Acknowledgement because it's fetched from the events
	err := json.Unmarshal(ackBytes, &ack)
	suite.Require().NoError(err)
	suite.Require().NotContains(ack, "error")
	suite.Require().Equal(ack["result"], "eyJjb250cmFjdF9yZXN1bHQiOiJkR2hwY3lCemFHOTFiR1FnWldOb2J3PT0iLCJpYmNfYWNrIjoiZXlKeVpYTjFiSFFpT2lKQlVUMDlJbjA9In0=")

	// Check that the token has now been transferred to the contract
	balance = suite.chainA.GetSimApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(1), balance.Amount)
}

// If the wasm call wails, the contract acknowledgement should be an error and the funds returned
func (suite *HooksTestSuite) TestFundsAreReturnedOnFailedContractExec() {
	// Setup contract
	suite.chainA.StoreContractCode(&suite.Suite, "./bytecode/echo.wasm")
	addr := suite.chainA.InstantiateContract(&suite.Suite, "{}", 1)

	// Check that the contract has no funds
	localDenom := utils.MustExtractDenomFromPacketOnRecv(suite.makeMockPacket("", "", 0))
	balance := suite.chainA.GetSimApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(0), balance.Amount)

	// Execute the contract via IBC with a message that the contract will reject
	ackBytes := suite.receivePacket(addr.String(), fmt.Sprintf(`{"wasm": {"contract": "%s", "msg": {"not_echo": {"msg": "test"} } } }`, addr))
	ackStr := string(ackBytes)
	fmt.Println(ackStr)
	var ack map[string]string // This can't be unmarshalled to Acknowledgement because it's fetched from the events
	err := json.Unmarshal(ackBytes, &ack)
	suite.Require().NoError(err)
	suite.Require().Contains(ack, "error")

	// Check that the token has now been transferred to the contract
	balance = suite.chainA.GetSimApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	fmt.Println(balance)
	suite.Require().Equal(sdk.NewInt(0), balance.Amount)
}

func (suite *HooksTestSuite) TestPacketsThatShouldBeSkipped() {
	var sequence uint64
	receiver := suite.chainB.SenderAccount.GetAddress().String()

	testCases := []struct {
		memo           string
		expPassthrough bool
	}{
		{"", true},
		{"{01]", true}, // bad json
		{"{}", true},
		{`{"something": ""}`, true},
		{`{"wasm": "test"}`, false},
		{`{"wasm": []`, true}, // invalid top level JSON
		{`{"wasm": {}`, true}, // invalid top level JSON
		{`{"wasm": []}`, false},
		{`{"wasm": {}}`, false},
		{`{"wasm": {"contract": "something"}}`, false},
		{`{"wasm": {"contract": "osmo1clpqr4nrk4khgkxj78fcwwh6dl3uw4epasmvnj"}}`, false},
		{`{"wasm": {"msg": "something"}}`, false},
		// invalid receiver
		{`{"wasm": {"contract": "osmo1clpqr4nrk4khgkxj78fcwwh6dl3uw4epasmvnj", "msg": {}}}`, false},
		// msg not an object
		{fmt.Sprintf(`{"wasm": {"contract": "%s", "msg": 1}}`, receiver), false},
	}

	for _, tc := range testCases {
		ackBytes := suite.receivePacketWithSequence(receiver, tc.memo, sequence)
		ackStr := string(ackBytes)
		fmt.Println(ackStr)
		var ack map[string]string // This can't be unmarshalled to Acknowledgement because it's fetched from the events
		err := json.Unmarshal(ackBytes, &ack)
		suite.Require().NoError(err)
		if tc.expPassthrough {
			suite.Require().Equal("AQ==", ack["result"], tc.memo)
		} else {
			suite.Require().Contains(ackStr, "error", tc.memo)
		}
		sequence += 1
	}
}

// After successfully executing a wasm call, the contract should have the funds sent via IBC
func (suite *HooksTestSuite) TestFundTracking() {
	// Setup contract
	suite.chainA.StoreContractCode(&suite.Suite, "./bytecode/counter.wasm")
	addr := suite.chainA.InstantiateContract(&suite.Suite, `{"count": 0}`, 1)

	// Check that the contract has no funds
	localDenom := utils.MustExtractDenomFromPacketOnRecv(suite.makeMockPacket("", "", 0))
	balance := suite.chainA.GetSimApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(0), balance.Amount)

	// Execute the contract via IBC
	suite.receivePacket(
		addr.String(),
		fmt.Sprintf(`{"wasm": {"contract": "%s", "msg": {"increment": {} } } }`, addr))

	senderLocalAcc, err := ibchookskeeper.DeriveIntermediateSender("channel-0", suite.chainB.SenderAccount.GetAddress().String(), "osmo")
	suite.Require().NoError(err)

	state := suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, senderLocalAcc)))
	suite.Require().Equal(`{"count":0}`, state)

	state = suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_total_funds": {"addr": "%s"}}`, senderLocalAcc)))
	suite.Require().Equal(`{"total_funds":[{"denom":"ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878","amount":"1"}]}`, state)

	suite.receivePacketWithSequence(
		addr.String(),
		fmt.Sprintf(`{"wasm": {"contract": "%s", "msg": {"increment": {} } } }`, addr), 1)

	state = suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, senderLocalAcc)))
	suite.Require().Equal(`{"count":1}`, state)

	state = suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_total_funds": {"addr": "%s"}}`, senderLocalAcc)))
	suite.Require().Equal(`{"total_funds":[{"denom":"ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878","amount":"2"}]}`, state)

	// Check that the token has now been transferred to the contract
	balance = suite.chainA.GetSimApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(2), balance.Amount)
}

// custom MsgTransfer constructor that supports Memo
func NewMsgTransfer(token sdk.Coin, sender, receiver, channel, memo string) *transfertypes.MsgTransfer {
	return &transfertypes.MsgTransfer{
		SourcePort:       "transfer",
		SourceChannel:    channel,
		Token:            token,
		Sender:           sender,
		Receiver:         receiver,
		TimeoutHeight:    clienttypes.NewHeight(0, 500),
		TimeoutTimestamp: 0,
		Memo:             memo,
	}
}

func (suite *HooksTestSuite) RelayPacket(packet channeltypes.Packet, direction Direction) (*sdk.Result, []byte) {
	sender, receiver := suite.GetEndpoints(direction)

	err := receiver.UpdateClient()
	suite.Require().NoError(err)

	// receiver Receives
	receiveResult, err := receiver.RecvPacketWithResult(packet)
	suite.Require().NoError(err)

	ack, err := ibctesting.ParseAckFromEvents(receiveResult.GetEvents())
	suite.Require().NoError(err)

	if strings.Contains(string(ack), "error") {
		errorCtx := gjson.Get(receiveResult.Log, "0.events.#(type==ibc-acknowledgement-error)#.attributes.#(key==error-context)#.value")
		fmt.Println("ibc-ack-error:", errorCtx)
	}

	// sender Acknowledges
	err = sender.AcknowledgePacket(packet, ack)
	suite.Require().NoError(err)

	err = sender.UpdateClient()
	suite.Require().NoError(err)
	err = receiver.UpdateClient()
	suite.Require().NoError(err)

	return receiveResult, ack
}

func (suite *HooksTestSuite) RelayPacketNoAck(packet channeltypes.Packet, direction Direction) *sdk.Result {
	sender, receiver := suite.GetEndpoints(direction)

	err := receiver.UpdateClient()
	suite.Require().NoError(err)

	// receiver Receives
	receiveResult, err := receiver.RecvPacketWithResult(packet)
	suite.Require().NoError(err)

	err = sender.UpdateClient()
	suite.Require().NoError(err)
	err = receiver.UpdateClient()
	suite.Require().NoError(err)

	return receiveResult
}

func (suite *HooksTestSuite) FullSend(msg sdk.Msg, direction Direction) (*sdk.Result, *sdk.Result, string, error) {
	var sender *ibctestingWrapper.TestChain
	switch direction {
	case AtoB:
		sender = suite.chainA
	case BtoA:
		sender = suite.chainB
	case BtoC:
		sender = suite.chainB
	case CtoB:
		sender = suite.chainC
	case AtoC:
		sender = suite.chainA
	case CtoA:
		sender = suite.chainC
	case CW20toA:
		sender = suite.chainB
	case AtoCW20:
		sender = suite.chainA
	default:
		panic("invalid direction")
	}
	sendResult, err := sender.SendMsgsNoCheck(msg)
	suite.Require().NoError(err)

	packet, err := ibctesting.ParsePacketFromEvents(sendResult.GetEvents())
	suite.Require().NoError(err)

	receiveResult, ack := suite.RelayPacket(packet, direction)
	if strings.Contains(string(ack), "error") {
		errorCtx := gjson.Get(receiveResult.Log, "0.events.#(type==ibc-acknowledgement-error)#.attributes.#(key==error-context)#.value")
		fmt.Println("ibc-ack-error:", errorCtx)
	}
	return sendResult, receiveResult, string(ack), err
}

func (suite *HooksTestSuite) TestAcks() {
	suite.chainA.StoreContractCode(&suite.Suite, "./bytecode/counter.wasm")
	addr := suite.chainA.InstantiateContract(&suite.Suite, `{"count": 0}`, 1)

	// Generate swap instructions for the contract
	callbackMemo := fmt.Sprintf(`{"ibc_callback":"%s"}`, addr)
	// Send IBC transfer with the memo with crosschain-swap instructions
	transferMsg := NewMsgTransfer(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)), suite.chainA.SenderAccount.GetAddress().String(), addr.String(), "channel-0", callbackMemo)
	_, _, _, err := suite.FullSend(transferMsg, AtoB)
	suite.Require().NoError(err)

	// The test contract will increment the counter for itself every time it receives an ack
	state := suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, addr)))
	suite.Require().Equal(`{"count":1}`, state)

	_, _, _, err = suite.FullSend(transferMsg, AtoB)
	suite.Require().NoError(err)
	state = suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, addr)))
	suite.Require().Equal(`{"count":2}`, state)
}

func (suite *HooksTestSuite) TestTimeouts() {
	suite.chainA.StoreContractCode(&suite.Suite, "./bytecode/counter.wasm")
	addr := suite.chainA.InstantiateContract(&suite.Suite, `{"count": 0}`, 1)

	// Generate swap instructions for the contract
	callbackMemo := fmt.Sprintf(`{"ibc_callback":"%s"}`, addr)
	// Send IBC transfer with the memo with crosschain-swap instructions
	transferMsg := NewMsgTransfer(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)), suite.chainA.SenderAccount.GetAddress().String(), addr.String(), "channel-0", callbackMemo)
	transferMsg.TimeoutTimestamp = uint64(suite.coordinator.CurrentTime.Add(time.Minute).UnixNano())
	sendResult, err := suite.chainA.SendMsgsNoCheck(transferMsg)
	suite.Require().NoError(err)

	packet, err := ibctesting.ParsePacketFromEvents(sendResult.GetEvents())
	suite.Require().NoError(err)

	// Move chainB forward one block
	suite.chainB.NextBlock()
	// One month later
	suite.coordinator.IncrementTimeBy(time.Hour)
	err = suite.pathAB.EndpointA.UpdateClient()
	suite.Require().NoError(err)

	err = suite.pathAB.EndpointA.TimeoutPacket(packet)
	suite.Require().NoError(err)

	// The test contract will increment the counter for itself by 10 when a packet times out
	state := suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, addr)))
	suite.Require().Equal(`{"count":10}`, state)
}

func (suite *HooksTestSuite) TestSendWithoutMemo() {
	// Sending a packet without memo to ensure that the ibc_callback middleware doesn't interfere with a regular send
	transferMsg := NewMsgTransfer(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)), suite.chainA.SenderAccount.GetAddress().String(), suite.chainA.SenderAccount.GetAddress().String(), "channel-0", "")
	_, _, ack, err := suite.FullSend(transferMsg, AtoB)
	suite.Require().NoError(err)
	suite.Require().Contains(ack, "result")
}

func (suite *HooksTestSuite) fundAccount(chain *ibctestingWrapper.TestChain, owner sdk.AccAddress) {
	// TODO: allow this function to fund with custom token names (calling them tokenA, tokenB, etc. would make tests easier to read, I think)
	bankKeeper := chain.GetSimApp().BankKeeper
	i, ok := sdk.NewIntFromString("20000000000000000000000")
	suite.Require().True(ok)
	amounts := sdk.NewCoins(sdk.NewCoin("uxprt", i), sdk.NewCoin(sdk.DefaultBondDenom, i), sdk.NewCoin("token0", i), sdk.NewCoin("token1", i))
	err := bankKeeper.MintCoins(chain.GetContext(), minttypes.ModuleName, amounts)
	suite.Require().NoError(err)
	err = bankKeeper.SendCoinsFromModuleToAccount(chain.GetContext(), minttypes.ModuleName, owner, amounts)
	suite.Require().NoError(err)
}

func (suite *HooksTestSuite) setChainChannelLinks(registryAddr sdk.AccAddress, chainName Chain) {
	chain := suite.GetChain(chainName)
	ctx := chain.GetContext()
	owner := chain.SenderAccount.GetAddress()
	osmosisApp := chain.GetSimApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(osmosisApp.WasmKeeper)

	// Add all chain channel links in a single message
	msg := `{
		"modify_chain_channel_links": {
		  "operations": [
			{"operation": "set","source_chain": "chainB","destination_chain": "osmosis","channel_id": "channel-0"},
			{"operation": "set","source_chain": "osmosis","destination_chain": "chainB","channel_id": "channel-0"},
			{"operation": "set","source_chain": "chainB","destination_chain": "chainC","channel_id": "channel-1"},
			{"operation": "set","source_chain": "chainC","destination_chain": "chainB","channel_id": "channel-0"},
			{"operation": "set","source_chain": "osmosis","destination_chain": "chainC","channel_id": "channel-1"},
			{"operation": "set","source_chain": "chainC","destination_chain": "osmosis","channel_id": "channel-1"},
			{"operation": "set","source_chain": "osmosis","destination_chain": "chainB-cw20","channel_id": "channel-2"},
			{"operation": "set","source_chain": "chainB-cw20","destination_chain": "osmosis","channel_id": "channel-2"}
		  ]
		}
	  }
	  `
	_, err := contractKeeper.Execute(ctx, registryAddr, owner, []byte(msg), sdk.NewCoins())
	suite.Require().NoError(err)
}

// modifyChainChannelLinks modifies the chain channel links in the crosschain registry utilizing set, remove, and change operations
func (suite *HooksTestSuite) modifyChainChannelLinks(registryAddr sdk.AccAddress, chainName Chain) {
	chain := suite.GetChain(chainName)
	ctx := chain.GetContext()
	owner := chain.SenderAccount.GetAddress()
	osmosisApp := chain.GetSimApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(osmosisApp.WasmKeeper)

	msg := `{
		"modify_chain_channel_links": {
		  "operations": [
			{"operation": "remove","source_chain": "chainB","destination_chain": "chainC","channel_id": "channel-1"},
			{"operation": "set","source_chain": "chainD","destination_chain": "ChainC","channel_id": "channel-1"},
			{"operation": "remove","source_chain": "chainC","destination_chain": "chainB","channel_id": "channel-0"},
			{"operation": "set","source_chain": "ChainC","destination_chain": "chainD","channel_id": "channel-0"},
			{"operation": "change","source_chain": "chainB","destination_chain": "osmosis","new_source_chain": "chainD"},
			{"operation": "change","source_chain": "osmosis","destination_chain": "chainB","new_destination_chain": "chainD"}
		  ]
		}
	  }
	  `
	_, err := contractKeeper.Execute(ctx, registryAddr, owner, []byte(msg), sdk.NewCoins())
	suite.Require().NoError(err)
}

// modifyChainChannelLinks modifies the chain channel links in the crosschain registry utilizing set, remove, and change operations
func (suite *HooksTestSuite) setContractAlias(registryAddr sdk.AccAddress, contractAlias string, chainName Chain) {
	chain := suite.GetChain(chainName)
	ctx := chain.GetContext()
	owner := chain.SenderAccount.GetAddress()
	osmosisApp := chain.GetSimApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(osmosisApp.WasmKeeper)

	msg := fmt.Sprintf(`{
		"modify_contract_alias": {
		  "operations": [
			{"operation": "set", "alias": "%s", "address": "%s"}
		  ]
		}
	  }
	  `, contractAlias, registryAddr)
	_, err := contractKeeper.Execute(ctx, registryAddr, owner, []byte(msg), sdk.NewCoins())
	suite.Require().NoError(err)
}

func (suite *HooksTestSuite) SetupIBCRouteOnChain(swaprouterAddr, owner sdk.AccAddress, poolId uint64, chainName Chain, denom string) {
	chain := suite.GetChain(chainName)
	osmosisApp := chain.GetSimApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(osmosisApp.WasmKeeper)

	msg := fmt.Sprintf(`{"set_route":{"input_denom":"%s","output_denom":"token0","pool_route":[{"pool_id":"%v","token_out_denom":"%s"},{"pool_id":"1","token_out_denom":"token0"}]}}`,
		denom, poolId, sdk.DefaultBondDenom)
	_, err := contractKeeper.Execute(chain.GetContext(), swaprouterAddr, owner, []byte(msg), sdk.NewCoins())
	suite.Require().NoError(err)

	msg2 := fmt.Sprintf(`{"set_route":{"input_denom":"token0","output_denom":"%s","pool_route":[{"pool_id":"1","token_out_denom":"%s"},{"pool_id":"%v","token_out_denom":"%s"}]}}`,
		denom, sdk.DefaultBondDenom, poolId, denom)
	_, err = contractKeeper.Execute(chain.GetContext(), swaprouterAddr, owner, []byte(msg2), sdk.NewCoins())
	suite.Require().NoError(err)

	// Move forward one block
	chain.NextBlock()
	chain.Coordinator.IncrementTime()

	// Update both clients
	err = suite.pathAB.EndpointA.UpdateClient()
	suite.Require().NoError(err)
	err = suite.pathAB.EndpointB.UpdateClient()
	suite.Require().NoError(err)
}

func (suite *HooksTestSuite) SetupIBCSimpleRouteOnChain(swaprouterAddr, owner sdk.AccAddress, poolId uint64, chainName Chain, denom1, denom2 string) {
	chain := suite.GetChain(chainName)
	osmosisApp := chain.GetSimApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(osmosisApp.WasmKeeper)

	msg := fmt.Sprintf(`{"set_route":{"input_denom":"%s","output_denom":"%s","pool_route":[{"pool_id":"%v","token_out_denom":"%s"}]}}`,
		denom1, denom2, poolId, denom2)
	_, err := contractKeeper.Execute(chain.GetContext(), swaprouterAddr, owner, []byte(msg), sdk.NewCoins())
	suite.Require().NoError(err)

	msg2 := fmt.Sprintf(`{"set_route":{"input_denom":"%s","output_denom":"%s","pool_route":[{"pool_id":"%v","token_out_denom":"%s"}]}}`,
		denom2, denom1, poolId, denom1)
	_, err = contractKeeper.Execute(chain.GetContext(), swaprouterAddr, owner, []byte(msg2), sdk.NewCoins())
	suite.Require().NoError(err)

	// Move forward one block
	chain.NextBlock()
	chain.Coordinator.IncrementTime()

	// Update both clients
	err = suite.pathAB.EndpointA.UpdateClient()
	suite.Require().NoError(err)
	err = suite.pathAB.EndpointB.UpdateClient()
	suite.Require().NoError(err)
}

// simple transfer
func (suite *HooksTestSuite) SimpleNativeTransfer(token string, amount sdk.Int, path []Chain) string {
	prev := path[0]
	prevPrefix := ""
	denom := token
	for i, path := range path {
		if i == 0 {
			continue
		}
		fromChain := prev
		toChain := path
		transferMsg := NewMsgTransfer(
			sdk.NewCoin(denom, amount),
			suite.GetChain(prev).SenderAccount.GetAddress().String(),
			suite.GetChain(path).SenderAccount.GetAddress().String(),
			suite.GetSenderChannel(fromChain, toChain),
			"",
		)
		_, _, _, err := suite.FullSend(transferMsg, suite.GetDirection(fromChain, toChain))
		suite.Require().NoError(err)
		receiveChannel := suite.GetSenderChannel(toChain, fromChain)
		prevPrefix += "/" + strings.TrimRight(transfertypes.GetDenomPrefix("transfer", receiveChannel), "/")
		prevPrefix = strings.TrimLeft(prevPrefix, "/")
		denom = transfertypes.DenomTrace{Path: prevPrefix, BaseDenom: token}.IBCDenom()
		prev = toChain
	}
	return denom
}

func (suite *HooksTestSuite) GetPath(chain1, chain2 Chain) string {
	return fmt.Sprintf("transfer/%s", suite.GetReceiverChannel(chain1, chain2))
}

func (suite *HooksTestSuite) GetIBCDenom(a Chain, b Chain, denom string) string {
	return transfertypes.DenomTrace{Path: suite.GetPath(a, b), BaseDenom: denom}.IBCDenom()
}

type ChainActorDefinition struct {
	Chain
	name    string
	address sdk.AccAddress
}

func (suite *HooksTestSuite) TestOutpostSimplified() {
	initializer := suite.chainB.SenderAccount.GetAddress()
	suite.ExecuteOutpostSwap(initializer, initializer, fmt.Sprintf(`chainB/%s`, initializer.String()))
}

func (suite *HooksTestSuite) TestOutpostExplicit() {
	initializer := suite.chainB.SenderAccount.GetAddress()
	suite.ExecuteOutpostSwap(initializer, initializer, fmt.Sprintf(`ibc:channel-0/%s`, initializer.String()))
}
