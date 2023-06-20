package oracle

import (
	"context"
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/client/cli"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/keeper"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

var (
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModule           = AppModule{}
	_ module.EndBlockAppModule   = AppModule{}
	_ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic implements the AppModuleBasic interface for the x/oracle module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// RegisterLegacyAminoCodec registers the x/oracle module's types with a legacy
// Amino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func NewAppModuleBasic(cdc codec.Codec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the x/oracle module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) ConsensusVersion() uint64 {
	return 1
}

// RegisterInterfaces registers the x/oracle module's interface types.
func (AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns the x/oracle module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the x/oracle module.
func (AppModuleBasic) ValidateGenesis(
	cdc codec.JSONCodec,
	config client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the x/oracle
// module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// GetTxCmd returns the x/oracle module's root tx command.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the x/oracle module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// AppModule implements the AppModule interface for the x/oracle module.
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// Name returns the x/oracle module's name.
func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}

// RegisterServices registers gRPC services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQuerier(am.keeper))
}

// RegisterInvariants registers the x/oracle module's invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the x/oracle module's genesis initialization. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState

	cdc.MustUnmarshalJSON(gs, &genState)
	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the x/oracle module's exported genesis state as raw
// JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// EndBlock executes all ABCI EndBlock logic respective to the x/oracle module.
// It returns no validator updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	if err := EndBlocker(ctx, am.keeper); err != nil {
		panic(err)
	}

	return []abci.ValidatorUpdate{}
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the pool-incentives module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	// simulation.RandomizedGenState(simState)
}

// RegisterStoreDecoder registers a decoder for supply module's types
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return nil // TODO
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }
