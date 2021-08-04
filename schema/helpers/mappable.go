/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

type Mappable interface {
	GetKey() Key
	RegisterCodec(protoCodec *codec.LegacyAmino)
	test_types.ProtoInterface

}
