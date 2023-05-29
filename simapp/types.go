package simapp

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// App implements the common methods for a Cosmos SDK-based application
// specific blockchain.
type App interface {
	Name() string
	BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock
	EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock
	Configurator() module.Configurator
	InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain
	LoadHeight(height int64) error
	LegacyAmino() *codec.LegacyAmino
	AppCodec() codec.Codec
	InterfaceRegistry() types.InterfaceRegistry
	TxConfig() client.TxConfig
	DefaultGenesis() map[string]json.RawMessage
	SimulationManager() *module.SimulationManager
}

var _ App = &SimApp{}
