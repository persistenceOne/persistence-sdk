/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	base2 "github.com/persistenceOne/persistenceSDK/schema/proto/types/base"

	"strconv"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func readOrderID(orderIDString string) types.ID {
	idList := strings.Split(orderIDString, constants.SecondOrderCompositeIDSeparator)

	if len(idList) == 7 {
		exchangeRate, Error := sdkTypes.NewDecFromStr(idList[3])
		if Error != nil {
			return orderID{ClassificationID: base2.NewID(""), MakerOwnableID: base2.NewID(""), TakerOwnableID: base2.NewID(""), RateID: base2.NewID(""), CreationID: base2.NewID(""), MakerID: base2.NewID(""), HashID: base2.NewID("")}
		}

		height, Error := strconv.ParseInt(idList[4], 10, 64)
		if Error != nil {
			return orderID{ClassificationID: base2.NewID(""), MakerOwnableID: base2.NewID(""), TakerOwnableID: base2.NewID(""), RateID: base2.NewID(""), CreationID: base2.NewID(""), MakerID: base2.NewID(""), HashID: base2.NewID("")}
		}

		return orderID{
			ClassificationID: base2.NewID(idList[0]),
			MakerOwnableID:   base2.NewID(idList[1]),
			TakerOwnableID:   base2.NewID(idList[2]),
			RateID:           base2.NewID(exchangeRate.String()),
			CreationID:       base2.NewID(strconv.FormatInt(height, 10)),
			MakerID:          base2.NewID(idList[5]),
			HashID:           base2.NewID(idList[6]),
		}
	}

	return orderID{ClassificationID: base2.NewID(""), MakerOwnableID: base2.NewID(""), TakerOwnableID: base2.NewID(""), RateID: base2.NewID(""), CreationID: base2.NewID(""), MakerID: base2.NewID(""), HashID: base2.NewID("")}
}
func orderIDFromInterface(i interface{}) orderID {
	switch value := i.(type) {
	case orderID:
		return value
	case types.ID:
		return orderIDFromInterface(readOrderID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(orderID protoTypes.ID) protoTypes.ID {
	return orderIDFromInterface(orderID).ClassificationID
}

func ReadRateID(orderID protoTypes.ID) protoTypes.ID {
	return orderIDFromInterface(orderID).RateID
}

func ReadCreationID(orderID protoTypes.ID) protoTypes.ID {
	return orderIDFromInterface(orderID).CreationID
}

func ReadMakerOwnableID(orderID protoTypes.ID) protoTypes.ID {
	return orderIDFromInterface(orderID).MakerOwnableID
}

func ReadTakerOwnableID(orderID protoTypes.ID) protoTypes.ID {
	return orderIDFromInterface(orderID).TakerOwnableID
}

func ReadMakerID(orderID protoTypes.ID) protoTypes.ID {
	return orderIDFromInterface(orderID).MakerID
}

func FromID(id protoTypes.ID) helpers.Key {
	return orderIDFromInterface(id)
}
