/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/persistenceOne/persistenceSDK/schema/test_types"

type MetaProperties interface {

	test_types.ProtoInterface
	Get(ID) MetaProperty

	GetList() []MetaProperty

	Add(...MetaProperty) MetaProperties
	Remove(...MetaProperty) MetaProperties
	Mutate(...MetaProperty) MetaProperties

	RemoveData() Properties
}
