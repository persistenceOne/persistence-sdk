/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"

	//"github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/tendermint/tendermint/crypto"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//type signature struct {
//	ID             types.ID     `json:"id"`
//	SignatureBytes []byte       `json:"signatureBytes"`
//	ValidityHeight types.Height `json:"validityHeight"`
//}

var _ types.Signature = &Signature{}
var _ cdctypes.UnpackInterfacesMessage = &Signature{}

//func (signature *Signature) SetID(a types.ID)error {
//	m, ok := a.(proto.Message)
//	if !ok {
//		return sdkerrors.Wrapf(sdkerrors.ErrPackAny, "can't proto marshal %T", m)
//	}
//	any, err := cdctypes.NewAnyWithValue(m)
//	if err != nil {
//		return err
//	}
//	signature.ID = any
//	return nil
//}

func (baseSignature *Signature) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	var ID types.ID
	var valHeight types.Height
	err := unpacker.UnpackAny(baseSignature.ID, &ID)
	err1 := unpacker.UnpackAny(baseSignature.ValidityHeight, &valHeight)
	if err != nil {
		return err
	} else if err1 != nil {
		return err1
	}
	return nil
}

func (baseSignature *Signature) Bytes() []byte {
	return baseSignature.SignatureBytes
}

func (baseSignature *Signature) GetID() types.ID {
	return baseSignature.ID.GetCachedValue().(types.ID)
}

func (baseSignature *Signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifySignature(bytes, baseSignature.Bytes())
}

func (baseSignature *Signature) GetValidityHeight() types.Height {
	return baseSignature.ValidityHeight.GetCachedValue().(types.Height)
}

func (baseSignature *Signature) HasExpired(height types.Height) bool {
	return baseSignature.GetValidityHeight().IsGreaterThan(height)
}
func NewSignature(id types.ID, signatureBytes []byte, validityHeight types.Height) types.Signature {
	n := Signature{SignatureBytes: signatureBytes}
	_, ok := id.(proto.Message)
	if !ok {
		return nil
	}
	any, err := cdctypes.NewAnyWithValue(id)
	if err != nil {
		return nil
	}
	n.ID = any
	any, err = cdctypes.NewAnyWithValue(validityHeight)
	if err != nil {
		return nil
	}
	n.ValidityHeight = any
	return &n
}

//func (baseSignature signature) String() string {
//	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
//}
//func (baseSignature signature) Bytes() []byte   { return baseSignature.SignatureBytes }
//func (baseSignature signature) GetID() types.ID { return baseSignature.ID }
//func (baseSignature signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
//	return pubKey.VerifyBytes(bytes, baseSignature.Bytes())
//}
//func (baseSignature signature) GetValidityHeight() types.Height {
//	return baseSignature.ValidityHeight
//}
//func (baseSignature signature) HasExpired(height types.Height) bool {
//	return baseSignature.GetValidityHeight().IsGreaterThan(height)
//}

//func NewSignature(id types.ID, signatureBytes []byte, validityHeight types.Height) Signature {
//	return Signature{
//		ID:             id,
//		SignatureBytes: signatureBytes,
//		ValidityHeight: validityHeight,
//	}
//}
