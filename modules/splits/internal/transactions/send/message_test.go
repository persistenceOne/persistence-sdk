/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	"fmt"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/utilities/transaction"
	"github.com/stretchr/testify/require"
)

func Test_Send_Message(t *testing.T) {

	testToID := test_types.NewID("toID")
	testFromID := test_types.NewID("fromID")
	testOwnableID := test_types.NewID("ownableID")
	testSplit := sdkTypes.NewDec(2)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)
	testMessage := newMessage(fromAccAddress, testFromID, testToID, testOwnableID, testSplit)
	//fmt.Println(testMessage)
	require.Equal(t, Message{From: fromAccAddress, FromID: testFromID, ToID: testToID, OwnableID: testOwnableID, Value: testSplit}, testMessage)
	//fmt.Println(module.Name)
	//fmt.Println(testMessage.Route())
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	//fmt.Println(Transaction.GetName())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, Message{}.ValidateBasic())
	fmt.Println(string(sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage))))
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())

	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, Message{}, messageFromInterface(nil))
	require.Equal(t, Message{}, messagePrototype())

}
