package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/persistenceOne/persistence-sdk/v4/simapp"
	"github.com/persistenceOne/persistence-sdk/v4/simapp/simd/cmd"
)

func main() {
	cfg := sdk.GetConfig()
	// Configure SDK with proper bech32 prefixes before sealing
	cfg.SetBech32PrefixForAccount("persistence", "persistencepub")
	cfg.SetBech32PrefixForValidator("persistencevaloper", "persistencevaloperpub")
	cfg.SetBech32PrefixForConsensusNode("persistencevalcons", "persistencevalconspub")

	rootCmd := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
