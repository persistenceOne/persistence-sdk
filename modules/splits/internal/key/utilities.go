/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"strings"
)

func readSplitID(splitIDString string) types.ID {
	idList := strings.Split(splitIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return splitID{
			OwnerID:   test_types.NewID(idList[0]),
			OwnableID: test_types.NewID(idList[1]),
		}
	}

	return splitID{OwnerID: test_types.NewID(""), OwnableID: test_types.NewID("")}
}

func splitIDFromInterface(i interface{}) splitID {
	switch value := i.(type) {
	case splitID:
		return value
	case types.ID:
		return splitIDFromInterface(readSplitID(value.String()))
	default:
		panic(i)
	}
}

func ReadOwnableID(id types.ID) test_types.ID {
	return splitIDFromInterface(id).OwnableID
}

func ReadOwnerID(id types.ID) test_types.ID {
	return splitIDFromInterface(id).OwnerID
}

func FromID(id types.ID) helpers.Key {
	return splitIDFromInterface(id)
}

func ToID(key helpers.Key) types.ID {
	return splitIDFromInterface(key)
}
