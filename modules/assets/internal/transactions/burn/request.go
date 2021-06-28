/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package burn

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

func newTransactionRequest(baseReq rest.BaseReq, fromID string, assetID string) *TransactionRequest {
	transReq := &TransactionRequest{
		FromID:  fromID,
		AssetID: assetID,
	}
	var bReq rest.BaseReq

	bReq = baseReq

	bReqAny, err := codectypes.NewAnyWithValue(bReq)

	if err != nil {
		panic(err)
	}
	transReq.BaseReq = bReqAny
	return transReq
}

//var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

//Implementing the helpers.TransactionRequest interface for Base Req



func (transactionRequest *TransactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest *TransactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.AssetID),
	), nil
}
func (transactionRequest *TransactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if Error := json.Unmarshal(rawMessage, &transactionRequest); Error != nil {
		return nil, Error
	}

	return transactionRequest, nil
}

func (transactionRequest *TransactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		return nil, Error
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.AssetID),
	), nil
}
func (*TransactionRequest) RegisterCodec(codec *codec.ProtoCodec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, TransactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return (*TransactionRequest){}
}

