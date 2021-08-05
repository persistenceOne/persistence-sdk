/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"strings"

	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type identityID struct {
	ClassificationID protoTypes.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	HashID           protoTypes.ID `json:"hashID" valid:"required~required field hashID missing"`
}


func (identityID identityID) Size() int {
	panic("implement me")
}

func (identityID identityID) MarshalTo(i []byte) (int, error) {
	panic("implement me")
}

func (identityID identityID) Unmarshal(i []byte) error {
	panic("implement me")
}

func (identityID identityID) MarshalToSizedBuffer(i []byte) (int, error) {
	panic("implement me")
}

var _ protoTypes.ID = (*identityID)(nil)
var _ helpers.Key = (*identityID)(nil)

func (identityID identityID) Bytes() []byte {
	return append(
		identityID.ClassificationID.Bytes(),
		identityID.HashID.Bytes()...,
	)
}
func (identityID identityID) String() string {
	var values []string
	values = append(values, identityID.ClassificationID.String())
	values = append(values, identityID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (identityID identityID) Equals(id protoTypes.ID) bool {
	return bytes.Equal(identityID.Bytes(), id.Bytes())
}
func (identityID identityID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(identityID.Bytes())
}
func (identityID) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, identityID{})
}
func (identityID identityID) IsPartial() bool {
	return len(identityID.HashID.Bytes()) == 0
}
func (identityID identityID) Matches(key helpers.Key) bool {
	return identityID.Equals(identityIDFromInterface(key))
}

func NewIdentityID(classificationID protoTypes.ID, immutableProperties types.Properties) protoTypes.ID {
	return identityID{
		ClassificationID: classificationID,
		HashID:           baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID(),
	}
}
