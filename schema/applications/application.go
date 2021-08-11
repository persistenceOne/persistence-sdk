/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"io"
)

type Application interface {
	serverTypes.Application
	LoadHeight(int64) error
	ExportApplicationStateAndValidators(bool, []string) (serverTypes.ExportedApp, error)

	Initialize(applicationName string, encodingConfig EncodingConfig, enabledProposals []wasm.ProposalType, moduleAccountPermissions map[string][]string, tokenReceiveAllowedModules map[string]bool,
		logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, applicationOptions serverTypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp)) Application
}

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}
