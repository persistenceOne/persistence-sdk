/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/traits/base"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type assetID struct {
	ClassificationID protoTypes.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	HashID           protoTypes.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ protoTypes.ID = (*assetID)(nil)
var _ helpers.Key = (*assetID)(nil)


func (assetID assetID) MarshalToSizedBuffer(i []byte) (int, error) {
	panic("implement me")
}

//TODO:generate via proto below methods would be generated

func (assetID assetID) Size() int {
	panic("implement me")
}

func (assetID assetID) MarshalTo(i []byte) (int, error) {
	panic("implement me")
}

func (assetID assetID) Unmarshal(i []byte) error {
	panic("implement me")
}


func (assetID assetID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, assetID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, assetID.HashID.Bytes()...)

	return Bytes
}
func (assetID assetID) String() string {
	var values []string
	values = append(values, assetID.ClassificationID.String())
	values = append(values, assetID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (assetID assetID) Equals(id protoTypes.ID) bool {
	return bytes.Equal(assetID.Bytes(), id.Bytes())
}
func (assetID assetID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(assetID.Bytes())
}
func (assetID) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, assetID{})
}
func (assetID assetID) IsPartial() bool {
	return len(assetID.HashID.Bytes()) == 0
}
func (assetID assetID) Matches(key helpers.Key) bool {
	return assetID.Equals(assetIDFromInterface(key))
}

func NewAssetID(classificationID protoTypes.ID, immutableProperties types.Properties) protoTypes.ID {
	return assetID{
		ClassificationID: classificationID,
		HashID:           base.HasImmutables{Properties: immutableProperties}.GenerateHashID(),
	}
}
