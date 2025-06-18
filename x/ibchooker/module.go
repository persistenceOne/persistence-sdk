package ibchooker

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	ibchookerkeeper "github.com/persistenceOne/persistence-sdk/v3/x/ibchooker/keeper"
	"github.com/persistenceOne/persistence-sdk/v3/x/ibchooker/types"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ module.AppModule      = AppModule{}
	_ porttypes.IBCModule   = AppModule{}
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

type AppModule struct {
	AppModuleBasic
	keeper ibchookerkeeper.Keeper
	// ONLY PASS IBC TRANSFER APP
	ibcApp porttypes.IBCModule
}

func NewAppModule(keeper ibchookerkeeper.Keeper, ibcTransferApp porttypes.IBCModule) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
		ibcApp:         ibcTransferApp,
	}
}

func (am AppModule) InitGenesis(_ sdk.Context, _ codec.JSONCodec, _ json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(_ sdk.Context, _ codec.JSONCodec) json.RawMessage {
	return nil
}

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) RegisterServices(_ module.Configurator) {}

func (am AppModule) ConsensusVersion() uint64 { return 1 }

func (am AppModule) OnChanOpenInit(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string, channelID string, channelCap *capabilitytypes.Capability, counterparty channeltypes.Counterparty, version string) (string, error) {
	return am.ibcApp.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, version)
}

func (am AppModule) OnChanOpenTry(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID, channelID string, channelCap *capabilitytypes.Capability, counterparty channeltypes.Counterparty, counterpartyVersion string) (version string, err error) {
	return am.ibcApp.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, counterpartyVersion)
}

func (am AppModule) OnChanOpenAck(ctx sdk.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string) error {
	return am.ibcApp.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

func (am AppModule) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	return am.ibcApp.OnChanOpenConfirm(ctx, portID, channelID)
}

func (am AppModule) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	return am.ibcApp.OnChanCloseInit(ctx, portID, channelID)
}

func (am AppModule) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	return am.ibcApp.OnChanCloseConfirm(ctx, portID, channelID)
}

func (am AppModule) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) exported.Acknowledgement {
	ack := am.ibcApp.OnRecvPacket(ctx, packet, relayer)
	am.keeper.OnRecvPacket(ctx, packet, relayer, ack)

	return ack
}

func (am AppModule) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	err := am.ibcApp.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	am.keeper.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer, err)

	return err
}

func (am AppModule) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	err := am.ibcApp.OnTimeoutPacket(ctx, packet, relayer)
	am.keeper.OnTimeoutPacket(ctx, packet, relayer, err)

	return err
}
