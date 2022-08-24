package keeper

import "github.com/persistenceOne/persistence-sdk/x/ibchooker/types"

type Keeper struct {
	hooks types.IBCTransferHooks
}

func NewKeeper() Keeper {
	return Keeper{
		hooks: nil,
	}
}

// Set the validator hooks
func (k *Keeper) SetHooks(transferHooks types.IBCTransferHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set hooks twice")
	}

	k.hooks = transferHooks

	return k
}
