package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/x/genaccounts"

	"github.com/commitHub/commitBlockchain/applications/hub"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tendermintABSITypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/commitHub/commitBlockchain/applications/hub/initialize"
)

const flagInvalidCheckPeriod = "invalid-check-period"

var invalidCheckPeriod uint

func main() {
	cdc := hub.MakeCodec()

	configuration := sdkTypes.GetConfig()
	configuration.SetBech32PrefixForAccount(sdkTypes.Bech32PrefixAccAddr, sdkTypes.Bech32PrefixAccPub)
	configuration.SetBech32PrefixForValidator(sdkTypes.Bech32PrefixValAddr, sdkTypes.Bech32PrefixValPub)
	configuration.SetBech32PrefixForConsensusNode(sdkTypes.Bech32PrefixConsAddr, sdkTypes.Bech32PrefixConsPub)
	configuration.Seal()

	serverContext := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCommand := &cobra.Command{
		Use:               "hubNode",
		Short:             "Commit Hub Node Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(serverContext),
	}

	rootCommand.AddCommand(initialize.InitializeCommand(serverContext, cdc, hub.ModuleBasics, hub.DefaultNodeHome))
	rootCommand.AddCommand(initialize.CollectGenesisTransactionsCommand(serverContext, cdc, genaccounts.AppModuleBasic{}, hub.DefaultNodeHome))
	rootCommand.AddCommand(initialize.MigrateGenesisCommand(serverContext, cdc))
	rootCommand.AddCommand(initialize.GenesisTransactionCommand(serverContext, cdc, hub.ModuleBasics, staking.AppModuleBasic{}, genaccounts.AppModuleBasic{}, hub.DefaultNodeHome, hub.DefaultClientHome))
	rootCommand.AddCommand(initialize.ValidateGenesisCommand(serverContext, cdc, hub.ModuleBasics))
	rootCommand.AddCommand(initialize.AddGenesisAccountCommand(serverContext, cdc, hub.DefaultNodeHome, hub.DefaultClientHome))
	rootCommand.AddCommand(client.NewCompletionCmd(rootCommand, true))
	rootCommand.AddCommand(initialize.TestnetCommand(serverContext, cdc, hub.ModuleBasics, genaccounts.AppModuleBasic{}))

	rootCommand.AddCommand(initialize.ReplayCmd())

	server.AddCommands(serverContext, cdc, rootCommand, newApplication, exportApplicationStateAndValidators)

	executor := cli.PrepareBaseCmd(rootCommand, "CA", hub.DefaultNodeHome)
	rootCommand.PersistentFlags().UintVar(
		&invalidCheckPeriod,
		flagInvalidCheckPeriod,
		0,
		"Assert registered invariants every N blocks",
	)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApplication(logger log.Logger, db tendermintDB.DB, traceStore io.Writer) tendermintABSITypes.Application {
	return hub.NewCommitHubApplication(
		logger,
		db,
		traceStore,
		true,
		invalidCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportApplicationStateAndValidators(
	logger log.Logger,
	db tendermintDB.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailWhiteList []string,
) (json.RawMessage, []tendermintTypes.GenesisValidator, error) {

	if height != -1 {
		genesisApplication := hub.NewCommitHubApplication(logger, db, traceStore, false, uint(1))
		err := genesisApplication.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return genesisApplication.ExportApplicationStateAndValidators(forZeroHeight, jailWhiteList)
	} else {
		genesisApplication := hub.NewCommitHubApplication(logger, db, traceStore, true, uint(1))
		return genesisApplication.ExportApplicationStateAndValidators(forZeroHeight, jailWhiteList)
	}
}
