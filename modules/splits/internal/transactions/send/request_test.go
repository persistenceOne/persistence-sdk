/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Send_Request(t *testing.T) {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.ToID, flags.OwnableID, flags.Value})
	cliContext := client.Context{}.WithLegacyAmino(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	testBaseReq := test_types.BaseReq{From: fromAddress, ChainId: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", "toID", "ownableID", "2")

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", ToID: "toID", OwnableID: "ownableID", Value: "2"}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, Error := transactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, Error)
	require.Equal(t, transactionRequest{BaseReq: test_types.BaseReq{From: cliContext.GetFromAddress().String(), ChainId: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", ToID: "", OwnableID: "", Value: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, Error := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, Error)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, Error := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, Error)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, Error := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, test_types.NewID("fromID"), test_types.NewID("toID"), test_types.NewID("ownableID"), sdkTypes.NewDec(2)), msg)
	require.Nil(t, Error)

	msg2, Error := newTransactionRequest(test_types.BaseReq{From: "randomFromAddress", ChainId: "test"}, "fromID", "toID", "ownableID", "2").MakeMsg()
	require.NotNil(t, Error)
	require.Nil(t, msg2)

	msg2, Error = newTransactionRequest(testBaseReq, "fromID", "toID", "ownableID", "randomString").MakeMsg()
	require.NotNil(t, Error)
	require.Nil(t, msg2)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterCodec(codec.NewLegacyAmino())
	})
}
