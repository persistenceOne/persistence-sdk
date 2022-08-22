package ibctransferhooks

import (
	"encoding/json"
	types2 "github.com/persistenceOne/persistence-sdk/x/ibctransferhooks/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/persistenceOne/persistence-sdk/x/ibctransferhooks/keeper"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = IBCTransferHooksWrapper{}
	_ module.AppModuleBasic = AppModuleBasic{}
	_ porttypes.IBCModule   = IBCTransferHooksWrapper{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {

	return types2.ModuleName
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {

	return
}

func (a AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {

	return
}

func (a AppModuleBasic) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {

	return nil
}

func (a AppModuleBasic) ValidateGenesis(jsonCodec codec.JSONCodec, config client.TxEncodingConfig, message json.RawMessage) error {

	return nil
}

func (a AppModuleBasic) RegisterRESTRoutes(context client.Context, router *mux.Router) {

	return
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {

	return
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {

	return nil
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {

	return nil
}

type IBCTransferHooksWrapper struct {
	AppModuleBasic
	k keeper.Keeper
	// ONLY PASS IBC TRANSFER APP
	ibcTransferApp porttypes.IBCModule
}

func (I IBCTransferHooksWrapper) InitGenesis(context sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}

func (I IBCTransferHooksWrapper) ExportGenesis(context sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {

	return nil
}

func (I IBCTransferHooksWrapper) RegisterInvariants(registry sdk.InvariantRegistry) {

	return
}

func (I IBCTransferHooksWrapper) Route() sdk.Route {

	return sdk.Route{}
}

func (I IBCTransferHooksWrapper) QuerierRoute() string {

	return ""
}

func (I IBCTransferHooksWrapper) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {

	return nil
}

func (I IBCTransferHooksWrapper) RegisterServices(configurator module.Configurator) {

	return
}

func (I IBCTransferHooksWrapper) ConsensusVersion() uint64 {

	return 1
}

func (I IBCTransferHooksWrapper) BeginBlock(context sdk.Context, block abci.RequestBeginBlock) {

	return
}

func (I IBCTransferHooksWrapper) EndBlock(context sdk.Context, block abci.RequestEndBlock) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}

func (I IBCTransferHooksWrapper) OnChanOpenInit(ctx sdk.Context, order types.Order, connectionHops []string, portID string, channelID string, channelCap *capabilitytypes.Capability, counterparty types.Counterparty, version string) error {

	return I.ibcTransferApp.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, version)
}

func (I IBCTransferHooksWrapper) OnChanOpenTry(ctx sdk.Context, order types.Order, connectionHops []string, portID, channelID string, channelCap *capabilitytypes.Capability, counterparty types.Counterparty, counterpartyVersion string) (version string, err error) {

	return I.ibcTransferApp.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, counterpartyVersion)
}

func (I IBCTransferHooksWrapper) OnChanOpenAck(ctx sdk.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string) error {

	return I.ibcTransferApp.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

func (I IBCTransferHooksWrapper) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {

	return I.ibcTransferApp.OnChanOpenConfirm(ctx, portID, channelID)
}

func (I IBCTransferHooksWrapper) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {

	return I.ibcTransferApp.OnChanCloseInit(ctx, portID, channelID)
}

func (I IBCTransferHooksWrapper) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {

	return I.ibcTransferApp.OnChanCloseConfirm(ctx, portID, channelID)
}

func (I IBCTransferHooksWrapper) OnRecvPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress) exported.Acknowledgement {

	ack := I.ibcTransferApp.OnRecvPacket(ctx, packet, relayer)
	if ack.Success() {
		// TODO add OnRecvPacketHook
	}

	return ack
}

func (I IBCTransferHooksWrapper) OnAcknowledgementPacket(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {

	err := I.ibcTransferApp.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	if err == nil {
		// TODO add OnAcknowledgementPacketHook
	}
	return err
}

func (I IBCTransferHooksWrapper) OnTimeoutPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress) error {

	err := I.ibcTransferApp.OnTimeoutPacket(ctx, packet, relayer)
	if err == nil {
		// TODO add OnTimeoutPacketHook
	}

	return err
}
