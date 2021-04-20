package base

import (
	"bytes"
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	"github.com/tendermint/tendermint/libs/log"
	tendermintTypes "github.com/tendermint/tendermint/types"
	db "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

func TestNewApplication(t *testing.T) {
	type args struct {
		name                        string
		moduleBasicManager          module.BasicManager
		enabledWasmProposalTypeList []wasm.ProposalType
		moduleAccountPermissions    map[string][]string
		tokenReceiveAllowedModules  map[string]bool
	}
	var tests []struct {
		name string
		args args
		want applications.Application
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApplication(tt.args.name, tt.args.moduleBasicManager, tt.args.enabledWasmProposalTypeList, tt.args.moduleAccountPermissions, tt.args.tokenReceiveAllowedModules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_application_ExportApplicationStateAndValidators(t *testing.T) {
	type fields struct {
		name                        string
		moduleBasicManager          module.BasicManager
		codec                       *codec.Codec
		enabledWasmProposalTypeList []wasm.ProposalType
		moduleAccountPermissions    map[string][]string
		tokenReceiveAllowedModules  map[string]bool
		keys                        map[string]*sdkTypes.KVStoreKey
		stakingKeeper               staking.Keeper
		slashingKeeper              slashing.Keeper
		distributionKeeper          distribution.Keeper
		crisisKeeper                crisis.Keeper
		moduleManager               *module.Manager
		BaseApp                     baseapp.BaseApp
	}
	type args struct {
		forZeroHeight bool
		jailWhiteList []string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    json.RawMessage
		want1   []tendermintTypes.GenesisValidator
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application := application{
				name:                        tt.fields.name,
				moduleBasicManager:          tt.fields.moduleBasicManager,
				codec:                       tt.fields.codec,
				enabledWasmProposalTypeList: tt.fields.enabledWasmProposalTypeList,
				moduleAccountPermissions:    tt.fields.moduleAccountPermissions,
				tokenReceiveAllowedModules:  tt.fields.tokenReceiveAllowedModules,
				keys:                        tt.fields.keys,
				stakingKeeper:               tt.fields.stakingKeeper,
				slashingKeeper:              tt.fields.slashingKeeper,
				distributionKeeper:          tt.fields.distributionKeeper,
				crisisKeeper:                tt.fields.crisisKeeper,
				moduleManager:               tt.fields.moduleManager,
				BaseApp:                     tt.fields.BaseApp,
			}
			got, got1, err := application.ExportApplicationStateAndValidators(tt.args.forZeroHeight, tt.args.jailWhiteList)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExportApplicationStateAndValidators() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExportApplicationStateAndValidators() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ExportApplicationStateAndValidators() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_application_GetCodec(t *testing.T) {
	type fields struct {
		name                        string
		moduleBasicManager          module.BasicManager
		codec                       *codec.Codec
		enabledWasmProposalTypeList []wasm.ProposalType
		moduleAccountPermissions    map[string][]string
		tokenReceiveAllowedModules  map[string]bool
		keys                        map[string]*sdkTypes.KVStoreKey
		stakingKeeper               staking.Keeper
		slashingKeeper              slashing.Keeper
		distributionKeeper          distribution.Keeper
		crisisKeeper                crisis.Keeper
		moduleManager               *module.Manager
		BaseApp                     baseapp.BaseApp
	}
	var tests []struct {
		name   string
		fields fields
		want   *codec.Codec
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application := application{
				name:                        tt.fields.name,
				moduleBasicManager:          tt.fields.moduleBasicManager,
				codec:                       tt.fields.codec,
				enabledWasmProposalTypeList: tt.fields.enabledWasmProposalTypeList,
				moduleAccountPermissions:    tt.fields.moduleAccountPermissions,
				tokenReceiveAllowedModules:  tt.fields.tokenReceiveAllowedModules,
				keys:                        tt.fields.keys,
				stakingKeeper:               tt.fields.stakingKeeper,
				slashingKeeper:              tt.fields.slashingKeeper,
				distributionKeeper:          tt.fields.distributionKeeper,
				crisisKeeper:                tt.fields.crisisKeeper,
				moduleManager:               tt.fields.moduleManager,
				BaseApp:                     tt.fields.BaseApp,
			}
			if got := application.GetCodec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCodec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_application_GetDefaultClientHome(t *testing.T) {
	type fields struct {
		name                        string
		moduleBasicManager          module.BasicManager
		codec                       *codec.Codec
		enabledWasmProposalTypeList []wasm.ProposalType
		moduleAccountPermissions    map[string][]string
		tokenReceiveAllowedModules  map[string]bool
		keys                        map[string]*sdkTypes.KVStoreKey
		stakingKeeper               staking.Keeper
		slashingKeeper              slashing.Keeper
		distributionKeeper          distribution.Keeper
		crisisKeeper                crisis.Keeper
		moduleManager               *module.Manager
		BaseApp                     baseapp.BaseApp
	}
	var tests []struct {
		name   string
		fields fields
		want   string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application := application{
				name:                        tt.fields.name,
				moduleBasicManager:          tt.fields.moduleBasicManager,
				codec:                       tt.fields.codec,
				enabledWasmProposalTypeList: tt.fields.enabledWasmProposalTypeList,
				moduleAccountPermissions:    tt.fields.moduleAccountPermissions,
				tokenReceiveAllowedModules:  tt.fields.tokenReceiveAllowedModules,
				keys:                        tt.fields.keys,
				stakingKeeper:               tt.fields.stakingKeeper,
				slashingKeeper:              tt.fields.slashingKeeper,
				distributionKeeper:          tt.fields.distributionKeeper,
				crisisKeeper:                tt.fields.crisisKeeper,
				moduleManager:               tt.fields.moduleManager,
				BaseApp:                     tt.fields.BaseApp,
			}
			if got := application.GetDefaultClientHome(); got != tt.want {
				t.Errorf("GetDefaultClientHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_application_GetDefaultNodeHome(t *testing.T) {
	type fields struct {
		name                        string
		moduleBasicManager          module.BasicManager
		codec                       *codec.Codec
		enabledWasmProposalTypeList []wasm.ProposalType
		moduleAccountPermissions    map[string][]string
		tokenReceiveAllowedModules  map[string]bool
		keys                        map[string]*sdkTypes.KVStoreKey
		stakingKeeper               staking.Keeper
		slashingKeeper              slashing.Keeper
		distributionKeeper          distribution.Keeper
		crisisKeeper                crisis.Keeper
		moduleManager               *module.Manager
		BaseApp                     baseapp.BaseApp
	}
	var tests []struct {
		name   string
		fields fields
		want   string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application := application{
				name:                        tt.fields.name,
				moduleBasicManager:          tt.fields.moduleBasicManager,
				codec:                       tt.fields.codec,
				enabledWasmProposalTypeList: tt.fields.enabledWasmProposalTypeList,
				moduleAccountPermissions:    tt.fields.moduleAccountPermissions,
				tokenReceiveAllowedModules:  tt.fields.tokenReceiveAllowedModules,
				keys:                        tt.fields.keys,
				stakingKeeper:               tt.fields.stakingKeeper,
				slashingKeeper:              tt.fields.slashingKeeper,
				distributionKeeper:          tt.fields.distributionKeeper,
				crisisKeeper:                tt.fields.crisisKeeper,
				moduleManager:               tt.fields.moduleManager,
				BaseApp:                     tt.fields.BaseApp,
			}
			if got := application.GetDefaultNodeHome(); got != tt.want {
				t.Errorf("GetDefaultNodeHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_application_GetModuleBasicManager(t *testing.T) {
	type fields struct {
		name                        string
		moduleBasicManager          module.BasicManager
		codec                       *codec.Codec
		enabledWasmProposalTypeList []wasm.ProposalType
		moduleAccountPermissions    map[string][]string
		tokenReceiveAllowedModules  map[string]bool
		keys                        map[string]*sdkTypes.KVStoreKey
		stakingKeeper               staking.Keeper
		slashingKeeper              slashing.Keeper
		distributionKeeper          distribution.Keeper
		crisisKeeper                crisis.Keeper
		moduleManager               *module.Manager
		BaseApp                     baseapp.BaseApp
	}
	var tests []struct {
		name   string
		fields fields
		want   module.BasicManager
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application := application{
				name:                        tt.fields.name,
				moduleBasicManager:          tt.fields.moduleBasicManager,
				codec:                       tt.fields.codec,
				enabledWasmProposalTypeList: tt.fields.enabledWasmProposalTypeList,
				moduleAccountPermissions:    tt.fields.moduleAccountPermissions,
				tokenReceiveAllowedModules:  tt.fields.tokenReceiveAllowedModules,
				keys:                        tt.fields.keys,
				stakingKeeper:               tt.fields.stakingKeeper,
				slashingKeeper:              tt.fields.slashingKeeper,
				distributionKeeper:          tt.fields.distributionKeeper,
				crisisKeeper:                tt.fields.crisisKeeper,
				moduleManager:               tt.fields.moduleManager,
				BaseApp:                     tt.fields.BaseApp,
			}
			if got := application.GetModuleBasicManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModuleBasicManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_application_Initialize(t *testing.T) {
	type fields struct {
		name                        string
		moduleBasicManager          module.BasicManager
		codec                       *codec.Codec
		enabledWasmProposalTypeList []wasm.ProposalType
		moduleAccountPermissions    map[string][]string
		tokenReceiveAllowedModules  map[string]bool
		keys                        map[string]*sdkTypes.KVStoreKey
		stakingKeeper               staking.Keeper
		slashingKeeper              slashing.Keeper
		distributionKeeper          distribution.Keeper
		crisisKeeper                crisis.Keeper
		moduleManager               *module.Manager
		BaseApp                     baseapp.BaseApp
	}
	type args struct {
		logger             log.Logger
		db                 db.DB
		loadLatest         bool
		invCheckPeriod     uint
		skipUpgradeHeights map[int64]bool
		home               string
		baseAppOptions     []func(*baseapp.BaseApp)
	}
	var tests []struct {
		name           string
		fields         fields
		args           args
		wantTraceStore string
		want           applications.Application
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application := application{
				name:                        tt.fields.name,
				moduleBasicManager:          tt.fields.moduleBasicManager,
				codec:                       tt.fields.codec,
				enabledWasmProposalTypeList: tt.fields.enabledWasmProposalTypeList,
				moduleAccountPermissions:    tt.fields.moduleAccountPermissions,
				tokenReceiveAllowedModules:  tt.fields.tokenReceiveAllowedModules,
				keys:                        tt.fields.keys,
				stakingKeeper:               tt.fields.stakingKeeper,
				slashingKeeper:              tt.fields.slashingKeeper,
				distributionKeeper:          tt.fields.distributionKeeper,
				crisisKeeper:                tt.fields.crisisKeeper,
				moduleManager:               tt.fields.moduleManager,
				BaseApp:                     tt.fields.BaseApp,
			}
			traceStore := &bytes.Buffer{}
			got := application.Initialize(tt.args.logger, tt.args.db, traceStore, tt.args.loadLatest, tt.args.invCheckPeriod, tt.args.skipUpgradeHeights, tt.args.home, tt.args.baseAppOptions...)
			if gotTraceStore := traceStore.String(); gotTraceStore != tt.wantTraceStore {
				t.Errorf("Initialize() gotTraceStore = %v, want %v", gotTraceStore, tt.wantTraceStore)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_application_LoadHeight(t *testing.T) {
	type fields struct {
		name                        string
		moduleBasicManager          module.BasicManager
		codec                       *codec.Codec
		enabledWasmProposalTypeList []wasm.ProposalType
		moduleAccountPermissions    map[string][]string
		tokenReceiveAllowedModules  map[string]bool
		keys                        map[string]*sdkTypes.KVStoreKey
		stakingKeeper               staking.Keeper
		slashingKeeper              slashing.Keeper
		distributionKeeper          distribution.Keeper
		crisisKeeper                crisis.Keeper
		moduleManager               *module.Manager
		BaseApp                     baseapp.BaseApp
	}
	type args struct {
		height int64
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application := application{
				name:                        tt.fields.name,
				moduleBasicManager:          tt.fields.moduleBasicManager,
				codec:                       tt.fields.codec,
				enabledWasmProposalTypeList: tt.fields.enabledWasmProposalTypeList,
				moduleAccountPermissions:    tt.fields.moduleAccountPermissions,
				tokenReceiveAllowedModules:  tt.fields.tokenReceiveAllowedModules,
				keys:                        tt.fields.keys,
				stakingKeeper:               tt.fields.stakingKeeper,
				slashingKeeper:              tt.fields.slashingKeeper,
				distributionKeeper:          tt.fields.distributionKeeper,
				crisisKeeper:                tt.fields.crisisKeeper,
				moduleManager:               tt.fields.moduleManager,
				BaseApp:                     tt.fields.BaseApp,
			}
			if err := application.LoadHeight(tt.args.height); (err != nil) != tt.wantErr {
				t.Errorf("LoadHeight() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_makeCodec(t *testing.T) {
	type args struct {
		moduleBasicManager module.BasicManager
	}
	var tests []struct {
		name string
		args args
		want *codec.Codec
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeCodec(tt.args.moduleBasicManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeCodec() = %v, want %v", got, tt.want)
			}
		})
	}
}
