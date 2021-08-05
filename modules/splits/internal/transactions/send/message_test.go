/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	"fmt"
	base2 "github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/utilities/transaction"
	"github.com/stretchr/testify/require"
)

func Test_Send_Message(t *testing.T) {

	testToID := base2.NewID("toID")
	testFromID := base2.NewID("fromID")
	testOwnableID := base2.NewID("ownableID")
	testSplit := sdkTypes.NewDec(2)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)
	testMessage := newMessage(fromAccAddress, testFromID, testToID, testOwnableID, testSplit)
	require.Equal(t, message{from: fromAccAddress, fromID: testFromID, toID: testToID, ownableID: testOwnableID, value: testSplit}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, message{}.ValidateBasic())
	fmt.Println(string(sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage))))
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())

	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, message{}, messageFromInterface(nil))
	require.Equal(t, message{}, messagePrototype())

}
