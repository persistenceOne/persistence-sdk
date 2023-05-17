package types

import (
	"fmt"

	"cosmossdk.io/errors"
	"github.com/cometbft/cometbft/crypto/tmhash"
)

// Oracle sentinel errors
var (
	ErrInvalidOraclePrice  = errors.Register(ModuleName, 1112, "invalid oracle price")
	ErrInvalidExchangeRate = errors.Register(ModuleName, 1, "invalid exchange rate")
	ErrNoVotingPermission  = errors.Register(ModuleName, 4, "unauthorized voter")
	ErrInvalidHash         = errors.Register(ModuleName, 5, "invalid hash")
	ErrInvalidHashLength   = errors.Register(ModuleName, 6,
		fmt.Sprintf("invalid hash length; should equal %d", tmhash.TruncatedSize))
	ErrVerificationFailed    = errors.Register(ModuleName, 7, "hash verification failed")
	ErrRevealPeriodMissMatch = errors.Register(ModuleName, 8,
		"reveal period of submitted vote does not match with registered prevote")
	ErrInvalidSaltLength  = errors.Register(ModuleName, 9, "invalid salt length; must be 64")
	ErrInvalidSaltFormat  = errors.Register(ModuleName, 10, "invalid salt format")
	ErrNoAggregatePrevote = errors.Register(ModuleName, 11, "no aggregate prevote")
	ErrNoAggregateVote    = errors.Register(ModuleName, 12, "no aggregate vote")
	ErrUnknownDenom       = errors.Register(ModuleName, 13, "unknown denom")
	ErrExistingPrevote    = errors.Register(ModuleName, 15, "prevote already submitted for this voting period")
	ErrBallotNotSorted    = errors.Register(ModuleName, 16, "ballot must be sorted before this operation")
)
