/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

type NFT interface {
	//GetID() ID
	GetClassificationID() test_types.ID
}
