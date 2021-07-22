/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type orderID struct {
	ClassificationID test_types.ID `json:"classificationID"`
	MakerOwnableID   test_types.ID `json:"makerOwnableID"`
	TakerOwnableID   test_types.ID `json:"takerOwnableID"`
	RateID           test_types.ID `json:"rateID"`
	CreationID       test_types.ID `json:"creationID"`
	MakerID          test_types.ID `json:"makerID"`
	HashID           test_types.ID `json:"hashID"`
}

var _ types.ID = (*orderID)(nil)
var _ helpers.Key = (*orderID)(nil)

func (orderID orderID) Bytes() []byte {
	var Bytes []byte

	rateIDBytes, Error := orderID.getRateIDBytes()
	if Error != nil {
		return Bytes
	}

	creationIDBytes, Error := orderID.getCreationHeightBytes()
	if Error != nil {
		return Bytes
	}

	Bytes = append(Bytes, orderID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, orderID.MakerOwnableID.Bytes()...)
	Bytes = append(Bytes, orderID.TakerOwnableID.Bytes()...)
	Bytes = append(Bytes, rateIDBytes...)
	Bytes = append(Bytes, creationIDBytes...)
	Bytes = append(Bytes, orderID.MakerID.Bytes()...)
	Bytes = append(Bytes, orderID.HashID.Bytes()...)

	return Bytes
}
func (orderID orderID) String() string {
	var values []string
	values = append(values, orderID.ClassificationID.String())
	values = append(values, orderID.MakerOwnableID.String())
	values = append(values, orderID.TakerOwnableID.String())
	values = append(values, orderID.RateID.String())
	values = append(values, orderID.CreationID.String())
	values = append(values, orderID.MakerID.String())
	values = append(values, orderID.HashID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (orderID orderID) Equals(id types.ID) bool {
	return bytes.Equal(orderID.Bytes(), id.Bytes())
}
func (orderID orderID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(orderID.Bytes())
}
func (orderID) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, orderID{})
}
func (orderID orderID) IsPartial() bool {
	return len(orderID.HashID.Bytes()) == 0
}
func (orderID orderID) Matches(key helpers.Key) bool {
	return orderID.Equals(orderIDFromInterface(key))
}

func (orderID orderID) getRateIDBytes() ([]byte, error) {
	var Bytes []byte

	if orderID.RateID.String() == "" {
		return Bytes, nil
	}

	exchangeRate, Error := sdkTypes.NewDecFromStr(orderID.RateID.String())
	if Error != nil {
		return Bytes, Error
	}

	Bytes = append(Bytes, uint8(len(strings.Split(exchangeRate.String(), ".")[0])))
	Bytes = append(Bytes, []byte(exchangeRate.String())...)

	return Bytes, Error
}

func (orderID orderID) getCreationHeightBytes() ([]byte, error) {
	var Bytes []byte

	if orderID.CreationID.String() == "" {
		return Bytes, nil
	}

	height, Error := strconv.ParseInt(orderID.CreationID.String(), 10, 64)
	if Error != nil {
		return Bytes, Error
	}

	Bytes = append(Bytes, uint8(len(orderID.CreationID.String())))
	Bytes = append(Bytes, []byte(strconv.FormatInt(height, 10))...)

	return Bytes, Error
}

func NewOrderID(classificationID test_types.ID, makerOwnableID test_types.ID, takerOwnableID test_types.ID, rateID test_types.ID, creationID test_types.ID, makerID test_types.ID, immutableProperties types.Properties) types.ID {
	return orderID{
		ClassificationID: classificationID,
		MakerOwnableID:   makerOwnableID,
		TakerOwnableID:   takerOwnableID,
		RateID:           rateID,
		CreationID:       creationID,
		MakerID:          makerID,
		HashID:           baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID(),
	}
}
