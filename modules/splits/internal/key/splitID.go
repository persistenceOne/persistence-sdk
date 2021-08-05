/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type splitID struct {
	OwnerID   protoTypes.ID `json:"ownerID" valid:"required~required field ownerID missing"`
	OwnableID protoTypes.ID `json:"ownableID" valid:"required~required field ownableID missing"`
}

func (splitID splitID) MarshalToSizedBuffer(i []byte) (int, error) {
	panic("implement me")
}

//TODO: would be generated via proto file
func (splitID splitID) Size() int {
	panic("implement me")
}

func (splitID splitID) MarshalTo(i []byte) (int, error) {
	panic("implement me")
}

func (splitID splitID) Unmarshal(i []byte) error {
	panic("implement me")
}

var _ protoTypes.ID = (*splitID)(nil)
var _ helpers.Key = (*splitID)(nil)

func (splitID splitID) Bytes() []byte {
	return append(
		splitID.OwnerID.Bytes(),
		splitID.OwnableID.Bytes()...)
}
func (splitID splitID) String() string {
	var values []string
	values = append(values, splitID.OwnerID.String())
	values = append(values, splitID.OwnableID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (splitID splitID) Equals(id protoTypes.ID) bool {
	return bytes.Equal(splitID.Bytes(), id.Bytes())
}
func (splitID splitID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(splitID.Bytes())
}
func (splitID) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, splitID{})
}
func (splitID splitID) IsPartial() bool {
	return len(splitID.OwnableID.Bytes()) == 0
}
func (splitID splitID) Matches(key helpers.Key) bool {
	return splitID.Equals(splitIDFromInterface(key))
}

func NewSplitID(ownerID protoTypes.ID, ownableID protoTypes.ID) protoTypes.ID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
