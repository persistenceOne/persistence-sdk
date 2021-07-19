/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func readIdentityID(identityIDString string) types.ID {
	idList := strings.Split(identityIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return identityID{
			ClassificationID: test_types.NewID(idList[0]),
			HashID:           test_types.NewID(idList[1]),
		}
	}

	return identityID{ClassificationID: test_types.NewID(""), HashID: test_types.NewID("")}
}

func identityIDFromInterface(i interface{}) identityID {
	switch value := i.(type) {
	case identityID:
		return value
	case types.ID:
		return identityIDFromInterface(readIdentityID(value.String()))
	default:
		panic(i)
	}
}

func FromID(id test_types.ID) helpers.Key {
	return identityIDFromInterface(id)
}
