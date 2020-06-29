package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type QueryKeeper interface {
	Query(sdkTypes.Context, QueryRequest) ([]byte, error)
}
