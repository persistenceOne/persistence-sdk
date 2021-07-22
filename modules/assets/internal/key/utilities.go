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

func readAssetID(assetIDString string) types.ID {
	idList := strings.Split(assetIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return assetID{
			ClassificationID: test_types.NewID(idList[0]),
			HashID:           test_types.NewID(idList[1]),
		}
	}

	return assetID{ClassificationID: test_types.NewID(""), HashID: test_types.NewID("")}
}
func assetIDFromInterface(i interface{}) assetID {
	switch value := i.(type) {
	case assetID:
		return value
	case types.ID:
		return assetIDFromInterface(readAssetID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(assetID test_types.ID) test_types.ID {
	return assetIDFromInterface(assetID).ClassificationID
}

func FromID(id test_types.ID) helpers.Key {
	return assetIDFromInterface(id)
}
