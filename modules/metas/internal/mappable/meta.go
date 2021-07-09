/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type meta struct {
	ID   types.ID   `json:"id" valid:"required field id missing"`
	Data types.Data `json:"data" valid:"required field data missing"`
}

func (meta meta) Reset() {
	panic("implement me")
}

func (meta meta) String() string {
	panic("implement me")
}

func (meta meta) ProtoMessage() {
	panic("implement me")
}

func (meta meta) Marshal() ([]byte, error) {
	panic("implement me")
}

func (meta meta) MarshalTo(data []byte) (n int, err error) {
	panic("implement me")
}

func (meta meta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	panic("implement me")
}

func (meta meta) Size() int {
	panic("implement me")
}

func (meta meta) Unmarshal(data []byte) error {
	panic("implement me")
}

var _ mappables.Meta = (*meta)(nil)

func (meta meta) GetData() types.Data { return meta.Data }
func (meta meta) GetID() types.ID     { return meta.ID }
func (meta meta) GetKey() helpers.Key {
	return key.FromID(meta.GetID())
}
func (meta) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, meta{})
}

func NewMeta(data types.Data) mappables.Meta {
	return meta{
		ID:   key.GenerateMetaID(data),
		Data: data,
	}
}
