/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	//"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	test_types "github.com/persistenceOne/persistenceSDK/schema/test_types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)



var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

func (transactionRequest TransactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest TransactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.ToID),
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

func (transactionRequest TransactionRequest) GetBaseReq() test_types.BaseReq   {
	return transactionRequest.BaseReq
}

func (transactionRequest TransactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		return nil, Error
	}

	value, Error := sdkTypes.NewDecFromStr(transactionRequest.Value)
	if Error != nil {
		return nil, Error
	}

	return newMessage(
		from,
		test_types.NewID(transactionRequest.FromID),
		test_types.NewID(transactionRequest.ToID),
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
func newTransactionRequest(baseReq test_types.BaseReq, fromID string, toID string, ownableID string, value string) TransactionRequest {
	return TransactionRequest{
		BaseReq:   baseReq,
		FromID:    fromID,
		ToID:      toID,
		OwnableID: ownableID,
		Value:     value,
	}
}
