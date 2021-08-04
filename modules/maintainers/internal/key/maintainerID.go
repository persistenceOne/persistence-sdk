/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type maintainerID struct {
	ClassificationID test_types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	IdentityID       test_types.ID `json:"identityID" valid:"required~required field identityID missing"`
}


func (maintainerID maintainerID) Size() int {
	panic("implement me")
}

func (maintainerID maintainerID) MarshalTo(i []byte) (int, error) {
	panic("implement me")
}

func (maintainerID maintainerID) Unmarshal(i []byte) error {
	panic("implement me")
}

func (maintainerID maintainerID) MarshalToSizedBuffer(i []byte) (int, error) {
	panic("implement me")
}

var _ test_types.ID = (*maintainerID)(nil)
var _ helpers.Key = (*maintainerID)(nil)

func (maintainerID maintainerID) Bytes() []byte {
	return append(
		maintainerID.ClassificationID.Bytes(),
		maintainerID.IdentityID.Bytes()...)
}
func (maintainerID maintainerID) String() string {
	var values []string
	values = append(values, maintainerID.ClassificationID.String())
	values = append(values, maintainerID.IdentityID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (maintainerID maintainerID) Equals(id test_types.ID) bool {
	return bytes.Equal(maintainerID.Bytes(), id.Bytes())
}
func (maintainerID maintainerID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(maintainerID.Bytes())
}
func (maintainerID) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, maintainerID{})
}
func (maintainerID maintainerID) IsPartial() bool {
	return len(maintainerID.IdentityID.Bytes()) == 0
}
func (maintainerID maintainerID) Matches(key helpers.Key) bool {
	return maintainerID.Equals(maintainerIDFromInterface(key))
}

func NewMaintainerID(classificationID test_types.ID, identityID test_types.ID) test_types.ID {
	return maintainerID{
		ClassificationID: classificationID,
		IdentityID:       identityID,
	}
}
