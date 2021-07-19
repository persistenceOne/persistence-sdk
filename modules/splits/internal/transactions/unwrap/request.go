/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package unwrap

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

//type transactionRequest struct {
//	BaseReq   rest.BaseReq `json:"baseReq"`
//	FromID    string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
//	OwnableID string       `json:"ownableID" valid:"required~required field ownableID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field ownableID"`
//	Value     string       `json:"value" valid:"required~required field value missing, matches(^[0-9]+$)~invalid field value"`
//}

var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

func (transactionRequest TransactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest TransactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.OwnableID),
		cliCommand.ReadString(flags.Value),
	), nil
}
func (transactionRequest TransactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if Error := json.Unmarshal(rawMessage, &transactionRequest); Error != nil {
		return nil, Error
	}

	return transactionRequest, nil
}
func (transactionRequest TransactionRequest) GetBaseReq() test_types.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest TransactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		return nil, Error
	}

	value, ok := sdkTypes.NewIntFromString(transactionRequest.Value)
	if !ok {
		return nil, xprtErrors.InvalidRequest
	}

	return newMessage(
		from,
		test_types.NewID(transactionRequest.FromID),
		test_types.NewID(transactionRequest.OwnableID),
		value,
	), nil
}
func (TransactionRequest) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, TransactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return TransactionRequest{}
}
func newTransactionRequest(baseReq test_types.BaseReq, fromID string, ownableID string, value string) TransactionRequest {
	return TransactionRequest{
		BaseReq:   baseReq,
		FromID:    fromID,
		OwnableID: ownableID,
		Value:     value,
	}
}
