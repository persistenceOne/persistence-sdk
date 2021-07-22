/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

var _ helpers.TransactionResponse = (*transactionResponse)(nil)

func (transactionResponse transactionResponse) IsSuccessful() bool {
	return transactionResponse.success
}
func (transactionResponse transactionResponse) GetError() error {
	return transactionResponse.error
}
func newTransactionResponse(error error) helpers.TransactionResponse {
	success := true
	if error != nil {
		success = false
	}

	return transactionResponse{
		success: success,
		error:   error,
	}
}
