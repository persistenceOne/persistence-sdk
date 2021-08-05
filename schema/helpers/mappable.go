/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
)

type Mappable interface {
	GetKey() Key
	RegisterCodec(protoCodec *codec.LegacyAmino)
	protoTypes.ProtoInterface

}
