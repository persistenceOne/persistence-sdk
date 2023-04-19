package v046

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"
)

// MigrateStore performs in-place store migrations from v0.43/v0.44/v0.45 to v0.46.
// The migration includes:
//
// - Setting the MinCommissionRate, ValidatorBondFactor and GlobalLiquidStakingCap params in the paramstore
// - Initializing TotalLiquidStakedTokens to 0
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, paramstore paramtypes.Subspace) error {
	migrateParamsStore(ctx, paramstore)

	store := ctx.KVStore(storeKey)
	migrateTotalLiquidStakedTokens(store, sdk.ZeroInt())

	return nil
}

// migrateParamsStore migrates the param store values
func migrateParamsStore(ctx sdk.Context, paramstore paramtypes.Subspace) {
	if paramstore.HasKeyTable() {
		paramstore.Set(ctx, types.KeyMinCommissionRate, types.DefaultMinCommissionRate)
		paramstore.Set(ctx, types.KeyValidatorBondFactor, types.DefaultValidatorBondFactor)
		paramstore.Set(ctx, types.KeyGlobalLiquidStakingCap, types.DefaultGlobalLiquidStakingCap)
	} else {
		paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore.Set(ctx, types.KeyMinCommissionRate, types.DefaultMinCommissionRate)
		paramstore.Set(ctx, types.KeyValidatorBondFactor, types.DefaultValidatorBondFactor)
		paramstore.Set(ctx, types.KeyGlobalLiquidStakingCap, types.DefaultGlobalLiquidStakingCap)
	}
}

// migrateTotalLiquidStakedTokens migrates the total outstanding tokens owned by a liquid staking provider
func migrateTotalLiquidStakedTokens(store sdk.KVStore, tokens sdk.Int) {
	tokensBz, err := tokens.Marshal()
	if err != nil {
		panic(err)
	}

	store.Set(types.TotalLiquidStakedTokensKey, tokensBz)
}
