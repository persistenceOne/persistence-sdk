/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Maintainer interface {
	GetClassificationID() types.ID
	GetIdentityID() types.ID
	GetMaintainedProperties() types.Properties

	CanAddMaintainer() bool
	CanRemoveMaintainer() bool
	CanMutateMaintainer() bool

	MaintainsProperty(protoTypes.ID) bool
	helpers.Mappable
}
