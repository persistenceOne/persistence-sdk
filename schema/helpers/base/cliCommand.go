/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"fmt"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/spf13/cobra"
)

type CliCommand struct {
	Use         string
	Short       string
	long        string
	cliFlagList []helpers.CLIFlag
}

var _ helpers.CLICommand = (*CliCommand)(nil)

func (cliCommand CliCommand) registerFlags(command *cobra.Command) {
	for _, cliFlag := range cliCommand.cliFlagList {
		cliFlag.Register(command)
	}
}

func (cliCommand CliCommand) ReadInt64(cliFlag helpers.CLIFlag) int64 {
	switch cliFlag.GetValue().(type) {
	case int64:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(int64)
			}
		}
	default:
		panic(fmt.Errorf("flag %v not an int64 flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue()))
	}
	panic(fmt.Errorf("uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue()))
}

func (cliCommand CliCommand) ReadInt(cliFlag helpers.CLIFlag) int {
	switch cliFlag.GetValue().(type) {
	case int:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(int)
			}
		}
	default:
		panic(fmt.Errorf("flag %v not an int flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue()))
	}
	panic(fmt.Errorf("uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue()))
}

func (cliCommand CliCommand) ReadBool(cliFlag helpers.CLIFlag) bool {
	switch cliFlag.GetValue().(type) {
	case bool:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(bool)
			}
		}
	default:
		panic(fmt.Errorf("flag %v not an bool flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue()))
	}
	panic(fmt.Errorf("uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue()))
}

func (cliCommand CliCommand) ReadString(cliFlag helpers.CLIFlag) string {
	switch cliFlag.GetValue().(type) {
	case string:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(string)
			}
		}
	default:
		panic(fmt.Errorf("falg %v not an string flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue()))
	}
	panic(fmt.Errorf("uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue()))
}

func (cliCommand CliCommand) ReadBaseReq(cliContext client.Context) test_types.BaseReq {
	return test_types.BaseReq{
		From:     cliContext.GetFromAddress().String(),
		ChainId:  cliContext.ChainID,
		Simulate: cliContext.Simulate,
	}
}
func (cliCommand CliCommand) CreateCommand(runE func(command *cobra.Command, args []string) error) *cobra.Command {
	command := &cobra.Command{
		Use:   cliCommand.Use,
		Short: cliCommand.Short,
		Long:  cliCommand.long,
		RunE:  runE,
	}
	cliCommand.registerFlags(command)

	flags.AddTxFlagsToCmd(command)

	return command
}

func NewCLICommand(use string, short string, long string, cliFlagList []helpers.CLIFlag) helpers.CLICommand {
	return CliCommand{
		Use:         use,
		Short:       short,
		long:        long,
		cliFlagList: cliFlagList,
	}
}
