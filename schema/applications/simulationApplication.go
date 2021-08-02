/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	"testing"
)

type SimulationApplication interface {
	Application
	simapp.App

	GetBaseApp() *baseapp.BaseApp
	GetKey(storeKey string) *sdk.KVStoreKey
	GetTKey(storeKey string) *sdk.TransientStoreKey
	GetSubspace(moduleName string) params.Subspace
	GetModuleAccountPermissions() map[string][]string
	GetBlackListedAddresses() map[string]bool
	ModuleManager() *module.Manager

	CheckBalance(*testing.T, sdk.AccAddress, sdk.Coins)

	AddTestAddresses(sdk.Context, int, sdk.Int) []sdk.AccAddress

	Setup(bool) SimulationApplication
	SetupWithGenesisAccounts([]exported.GenesisAccount) SimulationApplication
	NewTestApplication(bool) (SimulationApplication, sdk.Context)
}
