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
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func readClassificationID(classificationIDString string) test_types.ID {
	idList := strings.Split(classificationIDString, constants.IDSeparator)
	if len(idList) == 2 {
		return classificationID{
			ChainID: testBase.NewID(idList[0]),
			HashID:  testBase.NewID(idList[1]),
		}
	}

	return classificationID{ChainID: testBase.NewID(""), HashID: testBase.NewID("")}
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

func FromID(id test_types.ID) helpers.Key {
	return classificationIDFromInterface(id)
}
