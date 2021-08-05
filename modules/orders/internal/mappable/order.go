/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	base2 "github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type order struct {
	ID protoTypes.ID `json:"id" valid:"required~required field key missing"`
	traits.HasMutables
	traits.HasImmutables
}

var _ mappables.Order = (*order)(nil)

func (order order) GetID() protoTypes.ID {
	return order.ID
}
func (order order) GetClassificationID() protoTypes.ID {
	return key.ReadClassificationID(order.ID)
}
func (order order) GetRateID() protoTypes.ID {
	return key.ReadRateID(order.ID)
}
func (order order) GetCreationID() protoTypes.ID {
	return key.ReadCreationID(order.ID)
}
func (order order) GetMakerOwnableID() protoTypes.ID {
	return key.ReadMakerOwnableID(order.ID)
}
func (order order) GetTakerOwnableID() protoTypes.ID {
	return key.ReadTakerOwnableID(order.ID)
}
func (order order) GetMakerID() protoTypes.ID {
	return key.ReadMakerID(order.ID)
}
func (order order) GetCreation() types.MetaProperty {
	heightValue, Error := strconv.ParseInt(key.ReadCreationID(order.ID).String(), 10, 64)
	if Error != nil {
		return base.NewMetaProperty(base.NewID(properties.Creation), base.NewMetaFact(base.NewHeightData(base.NewHeight(0))))
	}

	return base.NewMetaProperty(base.NewID(properties.Creation), base.NewMetaFact(base.NewHeightData(base.NewHeight(heightValue))))
}
func (order order) GetExchangeRate() types.MetaProperty {
	decValue, Error := sdkTypes.NewDecFromStr(key.ReadRateID(order.ID).String())
	if Error != nil {
		return base.NewMetaProperty(base.NewID(properties.ExchangeRate), base.NewMetaFact(base.NewDecData(sdkTypes.ZeroDec())))
	}

	return base.NewMetaProperty(base.NewID(properties.ExchangeRate), base.NewMetaFact(base.NewDecData(decValue)))
}
func (order order) GetTakerID() types.Property {
	if takerID := order.HasImmutables.GetImmutableProperties().Get(base2.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else if takerID := order.HasMutables.GetMutableProperties().Get(base2.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else {
		data, _ := base.ReadIDData("")
		return base.NewProperty(base2.NewID(properties.TakerID), base.NewFact(data))
	}
}
func (order order) GetExpiry() types.Property {
	if property := order.HasImmutables.GetImmutableProperties().Get(base2.NewID(properties.Expiry)); property != nil {
		return property
	} else if property := order.HasMutables.GetMutableProperties().Get(base2.NewID(properties.Expiry)); property != nil {
		return property
	} else {
		return base.NewProperty(base2.NewID(properties.Expiry), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (order order) GetMakerOwnableSplit() types.Property {
	if split := order.HasImmutables.GetImmutableProperties().Get(base2.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else if split := order.HasMutables.GetMutableProperties().Get(base2.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else {
		data, _ := base.ReadDecData("")
		return base.NewProperty(base2.NewID(properties.MakerOwnableSplit), base.NewFact(data))
	}
}
func (order order) GetKey() helpers.Key {
	return key.FromID(order.ID)
}
func (order) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, order{})
}

func NewOrder(orderID protoTypes.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Order {
	return order{
		ID:            orderID,
		HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties},
		HasMutables:   baseTraits.HasMutables{Properties: mutableProperties},
	}
}
