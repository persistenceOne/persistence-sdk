/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package provision

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)


var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.To),
		cliCommand.ReadString(flags.IdentityID),
	), nil
}
func (transactionRequest transactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if Error := json.Unmarshal(rawMessage, &transactionRequest); Error != nil {
		return nil, Error
	}

	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetBaseReq() test_types.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		return nil, Error
	}

	to, Error := sdkTypes.AccAddressFromBech32(transactionRequest.To)
	if Error != nil {
		return nil, Error
	}

	return newMessage(
		from,
		to,
		testBase.NewID(transactionRequest.IdentityID),
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq test_types.BaseReq, to string, identityID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:    baseReq,
		To:         to,
		IdentityID: identityID,
	}
}
