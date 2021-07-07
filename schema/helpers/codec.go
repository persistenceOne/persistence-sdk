/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)


func RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codec.RegisterInterface((*Mappable)(nil), nil)
	codec.RegisterInterface((*QueryRequest)(nil), nil)
	codec.RegisterInterface((*QueryResponse)(nil), nil)
	codec.RegisterInterface((*TransactionRequest)(nil), nil)
}


//func RegisterCodec(codec *codec.ProtoCodec) {
//	//TODO: Register Interface TO be replaced by Any from protobuf
//	codec.RegisterInterface((*Mappable)(nil), nil)
//	codec.RegisterInterface((*QueryRequest)(nil), nil)
//	codec.RegisterInterface((*QueryResponse)(nil), nil)
//	codec.RegisterInterface((*TransactionRequest)(nil), nil)
//}
