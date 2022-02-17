package types

const (
	// ModuleName defines the module name
	ModuleName = "interchainquery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

// prefix bytes for the interchainquery persistent store
const (
	prefixData          = iota + 1
	prefixPeriodicQuery = iota + 1
	prefixSingleQuery   = iota + 1
)

var (
	KeyPrefixData          = []byte{prefixData}
	KeyPrefixPeriodicQuery = []byte{prefixPeriodicQuery}
	KeyPrefixSingleQuery   = []byte{prefixSingleQuery}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
