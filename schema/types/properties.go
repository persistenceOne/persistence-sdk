/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/gogo/protobuf/proto"

type Properties interface {
	proto.Message
	Get(ID) Property

	GetList() []Property

	Add(...Property) Properties
	Remove(...Property) Properties
	Mutate(...Property) Properties
}
