package types

const (
	// ModuleName defines the module name
	ModuleName = "interchainquery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName
)

// prefix bytes for the interchainquery persistent store
const (
	prefixData  = iota + 1
	prefixQuery = iota + 1
)

var (
	KeyPrefixData  = []byte{prefixData}
	KeyPrefixQuery = []byte{prefixQuery}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
