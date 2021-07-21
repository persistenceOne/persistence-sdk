/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package block

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

//type block struct {
//	mapper     helpers.Mapper
//	parameters helpers.Parameters
//}

var _ helpers.Block = (*Block)(nil)

func (m *Block) Begin(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {

}

func (m *Block) End(_ sdkTypes.Context, _ abciTypes.RequestEndBlock) {

}

func (block *Block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, i ...interface{}) helpers.Block {
	msg, ok := mapper.(proto.Message)
	if !ok {
		return nil
	}
	msg1, ok1 := parameters.(proto.Message)
	if !ok1 {
		return nil
	}
	mapp, err := types.NewAnyWithValue(msg)
	param, err1 := types.NewAnyWithValue(msg1)
	if err != nil {
		return nil
	}
	if err1 != nil {
		return nil
	}
	block.Mapper, block.Parameters = mapp, param
	return block
}

//func (block block) Begin(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {
//
//}
//
//func (block block) End(_ sdkTypes.Context, _ abciTypes.RequestEndBlock) {
//
//}
//
//func (block block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, _ ...interface{}) helpers.Block {
//	block.mapper, block.parameters = mapper, parameters
//	return block
//}
