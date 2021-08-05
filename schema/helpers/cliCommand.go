/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/spf13/cobra"
)

type CLICommand interface {
	ReadInt64(CLIFlag) int64
	ReadInt(CLIFlag) int
	ReadBool(CLIFlag) bool
	ReadString(CLIFlag) string
	ReadBaseReq(ctx client.Context) protoTypes.BaseReq

	CreateCommand(func(command *cobra.Command, args []string) error) *cobra.Command
}
