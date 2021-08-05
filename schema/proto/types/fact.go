package types

import "github.com/99designs/keyring"

type Fact interface {
	GetHashID() ID
	GetTypeID() ID
	GetSignatures() Signatures

	Sign(keyring.Keyring) Fact
	ProtoInterface
}
