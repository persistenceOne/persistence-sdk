package base

import (
	"encoding/base64"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"github.com/tendermint/tendermint/crypto"
)

var _ test_types.Signature = (*signature)(nil)

func (baseSignature signature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature signature) Bytes() []byte   { return baseSignature.SignatureBytes }
func (baseSignature signature) GetID() test_types.ID { return baseSignature.ID }
func (baseSignature signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifySignature(bytes, baseSignature.Bytes())
}
func (baseSignature signature) GetValidityHeight() test_types.Height {
	return baseSignature.ValidityHeight
}
func (baseSignature signature) HasExpired(height test_types.Height) bool {
	return baseSignature.GetValidityHeight().IsGreaterThan(height)
}

func NewSignature(id test_types.ID, signatureBytes []byte, validityHeight test_types.Height) test_types.Signature {
	return signature{
		ID:             id,
		SignatureBytes: signatureBytes,
		ValidityHeight: validityHeight,
	}
}

