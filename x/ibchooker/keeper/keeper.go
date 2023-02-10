package keeper

import "github.com/persistenceOne/persistence-sdk/v2/x/ibchooker/types"

type Keeper struct {
	hooks types.IBCHandshakeHooks
}

func NewKeeper() Keeper {
	return Keeper{
		hooks: nil,
	}
}

// Set the validator hooks
func (k *Keeper) SetHooks(transferHooks types.IBCHandshakeHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set hooks twice")
	}

	k.hooks = transferHooks

	return k
}
