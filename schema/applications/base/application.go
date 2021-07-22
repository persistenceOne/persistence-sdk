/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilityTypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisisKeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisisTypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributionKeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidenceKeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidenceTypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govKeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransferTypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintKeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramsProposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingKeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeKeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/persistenceOne/persistenceSDK/modules/assets"
	"github.com/persistenceOne/persistenceSDK/modules/classifications"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/deputize"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/maintain"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/revoke"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/modules/metas"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders"
	"github.com/persistenceOne/persistenceSDK/modules/splits"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	splitsMint "github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/renumerate"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tendermintLog "github.com/tendermint/tendermint/libs/log"
	tendermintOS "github.com/tendermint/tendermint/libs/os"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"honnef.co/go/tools/version"
	"io"
	"log"
	"os"
)

type application struct {
	name string
	//baseApp           *baseapp.BaseApp
	codec              *codec.LegacyAmino
	applicationCodec   codec.Marshaler
	interfaceRegistry  types.InterfaceRegistry
	moduleBasicManager module.BasicManager

	keys                        map[string]*sdkTypes.KVStoreKey
	enabledWasmProposalTypeList []wasm.ProposalType
	moduleAccountPermissions    map[string][]string
	tokenReceiveAllowedModules  map[string]bool

	stakingKeeper      stakingKeeper.Keeper
	slashingKeeper     slashingKeeper.Keeper
	distributionKeeper distributionKeeper.Keeper
	crisisKeeper       crisisKeeper.Keeper
	paramsKeeper       paramsKeeper.Keeper

	moduleManager *module.Manager

	baseapp.BaseApp
}

func (application application) MountStoreWithDB(key sdkTypes.StoreKey, typ sdkTypes.StoreType, db tendermintDB.DB) {
	panic("implement me")
}

func (application application) GetDefaultNodeHome() string {
	return os.ExpandEnv("$HOME/." + application.name + "/Node")
}

func (application application) GetModuleBasicManager() module.BasicManager {
	return application.moduleBasicManager
}

func (application application) RegisterAPIRoutes(server *api.Server, config config.APIConfig) {
	panic("implement me")
}

func (application application) RegisterGRPCServer(context client.Context, server grpc.Server) {
	panic("implement me")
}

func (application application) RegisterTxService(clientCtx client.Context) {
	panic("implement me")
}

func (application application) RegisterTendermintService(clientCtx client.Context) {
	panic("implement me")
}

func (application application) ListSnapshots(snapshots abciTypes.RequestListSnapshots) abciTypes.ResponseListSnapshots {
	panic("implement me")
}

func (application application) OfferSnapshot(snapshot abciTypes.RequestOfferSnapshot) abciTypes.ResponseOfferSnapshot {
	panic("implement me")
}

func (application application) LoadSnapshotChunk(chunk abciTypes.RequestLoadSnapshotChunk) abciTypes.ResponseLoadSnapshotChunk {
	panic("implement me")
}

func (application application) ApplySnapshotChunk(chunk abciTypes.RequestApplySnapshotChunk) abciTypes.ResponseApplySnapshotChunk {
	panic("implement me")
}

var _ applications.Application = (*application)(nil)

func (application application) Info(requestInfo abciTypes.RequestInfo) abciTypes.ResponseInfo {
	return application.Info(requestInfo)
}

func (application application) SetOption(requestSetOption abciTypes.RequestSetOption) abciTypes.ResponseSetOption {
	return application.SetOption(requestSetOption)
}

func (application application) Query(requestQuery abciTypes.RequestQuery) abciTypes.ResponseQuery {
	return application.Query(requestQuery)
}

func (application application) CheckTx(requestCheckTx abciTypes.RequestCheckTx) abciTypes.ResponseCheckTx {
	return application.CheckTx(requestCheckTx)
}

func (application application) InitChain(requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	return application.InitChain(requestInitChain)
}

func (application application) BeginBlock(requestBeginBlock abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return application.BeginBlock(requestBeginBlock)
}

