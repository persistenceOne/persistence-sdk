/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package revoke

import (
	"encoding/json"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
)

func Test_Revoke_Request(t *testing.T) {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.ImmutableMetaProperties, flags.ImmutableProperties, flags.MutableMetaProperties, flags.MutableProperties})
	cliContext := client.Context{}.WithLegacyAmino(Codec)


	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	require.Equal(t, nil, Error)

	testBaseReq := test_types.BaseReq{From: fromAddress, ChainId: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", "toID", "classificationID")

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", ToID: "toID", ClassificationID: "classificationID"}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, Error := transactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, Error)
	require.Equal(t, transactionRequest{BaseReq: test_types.BaseReq{From: cliContext.GetFromAddress().String(), ChainId: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", ToID: "", ClassificationID: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, Error := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, Error)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, Error := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, Error)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, Error := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, testBase.NewID("fromID"), testBase.NewID("toID"), testBase.NewID("classificationID")), msg)
	require.Nil(t, Error)

	msg2, Error := newTransactionRequest(test_types.BaseReq{From: "randomString", ChainId: "test", Fees: sdkTypes.NewCoins()}, "fromID", "toID", "classificationID").MakeMsg()
	require.NotNil(t, Error)
	require.Nil(t, msg2)

	msg2, Error = newTransactionRequest(testBaseReq, "fromID", "toID", "classificationID").MakeMsg()
	require.Nil(t, Error)
	require.NotNil(t, msg2)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterCodec(codec.NewLegacyAmino())
	})
}
