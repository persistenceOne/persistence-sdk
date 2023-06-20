/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceCore contributors
 SPDX-License-Identifier: Apache-2.0
*/

package halving

import (
	"fmt"
	"strconv"

	errors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/persistenceOne/persistence-sdk/v2/x/halving/types"
)

func EndBlocker(ctx sdk.Context, k Keeper) {
	params := k.GetParams(ctx)

	if params.BlockHeight != 0 && uint64(ctx.BlockHeight())%params.BlockHeight == 0 {
		mintParams := k.GetMintingParams(ctx)
		newMaxInflation := mintParams.InflationMax.QuoTruncate(sdk.NewDecFromInt(Factor))
		newMinInflation := mintParams.InflationMin.QuoTruncate(sdk.NewDecFromInt(Factor))

		if newMaxInflation.Sub(newMinInflation).LT(sdk.ZeroDec()) {
			panic(fmt.Sprintf("max inflation (%s) must be greater than or equal to min inflation (%s)", newMaxInflation.String(), newMinInflation.String()))
		}

		updatedParams := minttypes.NewParams(mintParams.MintDenom, mintParams.InflationRateChange, newMaxInflation, newMinInflation, mintParams.GoalBonded, mintParams.BlocksPerYear)

		if err := k.SetMintingParams(ctx, updatedParams); err != nil {
			panic(errors.Wrap(err, "unable to set minting params at halving EndBlocker"))
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeHalving,
				sdk.NewAttribute(types.AttributeKeyBlockHeight, strconv.FormatInt(ctx.BlockHeight(), 10)),
				sdk.NewAttribute(types.AttributeKeyNewInflationMax, updatedParams.InflationMax.String()),
				sdk.NewAttribute(types.AttributeKeyNewInflationMin, updatedParams.InflationMin.String()),
				sdk.NewAttribute(types.AttributeKeyNewInflationRateChange, updatedParams.InflationRateChange.String()),
			),
		)
	}
}
