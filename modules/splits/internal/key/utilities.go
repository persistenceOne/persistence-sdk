/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
)

func readSplitID(splitIDString string) test_types.ID {
	idList := strings.Split(splitIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return splitID{
			OwnerID:   testBase.NewID(idList[0]),
			OwnableID: testBase.NewID(idList[1]),
		}
	}

	return splitID{OwnerID: testBase.NewID(""), OwnableID: testBase.NewID("")}
}

func splitIDFromInterface(i interface{}) splitID {
	switch value := i.(type) {
	case splitID:
		return value
	case test_types.ID:
		return splitIDFromInterface(readSplitID(value.String()))
	default:
		panic(i)
	}
}

func ReadOwnableID(id test_types.ID) test_types.ID {
	return splitIDFromInterface(id).OwnableID
}

func ReadOwnerID(id test_types.ID) test_types.ID {
	return splitIDFromInterface(id).OwnerID
}

func FromID(id test_types.ID) helpers.Key {
	return splitIDFromInterface(id)
}

func ToID(key helpers.Key) test_types.ID {
	return splitIDFromInterface(key)
}
