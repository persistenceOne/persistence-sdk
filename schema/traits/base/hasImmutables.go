/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	base2"github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type HasImmutables struct {
	Properties types.Properties `json:"properties"`
}

var _ traits.HasImmutables = (*HasImmutables)(nil)

func (immutables HasImmutables) GetImmutableProperties() types.Properties {
	return immutables.Properties
}
func (immutables HasImmutables) GenerateHashID() protoTypes.ID {
	metaList := make([]string, len(immutables.Properties.GetList()))

	for i, immutableProperty := range immutables.Properties.GetList() {
		metaList[i] = immutableProperty.GetFact().GetHashID().String()
	}

	return base2.NewID(metaUtilities.Hash(metaList...))
}
