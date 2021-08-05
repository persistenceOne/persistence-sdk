/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
)

type TransactionRequest interface {
	//proto.Message
	GetBaseReq() protoTypes.BaseReq

	FromCLI(CLICommand, client.Context) (TransactionRequest, error)
	FromJSON(json.RawMessage) (TransactionRequest, error)
	MakeMsg() (sdkTypes.Msg, error)
	//TODO: how to use protocodec
	RegisterCodec(*codec.LegacyAmino)
	Request
}
