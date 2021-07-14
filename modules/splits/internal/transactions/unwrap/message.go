/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package unwrap

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

//type message struct {
//	From      sdkTypes.AccAddress `json:"from" valid:"required~required field from missing"`
//	FromID    types.ID            `json:"fromID" valid:"required~required field fromID missing"`
//	OwnableID types.ID            `json:"ownableID" valid:"required~required field ownableID missing"`
//	Value     sdkTypes.Int        `json:"value" valid:"required~required field value missing"`
//}

var _ sdkTypes.Msg = (*Message)(nil)

func (message Message) Route() string { return module.Name }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}

	return nil
}
func (message Message) GetSignBytes() []byte {
	a, _ := json.Marshal(messagePrototype)
	return sdkTypes.MustSortJSON(a)
	//return sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message))
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{sdkTypes.AccAddress(message.From)}
}
func (Message) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, Message{})
}
func messageFromInterface(msg sdkTypes.Msg) Message {
	switch value := msg.(type) {
	case *Message:
		return *value
	default:
		return Message{}
	}
}
func messagePrototype() helpers.Message {
	return &Message{}
}

func newMessage(from sdkTypes.AccAddress, fromID test_types.ID, ownableID test_types.ID, value sdkTypes.Dec) sdkTypes.Msg {
	return &Message{
		From:      string(from),
		FromID:    fromID,
		OwnableID: ownableID,
		Value:     value,
	}
}