func (application application) DeliverTx(requestDeliverTx abciTypes.RequestDeliverTx) abciTypes.ResponseDeliverTx {
	return application.DeliverTx(requestDeliverTx)
}

func (application application) EndBlock(requestEndBlock abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return application.EndBlock(requestEndBlock)
}

func (application application) Commit() abciTypes.ResponseCommit {
	return application.Commit()
}

func (application application) LoadHeight(height int64) error {
	return application.LoadVersion(height)
}
func (application application) ExportApplicationStateAndValidators(forZeroHeight bool, jailWhiteList []string) (serverTypes.ExportedApp, error) {
	context := application.NewContext(true, tendermintProto.Header{Height: application.LastBlockHeight()})

	height := application.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0
		applyWhiteList := false

		if len(jailWhiteList) > 0 {
			applyWhiteList = true
		}

		whiteListMap := make(map[string]bool)

		for _, address := range jailWhiteList {
			if _, Error := sdkTypes.ValAddressFromBech32(address); Error != nil {
				panic(Error)
			}

			whiteListMap[address] = true
		}

		application.crisisKeeper.AssertInvariants(context)

		application.stakingKeeper.IterateValidators(context, func(_ int64, val stakingTypes.ValidatorI) (stop bool) {
			_, _ = application.distributionKeeper.WithdrawValidatorCommission(context, val.GetOperator())
			return false
		})

		delegations := application.stakingKeeper.GetAllDelegations(context)
		for _, delegation := range delegations {
			validatorAddress, Error := sdkTypes.ValAddressFromBech32(delegation.ValidatorAddress)
			if Error != nil {
				panic(Error)
			}
			delegatorAddress, Error := sdkTypes.AccAddressFromBech32(delegation.DelegatorAddress)
			if Error != nil {
				panic(Error)
			}
			_, _ = application.distributionKeeper.WithdrawDelegationRewards(context, delegatorAddress, validatorAddress)
		}

		application.distributionKeeper.DeleteAllValidatorSlashEvents(context)

		application.distributionKeeper.DeleteAllValidatorHistoricalRewards(context)

		height := context.BlockHeight()
		context = context.WithBlockHeight(0)

		application.stakingKeeper.IterateValidators(context, func(_ int64, val stakingTypes.ValidatorI) (stop bool) {

			scraps := application.distributionKeeper.GetValidatorOutstandingRewardsCoins(context, val.GetOperator())
			feePool := application.distributionKeeper.GetFeePool(context)
			feePool.CommunityPool = feePool.CommunityPool.Add(scraps...)
			application.distributionKeeper.SetFeePool(context, feePool)

			application.distributionKeeper.Hooks().AfterValidatorCreated(context, val.GetOperator())
			return false
		})

		for _, delegation := range delegations {
			validatorAddress, Error := sdkTypes.ValAddressFromBech32(delegation.ValidatorAddress)
			if Error != nil {
				panic(Error)
			}
			delegatorAddress, Error := sdkTypes.AccAddressFromBech32(delegation.DelegatorAddress)
			if Error != nil {
				panic(Error)
			}
			application.distributionKeeper.Hooks().BeforeDelegationCreated(context, delegatorAddress, validatorAddress)
			application.distributionKeeper.Hooks().AfterDelegationModified(context, delegatorAddress, validatorAddress)
		}

		context = context.WithBlockHeight(height)

		application.stakingKeeper.IterateRedelegations(context, func(_ int64, redelegation stakingTypes.Redelegation) (stop bool) {
			for i := range redelegation.Entries {
				redelegation.Entries[i].CreationHeight = 0
			}
			application.stakingKeeper.SetRedelegation(context, redelegation)
			return false
		})

		application.stakingKeeper.IterateUnbondingDelegations(context, func(_ int64, unbondingDelegation stakingTypes.UnbondingDelegation) (stop bool) {
			for i := range unbondingDelegation.Entries {
				unbondingDelegation.Entries[i].CreationHeight = 0
			}
			application.stakingKeeper.SetUnbondingDelegation(context, unbondingDelegation)
			return false
		})

		store := context.KVStore(application.keys[stakingTypes.StoreKey])
		kvStoreReversePrefixIterator := sdkTypes.KVStoreReversePrefixIterator(store, stakingTypes.ValidatorsKey)
		counter := int16(0)

		for ; kvStoreReversePrefixIterator.Valid(); kvStoreReversePrefixIterator.Next() {
			addr := sdkTypes.ValAddress(kvStoreReversePrefixIterator.Key()[1:])
			validator, found := application.stakingKeeper.GetValidator(context, addr)

			if !found {
				panic("Validator not found!")
			}

			validator.UnbondingHeight = 0

			if applyWhiteList && !whiteListMap[addr.String()] {
				validator.Jailed = true
			}

			application.stakingKeeper.SetValidator(context, validator)
			counter++
		}

		kvStoreReversePrefixIterator.Close()

		_, Error := application.stakingKeeper.ApplyAndReturnValidatorSetUpdates(context)
		if Error != nil {
			log.Fatal(Error)
		}

		application.slashingKeeper.IterateValidatorSigningInfos(
			context,
			func(validatorConsAddress sdkTypes.ConsAddress, validatorSigningInfo slashingTypes.ValidatorSigningInfo) (stop bool) {
				validatorSigningInfo.StartHeight = 0
				application.slashingKeeper.SetValidatorSigningInfo(context, validatorConsAddress, validatorSigningInfo)
				return false
			},
		)
	}

	genesisState := application.moduleManager.ExportGenesis(context, application.applicationCodec)
	applicationState, Error := codec.MarshalJSONIndent(application.codec, genesisState)

	if Error != nil {
		return serverTypes.ExportedApp{}, Error
	}
	validators, err := staking.WriteValidators(context, application.stakingKeeper)

	return serverTypes.ExportedApp{
		AppState:        applicationState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: application.GetConsensusParams(context),
	}, err
}

