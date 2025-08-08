package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/persistenceOne/persistence-sdk/v4/simapp"
	"github.com/persistenceOne/persistence-sdk/v4/simapp/simd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
