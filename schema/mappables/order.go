/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Order interface {
	GetClassificationID() protoTypes.ID
	GetRateID() protoTypes.ID
	GetCreationID() protoTypes.ID
	GetMakerOwnableID() protoTypes.ID
	GetTakerOwnableID() protoTypes.ID
	GetMakerID() protoTypes.ID

	GetCreation() types.MetaProperty
	GetExchangeRate() types.MetaProperty

	GetTakerID() types.Property
	GetExpiry() types.Property
	GetMakerOwnableSplit() types.Property

	Document
}
