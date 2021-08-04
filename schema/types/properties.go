/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/persistenceOne/persistenceSDK/schema/test_types"

type Properties interface {
	test_types.ProtoInterface
	Get(test_types.ID) Property

	GetList() []Property

	Add(...Property) Properties
	Remove(...Property) Properties
	Mutate(...Property) Properties
}
