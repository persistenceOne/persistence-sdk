/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package test_helpers

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

//type transactions struct {
//	transactionList []helpers.Transaction
//}

//var _ helpers.Transactions = &Transactions{}
//var _ cdctypes.UnpackInterfacesMessage = &Transactions{}
//
//func (transactions *Transactions) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
//	var a helpers.Transactions
//	return unpacker.UnpackAny(transactions.TransactionList, &a)
//}
//
//func (transactions *Transactions) Get(name string) helpers.Transaction {
//	for _, transaction := range transactions.GetList() {
//		if transaction.GetName() == name {
//			return transaction
//		}
//	}
//
//	return nil
//}
//
//func (transactions *Transactions) GetList() []helpers.Transaction {
//	arr := make([]helpers.Transaction, len(transactions.TransactionList))
//	for  i, transany := range transactions.TransactionList {
//		transc, ok := transany.GetCachedValue().(helpers.Transaction)
//		if !ok { return nil }
//		arr[i] = transc
//	}
//	return arr
//}
//
//func NewTransactions(transactionList ...helpers.Transaction) helpers.Transactions {
//	g := &Transactions{}
//	for _, element := range transactionList {
//		msg, ok := element.(proto.Message)
//		if !ok { return &Transactions{}}
//		any, err := cdctypes.NewAnyWithValue(msg)
//		if err!=nil { return &Transactions{}}
//		g.TransactionList = append(g.TransactionList, any)
//	}
//	return g
//}

//func (transactions transactions) Get(name string) helpers.Transaction {
//	for _, transaction := range transactions.transactionList {
//		if transaction.GetName() == name {
//			return transaction
//		}
//	}
//
//	return nil
//}
//
//func (transactions transactions) GetList() []helpers.Transaction {
//	return transactions.transactionList
//}
//
//func NewTransactions(transactionList ...helpers.Transaction) helpers.Transactions {
//	return transactions{
//		transactionList: transactionList,
//	}
//}

var _ helpers.Transactions = (*Transactions)(nil)
var _ types.UnpackInterfacesMessage = (*Transactions)(nil)
var _ helpers.Transactions = (*Transaction)(nil)

func (m *Transaction) Get(s string) helpers.Transaction {
	panic("implement me")
}

func (m *Transaction) GetList() []helpers.Transaction {
	panic("implement me")
}
func (transactions *Transactions) UnpackInterfaces(unpacker types.AnyUnpacker) error {

	err_arr := make([]error, len(transactions.TransactionList))
	for i, j := range transactions.TransactionList {
		//k, _ := types.NewAnyWithValue(j)
		var transac helpers.Transactions
		err := unpacker.UnpackAny(j, &transac)
		err_arr[i] = err
	}
	for _, j := range err_arr {
		if j != nil {
			return j
		}
	}
	return nil
}

func (transactions *Transactions) Get(name string) helpers.Transaction {
	for _, transaction := range transactions.GetList() {
		if transaction.GetName() == name {
			return transaction
		}
	}

	return nil
}

func (transactions *Transactions) GetList() []helpers.Transaction {
	arr := make([]helpers.Transaction, len(transactions.TransactionList))
	for i, transany := range transactions.TransactionList {
		transc, ok := transany.GetCachedValue().(helpers.Transaction)
		if !ok {
			return nil
		}
		arr[i] = transc
	}
	return arr
}

func NewTransactions(transactionList ...helpers.Transaction) helpers.Transactions {
	g := Transactions{}
	for _, element := range transactionList {
		msg, ok := element.(proto.Message)
		if !ok {
			return &Transactions{}
		}
		any, err := types.NewAnyWithValue(msg)
		if err != nil {
			return &Transactions{}
		}
		g.TransactionList = append(g.TransactionList, any)
	}
	return &g
}

//var _ helpers.Transaction = (*Transaction)(nil)
//var _ types.UnpackInterfacesMessage = (*Transaction)(nil)
//
//func (m *Transaction) UnpackInterfaces(unpacker types.AnyUnpacker) error {
//	var a helpers.CLICommand
//	var b helpers.TransactionKeeper
//	err1 := unpacker.UnpackAny(m.CliCommand, &a)
//	err2 := unpacker.UnpackAny(m.Keeper, &b)
//	if err1 != nil {return err1}
//	if err2 != nil {return err2}
//	return nil
//}
//
//func (m *Transaction) GetName() string {
//	return m.Name
//}
//
//func (m *Transaction) Command() *cobra.Command {
//	panic("implement me")
//}
//
//func (m *Transaction) HandleMessage(context sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
//	panic("implement me")
//}
//
//func (m *Transaction) RESTRequestHandler(context client.Context) http.HandlerFunc {
//	panic("implement me")
//}
//
//func (m *Transaction) RegisterCodec(amino *codec.LegacyAmino) {
//	panic("implement me")
//}
//
//func (m *Transaction) DecodeTransactionRequest(message json.RawMessage) (sdkTypes.Msg, error) {
//	panic("implement me")
//}
//
//func (m *Transaction) InitializeKeeper(mapper helpers.Mapper, parameters helpers.Parameters, i ...interface{}) helpers.Transaction {
//	panic("implement me")
//}
