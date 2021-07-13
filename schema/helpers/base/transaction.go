/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/persistenceOne/persistenceSDK/utilities/random"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	//"github.com/cosmos/cosmos-sdk/x/auth"
	//authClient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	//"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/utilities/rest/queuing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type transaction struct {
	name             string
	cliCommand       helpers.CLICommand
	keeper           helpers.TransactionKeeper
	requestPrototype func() helpers.TransactionRequest
	messagePrototype func() helpers.Message
	keeperPrototype  func() helpers.TransactionKeeper
}

var _ helpers.Transaction = (*transaction)(nil)

func (transaction transaction) GetName() string { return transaction.name }
func (transaction transaction) Command(codec *codec.ProtoCodec) *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		//var cliContext client.Context
		//cliContext,err := client.GetClientTxContext(command)
		//if err!=nil{
		//	return err
		//}
		//txCfg:=cliContext.TxConfig
		//bufioReader := bufio.NewReader(command.InOrStdin())
		//txBuilder,err := txCfg.WrapTxBuilder()
		//transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(authClient.GetTxEncoder(codec))

		cliContext,Error := client.GetClientTxContext(command)
		if Error != nil {
			return Error
		}
		transactionRequest, Error := transaction.requestPrototype().FromCLI(transaction.cliCommand, cliContext)
		if Error != nil {
			return Error
		}

		msg, Error := transactionRequest.MakeMsg()
		if Error != nil {
			return Error
		}

		if Error := msg.ValidateBasic(); Error != nil {
			return Error
		}

		return tx.GenerateOrBroadcastTxCLI(cliContext,command.Flags(),msg)
	}

	return transaction.cliCommand.CreateCommand(runE)
}

func (transaction transaction) HandleMessage(context sdkTypes.Context, message sdkTypes.Msg) (*sdkTypes.Result, error) {
	if transactionResponse := transaction.keeper.Transact(context, message); !transactionResponse.IsSuccessful() {
		return nil, transactionResponse.GetError()
	}

	context.EventManager().EmitEvent(
		sdkTypes.NewEvent(
			sdkTypes.EventTypeMessage,
			sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, transaction.name),
		),
	)

	return &sdkTypes.Result{Events: context.EventManager().ABCIEvents()}, nil
}

