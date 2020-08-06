/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package recover

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type request struct {
	Name     string `json:"name"`
	Mnemonic string `json:"mnemonic"`
}

var _ helpers.Request = request{}
