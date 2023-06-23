package keepers

import (
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	epochstypes "github.com/persistenceOne/persistence-sdk/v2/x/epochs/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/halving"
	ibchookstypes "github.com/persistenceOne/persistence-sdk/v2/x/ibc-hooks/types"
	interchainquerytypes "github.com/persistenceOne/persistence-sdk/v2/x/interchainquery/types"
	oracletypes "github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
	buildertypes "github.com/skip-mev/pob/x/builder/types"
	routertypes "github.com/strangelove-ventures/packet-forward-middleware/v7/router/types"
)

func (appKeepers *AppKeepers) GenerateKeys() {
	// Define what keys will be used in the cosmos-sdk key/value store.
	// Cosmos-SDK modules each have a "key" that allows the application to reference what they've stored on the chain.
	appKeepers.keys = sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		authzkeeper.StoreKey,
		banktypes.StoreKey,
		capabilitytypes.StoreKey,
		consensustypes.StoreKey,
		crisistypes.StoreKey,
		distributiontypes.StoreKey,
		epochstypes.StoreKey,
		evidencetypes.StoreKey,
		feegrant.StoreKey,
		govtypes.StoreKey,
		group.StoreKey,
		halving.StoreKey,
		ibcexported.StoreKey,
		ibcfeetypes.StoreKey,
		ibchookstypes.StoreKey,
		ibctransfertypes.StoreKey,
		icacontrollertypes.StoreKey,
		icahosttypes.StoreKey,
		interchainquerytypes.StoreKey,
		minttypes.StoreKey,
		oracletypes.StoreKey,
		paramstypes.StoreKey,
		routertypes.StoreKey,
		slashingtypes.StoreKey,
		stakingtypes.StoreKey,
		upgradetypes.StoreKey,
		wasm.StoreKey,
		buildertypes.StoreKey,
	)

	// Define transient store keys
	appKeepers.tkeys = sdk.NewTransientStoreKeys(paramstypes.TStoreKey)

	// MemKeys are for information that is stored only in RAM.
	appKeepers.memKeys = sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
}

func (appKeepers *AppKeepers) GetKVStoreKey() map[string]*storetypes.KVStoreKey {
	return appKeepers.keys
}

func (appKeepers *AppKeepers) GetTransientStoreKey() map[string]*storetypes.TransientStoreKey {
	return appKeepers.tkeys
}

func (appKeepers *AppKeepers) GetMemoryStoreKey() map[string]*storetypes.MemoryStoreKey {
	return appKeepers.memKeys
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetKey(storeKey string) *storetypes.KVStoreKey {
	key := appKeepers.keys[storeKey]
	if key == nil {
		panic(fmt.Sprintf("expected kv store key at %s", storeKey))
	}

	return key
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	key := appKeepers.tkeys[storeKey]
	if key == nil {
		panic(fmt.Sprintf("expected kv store tkey at %s", storeKey))
	}

	return key
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (appKeepers *AppKeepers) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	key := appKeepers.memKeys[storeKey]
	if key == nil {
		panic(fmt.Sprintf("expected kv store memkey at %s", storeKey))
	}

	return key
}
