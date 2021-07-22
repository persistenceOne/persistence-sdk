/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
)

type Document interface {
	GetID() test_types.ID

	traits.HasImmutables
	traits.HasMutables

	helpers.Mappable
}
