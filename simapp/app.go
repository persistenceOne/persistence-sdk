package simapp

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cast"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tendermintos "github.com/cometbft/cometbft/libs/os"
	tendermintproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	ibctestingtypes "github.com/cosmos/ibc-go/v7/testing/types"

	"github.com/persistenceOne/persistence-sdk/v2/simapp/keepers"
	simappparams "github.com/persistenceOne/persistence-sdk/v2/simapp/params"
	"github.com/persistenceOne/persistence-sdk/v2/simapp/upgrades"
)

const appName = "SimApp"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(keepers.AppModuleBasics...)

	Upgrades = []upgrades.Upgrade{}
)

var (
	_ runtime.AppI            = (*SimApp)(nil)
	_ servertypes.Application = (*SimApp)(nil)
)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimApp struct {
	*baseapp.BaseApp
	*keepers.AppKeepers

	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	// the module manager
	ModuleManager *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".simapp")
}

// NewSimApp returns a reference to an initialized SimApp.
func NewSimApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SimApp {
	encodingConfig := makeEncodingConfig()

	appCodec := encodingConfig.Codec
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	bApp := baseapp.NewBaseApp(appName, logger, db, txConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	app := &SimApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		txConfig:          txConfig,
		interfaceRegistry: interfaceRegistry,
	}

	var wasmOpts []wasm.Option
	var enabledProposals = wasm.EnableAllProposals

	homePath := cast.ToString(appOpts.Get(flags.FlagHome))
	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// Setup keepers
	app.AppKeepers = keepers.NewAppKeeper(
		appCodec,
		bApp,
		legacyAmino,
		ModuleAccountPermissions,
		SendCoinBlockedAddrs(),
		appOpts,
		wasmDir,
		enabledProposals,
		wasmOpts,
		simappparams.Bech32MainPrefix,
	)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	app.ModuleManager = module.NewManager(appModules(app, encodingConfig, skipGenesisInvariants)...)

	app.ModuleManager.SetOrderBeginBlockers(orderBeginBlockers()...)
	app.ModuleManager.SetOrderEndBlockers(orderEndBlockers()...)
	app.ModuleManager.SetOrderInitGenesis(orderInitGenesis()...)
	app.ModuleManager.SetOrderExportGenesis(orderInitGenesis()...)

	// Uncomment if you want to set a custom migration order here.
	// app.ModuleManager.SetOrderMigrations(custom order)

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.ModuleManager.RegisterServices(app.configurator)

	app.sm = module.NewSimulationManagerFromAppModules(
		app.ModuleManager.Modules,
		overrideSimulationModules(app, encodingConfig, skipGenesisInvariants),
	)
	app.sm.RegisterStoreDecoders()

	app.RegisterGRPCServices()

	// initialize stores
	app.MountKVStores(app.GetKVStoreKey())
	app.MountTransientStores(app.GetTransientStoreKey())
	app.MountMemoryStores(app.GetMemoryStoreKey())

	app.setAnteHandler(encodingConfig.TxConfig, wasmConfig)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// app.setPostHandler()
	app.setupUpgradeHandlers()
	app.setupUpgradeStoreLoaders()

	// must be before Loading version
	// requires the snapshot store to be created and registered as a BaseAppOption
	// see cmd/wasmd/root.go: 206 - 214 approx
	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), app.WasmKeeper),
		)
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extension: %s", err))
		}
	}

	if loadLatest {
		if err := app.BaseApp.LoadLatestVersion(); err != nil {
			tendermintos.Exit(err.Error())
		}
		ctx := app.BaseApp.NewUncachedContext(true, tendermintproto.Header{})

		// Initialize pinned codes in wasmvm as they are not persisted there
		if err := app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tendermintos.Exit(fmt.Sprintf("failed initialize pinned codes %s", err))
		}
	}

	return app
}

func (app *SimApp) setAnteHandler(txConfig client.TxConfig, wasmConfig wasmtypes.WasmConfig) {
	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeegrantKeeper,
				SignModeHandler: txConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCKeeper:         app.IBCKeeper,
			WasmConfig:        &wasmConfig,
			TXCounterStoreKey: app.GetKVStoreKey()[wasm.StoreKey],
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}
	app.SetAnteHandler(anteHandler)
}

func (app *SimApp) RegisterGRPCServices() {
	autocliv1.RegisterQueryServer(
		app.GRPCQueryRouter(),
		runtimeservices.NewAutoCLIQueryService(app.ModuleManager.Modules),
	)

	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)
}

func (app *SimApp) setPostHandler() {
	postHandler, err := posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
	if err != nil {
		panic(err)
	}

	app.SetPostHandler(postHandler)
}

// Name returns the name of the App
func (app *SimApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *SimApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.ModuleManager.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *SimApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.ModuleManager.EndBlock(ctx, req)
}

func (a *SimApp) Configurator() module.Configurator {
	return a.configurator
}

// InitChainer application update at chain initialization
func (app *SimApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	return app.ModuleManager.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *SimApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns SimApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns SimApp's InterfaceRegistry
func (app *SimApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// TxConfig returns SimApp's TxConfig
func (app *SimApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// DefaultGenesis returns a default genesis from the registered AppModuleBasic's.
func (a *SimApp) DefaultGenesis() map[string]json.RawMessage {
	return ModuleBasics.DefaultGenesis(a.appCodec)
}

func SendCoinBlockedAddrs() map[string]bool {
	sendCoinBlockedAddrs := make(map[string]bool)
	for acc := range ModuleAccountPermissions {
		sendCoinBlockedAddrs[authtypes.NewModuleAddress(acc).String()] = !receiveAllowedMAcc[acc]
	}
	return sendCoinBlockedAddrs
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *SimApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

func (app *SimApp) setupUpgradeHandlers() {
	for _, upgrade := range Upgrades {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(upgrades.UpgradeHandlerArgs{
				ModuleManager:      app.ModuleManager,
				Configurator:       app.configurator,
				Keepers:            app.AppKeepers,
				Codec:              app.appCodec,
				CapabilityStoreKey: app.GetKVStoreKey()[capabilitytypes.StoreKey],
				CapabilityKeeper:   app.CapabilityKeeper,
			}),
		)
	}
}

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *SimApp) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &upgrade.StoreUpgrades))
		}
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *SimApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *SimApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

func (app *SimApp) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// TestingApp functions

// GetBaseApp implements the TestingApp interface.
func (app *SimApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetStakingKeeper implements the TestingApp interface.
func (app *SimApp) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.StakingKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *SimApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *SimApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *SimApp) GetTxConfig() client.TxConfig {
	return app.txConfig
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range GetMaccPerms() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	// allow the following addresses to receive funds
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

func makeEncodingConfig() simappparams.EncodingConfig {
	encodingConfig := simappparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
