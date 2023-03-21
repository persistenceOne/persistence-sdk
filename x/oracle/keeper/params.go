package keeper

import (
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetVotePeriod returns the number of blocks during which voting takes place.
func (k Keeper) GetVotePeriod(ctx sdk.Context) (res uint64) {
	k.paramSpace.Get(ctx, types.KeyVotePeriod, &res)
	return
}

// GetVoteThreshold returns the minimum percentage of votes that must be received
// for a ballot to pass.
func (k Keeper) GetVoteThreshold(ctx sdk.Context) (res sdk.Dec) {
	k.paramSpace.Get(ctx, types.KeyVoteThreshold, &res)
	return
}

// SetVoteThreshold sets min combined validator power voting on a denom to accept
// it as valid.
// TODO: this is used in tests, we should refactor the way how this is handled.
func (k Keeper) SetVoteThreshold(ctx sdk.Context, threshold sdk.Dec) error {
	if err := types.ValidateVoteThreshold(threshold); err != nil {
		return err
	}

	k.paramSpace.Set(ctx, types.KeyVoteThreshold, &threshold)

	return nil
}

// GetRewardDistributionWindow returns the number of vote periods during which
// seigniorage reward comes in and then is distributed.
func (k Keeper) GetRewardDistributionWindow(ctx sdk.Context) (res uint64) {
	k.paramSpace.Get(ctx, types.KeyRewardDistributionWindow, &res)
	return
}

// GetSlashFraction returns oracle voting penalty rate
func (k Keeper) GetSlashFraction(ctx sdk.Context) (res sdk.Dec) {
	k.paramSpace.Get(ctx, types.KeySlashFraction, &res)
	return
}

// GetSlashWindow returns # of vote period for oracle slashing
func (k Keeper) GetSlashWindow(ctx sdk.Context) (res uint64) {
	k.paramSpace.Get(ctx, types.KeySlashWindow, &res)
	return
}

// GetMinValidPerWindow returns oracle slashing threshold
func (k Keeper) GetMinValidPerWindow(ctx sdk.Context) (res sdk.Dec) {
	k.paramSpace.Get(ctx, types.KeyMinValidPerWindow, &res)
	return
}

// GetParams returns the total set of oracle parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// SetParams sets the total set of oracle parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetAcceptList returns the denom list that can be activated
func (k Keeper) GetAcceptList(ctx sdk.Context) (res types.DenomList) {
	k.paramSpace.Get(ctx, types.KeyAcceptList, &res)
	return
}

// SetAcceptList updates the accepted list of assets supported by the x/oracle
// module.
func (k Keeper) SetAcceptList(ctx sdk.Context, acceptList types.DenomList) {
	k.paramSpace.Set(ctx, types.KeyAcceptList, acceptList)
}
