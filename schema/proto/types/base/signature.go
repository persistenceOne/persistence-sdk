package base

import (
	"encoding/base64"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/tendermint/tendermint/crypto"
)

var _ protoTypes.Signature = (*signature)(nil)

func (baseSignature signature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature signature) Bytes() []byte        { return baseSignature.SignatureBytes }
func (baseSignature signature) GetID() protoTypes.ID { return baseSignature.ID }
func (baseSignature signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifySignature(bytes, baseSignature.Bytes())
}
func (baseSignature signature) GetValidityHeight() protoTypes.Height {
	return baseSignature.ValidityHeight
}
func (baseSignature signature) HasExpired(height protoTypes.Height) bool {
	return baseSignature.GetValidityHeight().IsGreaterThan(height)
}

func NewSignature(id protoTypes.ID, signatureBytes []byte, validityHeight protoTypes.Height) protoTypes.Signature {
	return signature{
		ID:             id,
		SignatureBytes: signatureBytes,
		ValidityHeight: validityHeight,
	}
}

