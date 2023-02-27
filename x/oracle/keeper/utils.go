package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// contains returns true if x is in ls
func contains[T comparable](x T, ls []T) bool {
	for i := range ls {
		if ls[i] == x {
			return true
		}
	}

	return false
}

// get sdk.ValAddress from iterator key by removing prefix
func getValAddrFromIteratorKey(key []byte) sdk.ValAddress {
	return sdk.ValAddress(key[2:])
}
