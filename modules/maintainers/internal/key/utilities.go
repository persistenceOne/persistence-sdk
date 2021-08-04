/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"

	"github.com/persistenceOne/persistenceSDK/constants"
)

func readMaintainerID(maintainerIDString string) test_types.ID {
	idList := strings.Split(maintainerIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return maintainerID{
			ClassificationID: testBase.NewID(idList[0]),
			IdentityID:       testBase.NewID(idList[1]),
		}
	}

	return maintainerID{IdentityID: testBase.NewID(""), ClassificationID: testBase.NewID("")}
}
func maintainerIDFromInterface(i interface{}) maintainerID {
	switch value := i.(type) {
	case maintainerID:
		return value
	case test_types.ID:
		return maintainerIDFromInterface(readMaintainerID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(assetID test_types.ID) test_types.ID {
	return maintainerIDFromInterface(assetID).ClassificationID
}

func ReadIdentityID(assetID test_types.ID) test_types.ID {
	return maintainerIDFromInterface(assetID).IdentityID
}

func FromID(id test_types.ID) helpers.Key {
	return maintainerIDFromInterface(id)
}