func (transaction transaction) RESTRequestHandler(cliContext client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := transaction.requestPrototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.LegacyAmino, &transactionRequest) {
			return
		} else if reflect.TypeOf(transaction.requestPrototype()) != reflect.TypeOf(transactionRequest) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		Error := transactionRequest.Validate()
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		bReq := transactionRequest.GetBaseReq()

		msg, Error := transactionRequest.MakeMsg()
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		//convert persistence base request to cosmos base request
		baseReq := rest.NewBaseReq(bReq.From,bReq.Memo,bReq.ChainId,bReq.Gas,bReq.GasAdjustment,bReq.AccountNumber,bReq.Sequence,bReq.Fees,bReq.GasPrices,bReq.Simulate)

		if !baseReq.ValidateBasic(responseWriter) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		if Error := msg.ValidateBasic(); Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		if viper.GetBool(flags.FlagGenerateOnly) {
			tx.WriteGeneratedTxResponse( cliContext, responseWriter,baseReq, msg)
			return
		}

		kr,err := keyring.New(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend),viper.GetString(flags.FlagHome),strings.NewReader(keys.DefaultKeyPass))
		if err!=nil{
			rest.WriteErrorResponse(responseWriter,http.StatusBadRequest,err.Error())
		}

		fromAddress, fromName, _, Error := client.GetFromFields(kr, baseReq.From, viper.GetBool(flags.FlagGenerateOnly))
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		cliContext = cliContext.WithFromAddress(fromAddress)
		cliContext = cliContext.WithFromName(fromName)
		cliContext = cliContext.WithBroadcastMode(viper.GetString(flags.FlagBroadcastMode))


		gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(responseWriter, baseReq.GasAdjustment, flags.DefaultGasAdjustment)
		if !ok {
			return
		}

		gasSetting,err := flags.ParseGasSetting(baseReq.Gas)
		if rest.CheckBadRequestError(responseWriter, err) {
			return
		}

		txf := tx.Factory{}.
			WithAccountNumber(baseReq.AccountNumber).
			WithSequence(baseReq.Sequence).
			WithGas(gasSetting.Gas).
			WithGasAdjustment(gasAdj).
			WithMemo(baseReq.Memo).
			WithChainID(baseReq.ChainID).
			WithSimulateAndExecute(baseReq.Simulate).
			WithTxConfig(cliContext.TxConfig).
			WithTimeoutHeight(baseReq.TimeoutHeight).
			WithFees(baseReq.Fees.String()).
			WithGasPrices(baseReq.GasPrices.String())

		txf,err = tx.PrepareFactory(cliContext,txf)
		if rest.CheckBadRequestError(responseWriter, err) {
			return
		}

		//simAndExec, gas, Error := flags.ParseGas(baseReq.Gas)
		//if Error != nil {
		//	rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
		//	return
		//}

		//txBuilder := types.NewTxBuilder(
		//	authClient.GetTxEncoder(cliContext.Codec), baseReq.AccountNumber, baseReq.Sequence, gas, gasAdj,
		//	baseReq.Simulate, baseReq.ChainID, baseReq.Memo, baseReq.Fees, baseReq.GasPrices,
		//)
		//msgList := []sdkTypes.Msg{msg}

		if baseReq.Simulate || gasSetting.Simulate {
			if gasAdj < 0 {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, errors.ErrOutOfGas.Error())
				return
			}

			_, adjusted, err := tx.CalculateGas(cliContext.QueryWithData, txf, msg)
			if rest.CheckInternalServerError(responseWriter, err) {
				return
			}

			txf = txf.WithGas(adjusted)

			//txBuilder, Error = authClient.EnrichWithGas(txBuilder, cliContext, msgList)
			//if Error != nil {
			//	rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			//	return
			//}

			if baseReq.Simulate {
				rest.WriteSimulationResponse(responseWriter, cliContext.LegacyAmino, txf.Gas())
				return
			}
		}




		//fromAddress, fromName, Error := cliContext.
		//if Error != nil {
		//	rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
		//	return
		//}
		//
		//cliContext = cliContext.WithFromAddress(fromAddress)
		//cliContext = cliContext.WithFromName(fromName)
		//cliContext = cliContext.WithBroadcastMode(viper.GetString(flags.FlagBroadcastMode))

		if queuing.KafkaState.IsEnabled {
			ticketID := queuing.TicketID(random.GenerateID(transaction.name))
			jsonResponse := queuing.SendToKafka(queuing.NewKafkaMsgFromRest(msg, ticketID, baseReq, cliContext), queuing.KafkaState, cliContext.LegacyAmino)

			responseWriter.WriteHeader(http.StatusAccepted)
			_, _ = responseWriter.Write(jsonResponse)
		} else {
			//accountNumber, sequence, Error := types.NewAccountRetriever(cliContext).GetAccountNumberSequence(fromAddress)
			//if Error != nil {
			//	rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			//	return
			//}
			//
			//txBuilder = txBuilder.WithAccountNumber(accountNumber)
			//txBuilder = txBuilder.WithSequence(sequence)
			//
			//stdMsg, Error := txBuilder.BuildAndSign(fromName, keys.DefaultKeyPass, msgList)
			//if Error != nil {
			//	rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			//	return
			//}
			txBuilder, err := tx.BuildUnsignedTx(txf,msg)
			if rest.CheckBadRequestError(responseWriter, err) {
				return
			}

			err = tx.Sign(txf, cliContext.FromName, txBuilder, true)
			if rest.CheckBadRequestError(responseWriter, err) {
				return
			}

			txBytes, err := cliContext.TxConfig.TxEncoder()(txBuilder.GetTx())
			if rest.CheckInternalServerError(responseWriter, err) {
				return
			}


			// broadcast to a node
			response, err := cliContext.BroadcastTx(txBytes)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}
			responseBytes, err := cliContext.JSONMarshaler.MarshalJSON(response)
			if err!=nil {
				rest.WriteErrorResponse(responseWriter,http.StatusBadRequest,err.Error())
				return
			}

			wrappedResponse := rest.NewResponseWithHeight(cliContext.Height, responseBytes)

			output, err := cliContext.LegacyAmino.MarshalJSON(wrappedResponse)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			responseWriter.Header().Set("Content-Type", "application/json")
			if _, Error := responseWriter.Write(output); Error != nil {
				log.Printf("could not write response: %v", Error)
			}
		}
	}
}

func (transaction transaction) RegisterCodec(codec *codec.LegacyAmino) {
	transaction.messagePrototype().RegisterCodec(codec)
	transaction.requestPrototype().RegisterCodec(codec)
}
func (transaction transaction) DecodeTransactionRequest(rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	transactionRequest, Error := transaction.requestPrototype().FromJSON(rawMessage)
	if Error != nil {
		return nil, Error
	}

	return transactionRequest.MakeMsg()
}

func (transaction transaction) InitializeKeeper(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaryKeepers ...interface{}) helpers.Transaction {
	transaction.keeper = transaction.keeperPrototype().Initialize(mapper, parameters, auxiliaryKeepers).(helpers.TransactionKeeper)
	return transaction
}

func NewTransaction(name string, short string, long string, requestPrototype func() helpers.TransactionRequest, messagePrototype func() helpers.Message, keeperPrototype func() helpers.TransactionKeeper, flagList ...helpers.CLIFlag) helpers.Transaction {
	return transaction{
		name:             name,
		cliCommand:       NewCLICommand(name, short, long, flagList),
		requestPrototype: requestPrototype,
		messagePrototype: messagePrototype,
		keeperPrototype:  keeperPrototype,
	}
}
