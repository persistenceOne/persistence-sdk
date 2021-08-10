/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func readClassificationID(classificationIDString string) protoTypes.ID {
	idList := strings.Split(classificationIDString, constants.IDSeparator)
	if len(idList) == 2 {
		return classificationID{
			ChainID: base.NewID(idList[0]),
			HashID:  base.NewID(idList[1]),

		}
	}

	return classificationID{ChainID: base.NewID(""), HashID: base.NewID("")}
}
func classificationIDFromInterface(i interface{}) classificationID {
	switch value := i.(type) {
	case classificationID:
		return value
	case types.ID:
		return classificationIDFromInterface(readClassificationID(value.String()))
	default:
		panic(i)
	}
}

func FromID(id protoTypes.ID) helpers.Key {
	return classificationIDFromInterface(id)
}
