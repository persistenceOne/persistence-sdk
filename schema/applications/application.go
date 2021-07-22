/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"io"
)

type Application interface {
	serverTypes.Application
	LoadHeight(int64) error
	ExportApplicationStateAndValidators(bool, []string) (serverTypes.ExportedApp, error)
	Name() string
	AppVersion() string
	Logger() log.Logger
	MountStores(keys ...sdkTypes.StoreKey)
	MountKVStores(keys map[string]*sdkTypes.KVStoreKey)
	MountTransientStores(keys map[string]*sdkTypes.TransientStoreKey)
	MountStoreWithDB(key sdkTypes.StoreKey, typ sdkTypes.StoreType, db tendermintDB.DB)
	MountStore(key sdkTypes.StoreKey, typ sdkTypes.StoreType)
	LoadLatestVersion() error
	LoadVersion(version int64) error
	LastCommitID() sdkTypes.CommitID
	LastBlockHeight() int64
	Router() sdkTypes.Router
	QueryRouter() sdkTypes.QueryRouter
	Seal()
	IsSealed() bool



	Initialize(applicationName string, encodingConfig EncodingConfig,
		logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) Application
	GetCodec() *codec.LegacyAmino
	GetDefaultNodeHome() string
	GetModuleBasicManager() module.BasicManager
}
