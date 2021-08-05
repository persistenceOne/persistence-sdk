/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	base2 "github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func readAssetID(assetIDString string) protoTypes.ID {
	idList := strings.Split(assetIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return assetID{
			ClassificationID: base2.NewID(idList[0]),
			HashID:           base2.NewID(idList[1]),
		}
	}

	return assetID{ClassificationID: base2.NewID(""), HashID: base2.NewID("")}
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

func ReadClassificationID(assetID protoTypes.ID) protoTypes.ID {
	return assetIDFromInterface(assetID).ClassificationID
}

func FromID(id protoTypes.ID) helpers.Key {
	return assetIDFromInterface(id)
}
