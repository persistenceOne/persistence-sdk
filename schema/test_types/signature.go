package test_types

import "github.com/tendermint/tendermint/crypto"

type Signature interface {
	String() string
	Bytes() []byte

	GetID() ID

	Verify(crypto.PubKey, []byte) bool
	GetValidityHeight() Height
	HasExpired(Height) bool

	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
	MarshalToSizedBuffer([]byte) (int, error)
	Size() int
}
