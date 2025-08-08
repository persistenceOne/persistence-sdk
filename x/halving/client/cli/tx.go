/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceCore contributors
 SPDX-License-Identifier: Apache-2.0
*/

package cli

import (
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/persistenceOne/persistence-sdk/v4/x/halving/types"
)

// GetTxCmd returns the transaction commands for the halving module.
func GetTxCmd() *cobra.Command {
	halvingTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Halving transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	halvingTxCmd.AddCommand(
		GetCmdUpdateParams(),
	)

	return halvingTxCmd
}

// GetCmdUpdateParams implements a command to update halving parameters.
func GetCmdUpdateParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-params [block-height]",
		Short: "Update halving parameters",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse block height
			blockHeight, ok := math.NewIntFromString(args[0])
			if !ok {
				return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid block height: %s", args[0])
			}

			// Create params
			params := types.NewParams(blockHeight.Uint64())

			// Create message
			msg := types.NewMsgUpdateParams(
				clientCtx.GetFromAddress().String(),
				params,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
