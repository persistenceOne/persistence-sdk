/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
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
			return orderID{ClassificationID: test_types.NewID(""), MakerOwnableID: test_types.NewID(""), TakerOwnableID: test_types.NewID(""), RateID: test_types.NewID(""), CreationID: test_types.NewID(""), MakerID: test_types.NewID(""), HashID: test_types.NewID("")}
		}

		height, Error := strconv.ParseInt(idList[4], 10, 64)
		if Error != nil {
			return orderID{ClassificationID: test_types.NewID(""), MakerOwnableID: test_types.NewID(""), TakerOwnableID: test_types.NewID(""), RateID: test_types.NewID(""), CreationID: test_types.NewID(""), MakerID: test_types.NewID(""), HashID: test_types.NewID("")}
		}

		return orderID{
			ClassificationID: test_types.NewID(idList[0]),
			MakerOwnableID:   test_types.NewID(idList[1]),
			TakerOwnableID:   test_types.NewID(idList[2]),
			RateID:           test_types.NewID(exchangeRate.String()),
			CreationID:       test_types.NewID(strconv.FormatInt(height, 10)),
			MakerID:          test_types.NewID(idList[5]),
			HashID:           test_types.NewID(idList[6]),
		}
	}

	return orderID{ClassificationID: test_types.NewID(""), MakerOwnableID: test_types.NewID(""), TakerOwnableID: test_types.NewID(""), RateID: test_types.NewID(""), CreationID: test_types.NewID(""), MakerID: test_types.NewID(""), HashID: test_types.NewID("")}
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

func ReadClassificationID(orderID test_types.ID) test_types.ID {
	return orderIDFromInterface(orderID).ClassificationID
}

func ReadRateID(orderID test_types.ID) test_types.ID {
	return orderIDFromInterface(orderID).RateID
}

func ReadCreationID(orderID test_types.ID) test_types.ID {
	return orderIDFromInterface(orderID).CreationID
}

func ReadMakerOwnableID(orderID test_types.ID) test_types.ID {
	return orderIDFromInterface(orderID).MakerOwnableID
}

func ReadTakerOwnableID(orderID test_types.ID) test_types.ID {
	return orderIDFromInterface(orderID).TakerOwnableID
}

func ReadMakerID(orderID test_types.ID) test_types.ID {
	return orderIDFromInterface(orderID).MakerID
}

func FromID(id test_types.ID) helpers.Key {
	return orderIDFromInterface(id)
}
