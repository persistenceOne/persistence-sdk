/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package wrap

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

//type transactionResponse struct {
//	Success bool
//	Error   error
//}

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (transactionResponse TransactionResponse) IsSuccessful() bool {
	return transactionResponse.Success
}
func (transactionResponse TransactionResponse) GetError() error {
	return transactionResponse.Error
}
func newTransactionResponse(error error) helpers.TransactionResponse {
	success := true
	if error != nil {
		success = false
	}

	return TransactionResponse{
		Success: success,
		Error:   error,
	}
}
