package ibctransferhooks

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/persistenceOne/persistence-sdk/x/ibctransferhooks/keeper"
	"github.com/persistenceOne/persistence-sdk/x/ibctransferhooks/types"
)

var (
	_ module.AppModule      = Wrapper{}
	_ module.AppModuleBasic = AppModuleBasic{}
	_ porttypes.IBCModule   = Wrapper{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string { return types.ModuleName }

func (a AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(_ cdctypes.InterfaceRegistry) {}

func (a AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return nil
}

func (a AppModuleBasic) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

func (a AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

type Wrapper struct {
	AppModuleBasic
	k keeper.Keeper
	// ONLY PASS IBC TRANSFER APP
	ibcTransferApp porttypes.IBCModule
}

func NewAppModule(keeper keeper.Keeper) Wrapper {
	return Wrapper{
		AppModuleBasic: AppModuleBasic{},
		k:              keeper,
	}
}

func (w Wrapper) InitGenesis(_ sdk.Context, _ codec.JSONCodec, _ json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (w Wrapper) ExportGenesis(_ sdk.Context, _ codec.JSONCodec) json.RawMessage {
	return nil
}

func (w Wrapper) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (w Wrapper) Route() sdk.Route {
	return sdk.Route{}
}

func (w Wrapper) QuerierRoute() string { return "" }

func (w Wrapper) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

func (w Wrapper) RegisterServices(_ module.Configurator) {}

func (w Wrapper) ConsensusVersion() uint64 { return 1 }

func (w Wrapper) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (w Wrapper) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (w Wrapper) OnChanOpenInit(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string, channelID string, channelCap *capabilitytypes.Capability, counterparty channeltypes.Counterparty, version string) error {
	return w.ibcTransferApp.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, version)
}

func (w Wrapper) OnChanOpenTry(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID, channelID string, channelCap *capabilitytypes.Capability, counterparty channeltypes.Counterparty, counterpartyVersion string) (version string, err error) {
	return w.ibcTransferApp.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, counterpartyVersion)
}

func (w Wrapper) OnChanOpenAck(ctx sdk.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string) error {
	return w.ibcTransferApp.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

func (w Wrapper) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	return w.ibcTransferApp.OnChanOpenConfirm(ctx, portID, channelID)
}

func (w Wrapper) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	return w.ibcTransferApp.OnChanCloseInit(ctx, portID, channelID)
}

func (w Wrapper) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	return w.ibcTransferApp.OnChanCloseConfirm(ctx, portID, channelID)
}

func (w Wrapper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) exported.Acknowledgement {
	ack := w.ibcTransferApp.OnRecvPacket(ctx, packet, relayer)
	w.k.OnRecvPacket(ctx, packet, relayer, ack)

	return ack
}

func (w Wrapper) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	err := w.ibcTransferApp.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	w.k.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer, err)

	return err
}

func (w Wrapper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	err := w.ibcTransferApp.OnTimeoutPacket(ctx, packet, relayer)
	w.k.OnTimeoutPacket(ctx, packet, relayer, err)

	return err
}
