/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/99designs/keyring"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

//type fact struct {
//	HashID     types.ID         `json:"hashID"`
//	TypeID     types.ID         `json:"typeID"`
//	Signatures types.Signatures `json:"signatures"`
//}

var _ types.Fact = &Fact{}
var _ cdctypes.UnpackInterfacesMessage = &Fact{}

func (fact *Fact) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	var hashID types.ID
	var typeID types.ID
	var signature types.Signatures
	err := unpacker.UnpackAny(fact.HashID, &hashID)
	err1 := unpacker.UnpackAny(fact.TypeID, &typeID)
	err2 := unpacker.UnpackAny(fact.Signatures, &signature)
	if err != nil {
		return err
	} else if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	}
	return nil
}

func (fact *Fact) GetHashID() types.ID { return fact.HashID.GetCachedValue().(types.ID) }
func (fact *Fact) GetTypeID() types.ID { return fact.TypeID.GetCachedValue().(types.ID) }
func (fact *Fact) GetSignatures() types.Signatures {
	return fact.Signatures.GetCachedValue().(types.Signatures)
}
func (fact *Fact) IsMeta() bool {
	return false
}
func (fact *Fact) Sign(_ keyring.Keyring) types.Fact {
	clicont := client.Context{}
	sign, _, _ := clicont.Keybase.Sign(clicont.FromName, keys.DefaultKeyPass, fact.GetHashID().Bytes())
	signature := Signature{
		ID:             Id{IDstring: fact.GetHashID().String()},
		SignatureBytes: sign,
		ValidityHeight: Height{Value: clicont.Height},
	}
	fact.GetSignatures().Add(signature)

	return fact
}

func NewFact(data types.Data) types.Fact {
	_, ok := data.(proto.Message)
	if !ok {
		return &Fact{}
	}
	any, err := cdctypes.NewAnyWithValue(data.GenerateHashID())
	if err != nil {
		return &Fact{}
	}
	any1, err := cdctypes.NewAnyWithValue(data.GetTypeID())
	if err != nil {
		return &Fact{}
	}
	return &Fact{
		HashID:     any,
		TypeID:     any1,
		Signatures: signatures{},
	}
}

func ReadFact(metaFactString string) (types.Fact, error) {
	metaFact, Error := ReadMetaFact(metaFactString)
	if Error != nil {
		return nil, Error
	}

	return metaFact.RemoveData(), nil
}