func (application application) GetCodec() *codec.LegacyAmino {
	return application.codec
}

func (application application) Initialize(applicationName string, encodingConfig applications.EncodingConfig, logger tendermintLog.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) applications.Application {
	applicationCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	baseApp := baseapp.NewBaseApp(
		applicationName,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	baseApp.SetCommitMultiStoreTracer(traceStore)
	baseApp.SetAppVersion(version.Version)
	baseApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdkTypes.NewKVStoreKeys(
		authTypes.StoreKey, bankTypes.StoreKey, stakingTypes.StoreKey,
		mintTypes.StoreKey, distributionTypes.StoreKey, slashingTypes.StoreKey,
		govTypes.StoreKey, paramsTypes.StoreKey, ibchost.StoreKey, upgradeTypes.StoreKey,
		evidenceTypes.StoreKey, ibctransferTypes.StoreKey, capabilityTypes.StoreKey,
		wasm.StoreKey,
		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)

	transientStoreKeys := sdkTypes.NewTransientStoreKeys(paramsTypes.TStoreKey)
	memoryKeys := sdkTypes.NewMemoryStoreKeys(capabilityTypes.MemStoreKey)

	//application.ba = baseApp
	application.codec = legacyAmino
	application.applicationCodec = applicationCodec
	application.interfaceRegistry = interfaceRegistry
	application.keys = keys
	application.SetParamStore(application.paramsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramsKeeper.ConsensusParamsKeyTable()))

	paramsKeeper := paramsKeeper.NewKeeper(
		applicationCodec,
		legacyAmino,
		keys[paramsTypes.StoreKey],
		transientStoreKeys[paramsTypes.TStoreKey],
	)

	accountKeeper := authKeeper.NewAccountKeeper(
		applicationCodec,
		keys[authTypes.StoreKey],
		paramsKeeper.Subspace(authTypes.ModuleName),
		authTypes.ProtoBaseAccount,
		application.moduleAccountPermissions,
	)

	blacklistedAddresses := make(map[string]bool)
	for account := range application.moduleAccountPermissions {
		blacklistedAddresses[authTypes.NewModuleAddress(account).String()] = !application.tokenReceiveAllowedModules[account]
	}

	bankKeeper := bankKeeper.NewBaseKeeper(
		applicationCodec,
		keys[bankTypes.StoreKey],
		accountKeeper,
		paramsKeeper.Subspace(bankTypes.ModuleName),
		blacklistedAddresses,
	)

	stakingKeeper := stakingKeeper.NewKeeper(
		applicationCodec,
		keys[stakingTypes.StoreKey],
		accountKeeper,
		bankKeeper,
		paramsKeeper.Subspace(stakingTypes.ModuleName),
	)

	mintKeeper := mintKeeper.NewKeeper(
		applicationCodec,
		keys[mintTypes.StoreKey],
		paramsKeeper.Subspace(mintTypes.ModuleName),
		&stakingKeeper,
		accountKeeper,
		bankKeeper,
		authTypes.FeeCollectorName,
	)

	blackListedModuleAddresses := make(map[string]bool)
	for moduleAccount := range application.moduleAccountPermissions {
		blackListedModuleAddresses[authTypes.NewModuleAddress(moduleAccount).String()] = true
	}

	application.distributionKeeper = distributionKeeper.NewKeeper(
		applicationCodec,
		keys[distributionTypes.StoreKey],
		paramsKeeper.Subspace(distributionTypes.ModuleName),
		accountKeeper,
		bankKeeper,
		&stakingKeeper,
		authTypes.FeeCollectorName,
		blackListedModuleAddresses,
	)
	application.slashingKeeper = slashingKeeper.NewKeeper(
		applicationCodec,
		keys[slashingTypes.StoreKey],
		&stakingKeeper,
		paramsKeeper.Subspace(slashingTypes.ModuleName),
	)
	application.crisisKeeper = crisisKeeper.NewKeeper(
		paramsKeeper.Subspace(crisisTypes.ModuleName),
		invCheckPeriod,
		bankKeeper,
		authTypes.FeeCollectorName,
	)
	upgradeKeeper := upgradeKeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradeTypes.StoreKey],
		applicationCodec,
		home,
	)

	evidenceKeeper := evidenceKeeper.NewKeeper(
		applicationCodec,
		keys[evidenceTypes.StoreKey],
		&stakingKeeper,
		application.slashingKeeper,
	)

	govRouter := govTypes.NewRouter()
	govRouter.AddRoute(
		govTypes.RouterKey,
		govTypes.ProposalHandler,
	).AddRoute(
		paramsProposal.RouterKey,
		params.NewParamChangeProposalHandler(paramsKeeper),
	).AddRoute(
		distributionTypes.RouterKey,
		distribution.NewCommunityPoolSpendProposalHandler(application.distributionKeeper),
	).AddRoute(
		upgradeTypes.RouterKey,
		upgrade.NewSoftwareUpgradeProposalHandler(upgradeKeeper),
	)

	application.stakingKeeper = *stakingKeeper.SetHooks(
		stakingTypes.NewMultiStakingHooks(application.distributionKeeper.Hooks(), application.slashingKeeper.Hooks()),
	)

	metasModule := metas.Prototype().Initialize(
		keys[metas.Prototype().Name()],
		paramsKeeper.Subspace(metas.Prototype().Name()),
	)
	classificationsModule := classifications.Prototype().Initialize(
		keys[classifications.Prototype().Name()],
		paramsKeeper.Subspace(classifications.Prototype().Name()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
	)
	maintainersModule := maintainers.Prototype().Initialize(
		keys[metas.Prototype().Name()],
		paramsKeeper.Subspace(maintainers.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
	)
	identitiesModule := identities.Prototype().Initialize(
		keys[identities.Prototype().Name()],
		paramsKeeper.Subspace(identities.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(deputize.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(revoke.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),

		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
	)
	splitsModule := splits.Prototype().Initialize(
		keys[splits.Prototype().Name()],
		paramsKeeper.Subspace(splits.Prototype().Name()),
		bankKeeper,
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
	)
	assetsModule := assets.Prototype().Initialize(
		keys[assets.Prototype().Name()],
		paramsKeeper.Subspace(assets.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(deputize.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(revoke.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(splitsMint.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(burn.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(renumerate.Auxiliary.GetName()),
	)
	ordersModule := orders.Prototype().Initialize(
		keys[orders.Prototype().Name()],
		paramsKeeper.Subspace(orders.Prototype().Name()),
		classificationsModule.GetAuxiliary(conform.Auxiliary.GetName()),
		classificationsModule.GetAuxiliary(define.Auxiliary.GetName()),
		identitiesModule.GetAuxiliary(verify.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(super.Auxiliary.GetName()),
		maintainersModule.GetAuxiliary(maintain.Auxiliary.GetName()),
		metasModule.GetAuxiliary(scrub.Auxiliary.GetName()),
		metasModule.GetAuxiliary(supplement.Auxiliary.GetName()),
		splitsModule.GetAuxiliary(transfer.Auxiliary.GetName()),
	)

	//TODO:wasmkeeper changed commented for now

	//var wasmRouter = baseApp.Router()

	//wasmDir := filepath.Join(home, wasm.ModuleName)

	//wasmConfig, err := wasm.ReadWasmConfig(applicationOptions)
	//if err != nil {
	//	panic("error while reading wasm config: " + err.Error())
	//}
	//
	//wasmKeeper := wasm.NewKeeper(
	//	applicationCodec,
	//	keys[wasm.StoreKey],
	//	paramsKeeper.Subspace(wasm.DefaultParamspace),
	//	accountKeeper,
	//	bankKeeper,
	//	application.stakingKeeper,
	//	application.distributionKeeper,
	//	nil,
	//	nil,
	//	nil,
	//	nil,
	//	wasmRouter,
	//	nil,
	//	wasmDir,
	//	wasmConfig,
	//	&wasm.MessageEncoders{Custom: wasmUtilities.CustomEncoder(assets.Prototype(), classifications.Prototype(), identities.Prototype(), maintainers.Prototype(), metas.Prototype(), orders.Prototype(), splits.Prototype())},
	//	nil)

	//govRouter := gov.NewRouter().AddRoute(
	//	gov.RouterKey,
	//	gov.ProposalHandler,
	//).AddRoute(
	//	params.RouterKey,
	//	params.NewParamChangeProposalHandler(paramsKeeper),
	//).AddRoute(
	//	distribution.RouterKey,
	//	distribution.NewCommunityPoolSpendProposalHandler(application.distributionKeeper),
	//).AddRoute(
	//	upgrade.RouterKey,
	//	upgrade.NewSoftwareUpgradeProposalHandler(upgradeKeeper),
	//)

	//if len(application.enabledWasmProposalTypeList) != 0 {
	//	govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmKeeper, application.enabledWasmProposalTypeList))
	//}

	govKeeper := govKeeper.NewKeeper(
		applicationCodec,
		keys[govTypes.StoreKey],
		paramsKeeper.Subspace(govTypes.ModuleName).WithKeyTable(govTypes.ParamKeyTable()),
		accountKeeper,
		bankKeeper,
		&stakingKeeper,
		govRouter,
	)
	/****  Module Options ****/
	//var skipGenesisInvariants = false
	//opt := applicationOptions.Get(crisis.FlagSkipGenesisInvariants)
	//if opt, ok := opt.(bool); ok {
	//	skipGenesisInvariants = opt
	//}
	application.moduleManager = sdkTypesModule.NewManager(
		genutil.NewAppModule(accountKeeper, application.stakingKeeper, application.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(applicationCodec, accountKeeper, nil),
		bank.NewAppModule(applicationCodec, bankKeeper, accountKeeper),
		//crisis.NewAppModule(&application.crisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(applicationCodec, govKeeper, accountKeeper, bankKeeper),
		mint.NewAppModule(applicationCodec, mintKeeper, accountKeeper),
		slashing.NewAppModule(applicationCodec, application.slashingKeeper, accountKeeper, bankKeeper, application.stakingKeeper),
		distribution.NewAppModule(applicationCodec, application.distributionKeeper, accountKeeper, bankKeeper, application.stakingKeeper),
		staking.NewAppModule(applicationCodec, application.stakingKeeper, accountKeeper, bankKeeper),
		upgrade.NewAppModule(upgradeKeeper),
		//wasm.NewAppModule(applicationCodec,&wasmKeeper, stakingKeeper),
		evidence.NewAppModule(*evidenceKeeper),

		assetsModule,
		classificationsModule,
		identitiesModule,
		maintainersModule,
		metasModule,
		ordersModule,
		splitsModule,
	)

	application.moduleManager.SetOrderBeginBlockers(
		upgradeTypes.ModuleName,
		mintTypes.ModuleName,
		distributionTypes.ModuleName,
		slashingTypes.ModuleName,
	)
	application.moduleManager.SetOrderEndBlockers(
		crisisTypes.ModuleName,
		govTypes.ModuleName,
		stakingTypes.ModuleName,
		ordersModule.Name(),
	)
	application.moduleManager.SetOrderInitGenesis(
		authTypes.ModuleName,
		distributionTypes.ModuleName,
		stakingTypes.ModuleName,
		bankTypes.ModuleName,
		slashingTypes.ModuleName,
		govTypes.ModuleName,
		mintTypes.ModuleName,
		crisisTypes.ModuleName,
		genutilTypes.ModuleName,
		evidenceTypes.ModuleName,
		wasm.ModuleName,
		assets.Prototype().Name(),
		classifications.Prototype().Name(),
		identities.Prototype().Name(),
		maintainers.Prototype().Name(),
		metas.Prototype().Name(),
		orders.Prototype().Name(),
		splits.Prototype().Name(),
	)
	application.moduleManager.RegisterInvariants(&application.crisisKeeper)
	application.moduleManager.RegisterRoutes(application.Router(), application.QueryRouter(), legacyAmino)

	simulationManager := sdkTypesModule.NewSimulationManager(
		auth.NewAppModule(applicationCodec, accountKeeper, nil),
		bank.NewAppModule(applicationCodec, bankKeeper, accountKeeper),
		gov.NewAppModule(applicationCodec, govKeeper, accountKeeper, bankKeeper),
		mint.NewAppModule(applicationCodec, mintKeeper, accountKeeper),
		staking.NewAppModule(applicationCodec, application.stakingKeeper, accountKeeper, bankKeeper),
		distribution.NewAppModule(applicationCodec, application.distributionKeeper, accountKeeper, bankKeeper, application.stakingKeeper),
		slashing.NewAppModule(applicationCodec, application.slashingKeeper, accountKeeper, bankKeeper, application.stakingKeeper),
		params.NewAppModule(paramsKeeper),
		assets.Prototype(),
		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	)

	simulationManager.RegisterStoreDecoders()

	application.MountKVStores(keys)
	application.MountTransientStores(transientStoreKeys)
	application.MountMemoryStores(memoryKeys)

	application.SetBeginBlocker(application.moduleManager.BeginBlock)
	application.SetEndBlocker(application.moduleManager.EndBlock)
	application.SetInitChainer(func(context sdkTypes.Context, requestInitChain abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
		var genesisState map[string]json.RawMessage
		legacyAmino.MustUnmarshalJSON(requestInitChain.AppStateBytes, &genesisState)
		return application.moduleManager.InitGenesis(context, applicationCodec, genesisState)
	})
	application.SetAnteHandler(ante.NewAnteHandler(accountKeeper, bankKeeper, ante.DefaultSigVerificationGasConsumer, encodingConfig.TxConfig.SignModeHandler()))

	if loadLatest {
		err := application.LoadLatestVersion()
		if err != nil {
			tendermintOS.Exit(err.Error())
		}
	}

	return &application
}

func makeCodec(moduleBasicManager module.BasicManager) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	moduleBasicManager.RegisterLegacyAminoCodec(Codec)
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	return Codec
}

func NewApplication(name string, moduleBasicManager module.BasicManager, enabledWasmProposalTypeList []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool) applications.Application {

	return &application{
		name:                        name,
		moduleBasicManager:          moduleBasicManager,
		codec:                       makeCodec(moduleBasicManager),
		enabledWasmProposalTypeList: enabledWasmProposalTypeList,
		moduleAccountPermissions:    moduleAccountPermissions,
		tokenReceiveAllowedModules:  tokenReceiveAllowedModules,
	}
}
