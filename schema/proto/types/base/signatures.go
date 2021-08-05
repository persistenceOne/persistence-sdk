package base

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
)

var _ protoTypes.Signatures = (*signatures)(nil)

func (signatures signatures) Get(id protoTypes.ID) protoTypes.Signature {
	for _, signature := range signatures.SignatureList {
		if signature.GetID().Equals(id) {
			return signature
		}
	}

	return nil
}
func (signatures signatures) GetList() []protoTypes.Signature {
	return signatures.SignatureList
}
func (signatures signatures) Add(signature protoTypes.Signature) protoTypes.Signatures {
	signatures.SignatureList = append(signatures.SignatureList, signature)
	return signatures
}
func (signatures signatures) Remove(signature protoTypes.Signature) protoTypes.Signatures {
	signatureList := signatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Equals(signature.GetID()) {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}

	return NewSignatures(signatureList)
}
func (signatures signatures) Mutate(signature protoTypes.Signature) protoTypes.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Equals(signature.GetID()) {
			signatureList[i] = signature
		}
	}

	return NewSignatures(signatureList)
}
func NewSignatures(signatureList []protoTypes.Signature) protoTypes.Signatures {
	return signatures{SignatureList: signatureList}
}
