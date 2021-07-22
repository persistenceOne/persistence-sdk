/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type asset struct {
	ID test_types.ID `json:"id" valid:"required~required field id missing"`
	baseTraits.HasImmutables
	baseTraits.HasMutables //nolint:govet
}

var _ mappables.InterNFT = (*asset)(nil)

func (asset asset) GetID() test_types.ID {
	return asset.ID
}
func (asset asset) GetClassificationID() test_types.ID {
	return key.ReadClassificationID(asset.ID)
}
func (asset asset) GetBurn() types.Property {
	if burnProperty := asset.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.Burn)); burnProperty != nil {
		return burnProperty
	} else if burnProperty := asset.HasMutables.GetMutableProperties().Get(base.NewID(properties.Burn)); burnProperty != nil {
		return burnProperty
	} else {
		return base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (asset asset) GetLock() types.Property {
	if lockProperty := asset.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.Lock)); lockProperty != nil {
		return lockProperty
	} else if lockProperty := asset.HasMutables.GetMutableProperties().Get(base.NewID(properties.Lock)); lockProperty != nil {
		return lockProperty
	} else {
		return base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	}
}
func (asset asset) GetValue() types.Property {
	if splitProperty := asset.HasImmutables.GetImmutableProperties().Get(base.NewID(properties.Value)); splitProperty != nil {
		return splitProperty
	} else if splitProperty := asset.HasMutables.GetMutableProperties().Get(base.NewID(properties.Value)); splitProperty != nil {
		return splitProperty
	} else {
		return base.NewProperty(base.NewID(properties.Value), base.NewFact(base.NewDecData(sdkTypes.SmallestDec())))
	}
}
func (asset asset) GetKey() helpers.Key {
	return key.FromID(asset.ID)
}
func (asset) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, asset{})
}

func NewAsset(assetID test_types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.InterNFT {
	return asset{
		ID:            assetID,
		HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties},
		HasMutables:   baseTraits.HasMutables{Properties: mutableProperties},
	}
}
