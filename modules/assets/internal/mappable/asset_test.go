/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Asset_Methods(t *testing.T) {
	classificationID := testBase.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(testBase.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))
	mutableProperties := base.NewProperties(base.NewProperty(testBase.NewID("ID2"), base.NewFact(base.NewStringData("MutableData"))))

	assetID := key.NewAssetID(classificationID, immutableProperties)
	testAsset := NewAsset(assetID, immutableProperties, mutableProperties).(asset)

	require.Equal(t, asset{ID: assetID, HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties}, HasMutables: baseTraits.HasMutables{Properties: mutableProperties}}, testAsset)
	require.Equal(t, assetID, testAsset.GetID())
	require.Equal(t, classificationID, testAsset.GetClassificationID())
	require.Equal(t, immutableProperties, testAsset.GetImmutableProperties())
	require.Equal(t, mutableProperties, testAsset.GetMutableProperties())
	data, _ := base.ReadHeightData("-1")
	require.Equal(t, testAsset.GetBurn(), base.NewProperty(testBase.NewID(properties.Burn), base.NewFact(data)))
	require.Equal(t, base.NewProperty(testBase.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnImmutableData"))), asset{ID: assetID, HasImmutables: baseTraits.HasImmutables{Properties: base.NewProperties(base.NewProperty(testBase.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnImmutableData"))))}, HasMutables: baseTraits.HasMutables{Properties: mutableProperties}}.GetBurn())
	require.Equal(t, base.NewProperty(testBase.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnMutableData"))), asset{ID: assetID, HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties}, HasMutables: baseTraits.HasMutables{Properties: base.NewProperties(base.NewProperty(testBase.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnMutableData"))))}}.GetBurn())
	require.Equal(t, base.NewProperty(testBase.NewID(properties.Lock), base.NewFact(data)), testAsset.GetLock())
	require.Equal(t, base.NewProperty(testBase.NewID(properties.Lock), base.NewFact(base.NewStringData("LockImmutableData"))), asset{ID: assetID, HasImmutables: baseTraits.HasImmutables{Properties: base.NewProperties(base.NewProperty(testBase.NewID(properties.Lock), base.NewFact(base.NewStringData("LockImmutableData"))))}, HasMutables: baseTraits.HasMutables{Properties: mutableProperties}}.GetLock())
	require.Equal(t, base.NewProperty(testBase.NewID(properties.Lock), base.NewFact(base.NewStringData("LockMutableData"))), asset{ID: assetID, HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties}, HasMutables: baseTraits.HasMutables{Properties: base.NewProperties(base.NewProperty(testBase.NewID(properties.Lock), base.NewFact(base.NewStringData("LockMutableData"))))}}.GetLock())
	require.Equal(t, assetID, testAsset.GetKey())

}
