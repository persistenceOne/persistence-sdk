/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

const moduleName = "queuing"

func RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, kafkaCliCtx{})
	codecUtilities.RegisterXPRTConcrete(codec, moduleName, kafkaMsg{})
}

var ModuleCdc *codec.LegacyAmino

func init() {
	ModuleCdc = codec.NewLegacyAmino()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
