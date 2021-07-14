/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

type Ownable interface {
	GetOwnerID() test_types.ID
	GetOwnableID() test_types.ID
}
