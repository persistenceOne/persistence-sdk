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
)

func readIdentityID(identityIDString string) protoTypes.ID {
	idList := strings.Split(identityIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return identityID{
			ClassificationID: base2.NewID(idList[0]),
			HashID:           base2.NewID(idList[1]),
		}
	}

	return identityID{ClassificationID: base2.NewID(""), HashID: base2.NewID("")}
}

func identityIDFromInterface(i interface{}) identityID {
	switch value := i.(type) {
	case identityID:
		return value
	case protoTypes.ID:
		return identityIDFromInterface(readIdentityID(value.String()))
	default:
		panic(i)
	}
}

func FromID(id protoTypes.ID) helpers.Key {
	return identityIDFromInterface(id)
}
