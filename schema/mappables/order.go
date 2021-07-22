/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Order interface {
	GetClassificationID() test_types.ID
	GetRateID() test_types.ID
	GetCreationID() test_types.ID
	GetMakerOwnableID() test_types.ID
	GetTakerOwnableID() test_types.ID
	GetMakerID() test_types.ID

	GetCreation() types.MetaProperty
	GetExchangeRate() types.MetaProperty

	GetTakerID() types.Property
	GetExpiry() types.Property
	GetMakerOwnableSplit() types.Property

	Document
}
