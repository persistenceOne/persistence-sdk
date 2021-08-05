/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
)

type NFT interface {
	ID() ID
	GetClassificationID() protoTypes.ID
}
