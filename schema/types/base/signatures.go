/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

//type signatures struct {
//	SignatureList []types.Signature `json:"signatureList"`
//}

var _ types.Signatures = &Signatures{}

func (signatures *Signatures) Get(id types.ID) types.Signature {
	for _, signature := range signatures.GetList() {
		if signature.GetID().Equals(id) {
			return signature
		}
	}

	return nil
}
func (signatures *Signatures) GetList() []types.Signature {
	a := make([]types.Signature, len(signatures.SignatureList))
	for i, listany := range signatures.SignatureList {
		lis, ok := listany.GetCachedValue().(types.Signature)
		if !ok {
			return nil
		}
		a[i] = lis
	}
	return a
}
func (signatures *Signatures) Add(signature types.Signature) types.Signatures {
	signatures.GetList() = append(signatures.GetList(), signature)
	return signatures
}
func (signatures *Signatures) Remove(signature types.Signature) types.Signatures {
	signatureList := signatures.SignatureList
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Equals(signature.GetID()) {
			signatureList = append(signatureList[:i], signatureList[i+1:]...)
		}
	}

	return NewSignatures(signatureList)
}
func (signatures *Signatures) Mutate(signature types.Signature) types.Signatures {
	signatureList := signatures.GetList()
	for i, oldSignature := range signatureList {
		if oldSignature.GetID().Equals(signature.GetID()) {
			signatureList[i] = signature
		}
	}

	return NewSignatures(signatureList)
}
func NewSignatures(signatureList []types.Signature) types.Signatures {
	return signatures{SignatureList: signatureList}
}
