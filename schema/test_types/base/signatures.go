package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

var _ test_types.Signatures = (*signatures)(nil)

func (signatures signatures) Get(id test_types.ID) test_types.Signature {
	for _, signature := range signatures.SignatureList {
		if signature.GetID().Equals(id) {
			return signature
		}
	}

	return nil
}
func (signatures signatures) GetList() []test_types.Signature {
	return signatures.SignatureList
}
func (signatures signatures) Add(signature test_types.Signature) test_types.Signatures {
	signatures.SignatureList = append(signatures.SignatureList, signature)
	return signatures
}
func (signatures signatures) Remove(signature test_types.Signature) test_types.Signatures {
	signatureList := signatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Equals(signature.GetID()) {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}

	return NewSignatures(signatureList)
}
func (signatures signatures) Mutate(signature test_types.Signature) test_types.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Equals(signature.GetID()) {
			signatureList[i] = signature
		}
	}

	return NewSignatures(signatureList)
}
func NewSignatures(signatureList []test_types.Signature) test_types.Signatures {
	return signatures{SignatureList: signatureList}
}
