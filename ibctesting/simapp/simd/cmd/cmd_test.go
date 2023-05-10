package cmd_test

import (
	"fmt"
	"testing"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistence-sdk/v2/ibctesting/simapp"
	"github.com/persistenceOne/persistence-sdk/v2/ibctesting/simapp/simd/cmd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",        // Test the init cmd
		"simapp-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
	})

	require.NoError(t, svrcmd.Execute(rootCmd, "simd", simapp.DefaultNodeHome))
}
