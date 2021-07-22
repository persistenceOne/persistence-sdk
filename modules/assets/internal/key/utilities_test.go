package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	testBase"github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	typesBase"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromID(t *testing.T) {

	classificationID := testBase.NewID("classificationID")
	immutableProperties := typesBase.NewProperties(typesBase.NewProperty(testBase.NewID("ID1"), typesBase.NewFact(typesBase.NewStringData("ImmutableData"))))
	newAssetID := NewAssetID(classificationID, immutableProperties)
	require.Equal(t, assetIDFromInterface(newAssetID), FromID(newAssetID))

	id := testBase.NewID("")
	testAssetID := assetID{ClassificationID: testBase.NewID(""), HashID: testBase.NewID("")}
	require.Equal(t, FromID(id), testAssetID)

	testString1 := "string1"
	testString2 := "string2"
	id2 := testBase.NewID(testString1 + constants.FirstOrderCompositeIDSeparator + testString2)
	testAssetID2 := assetID{ClassificationID: testBase.NewID(testString1), HashID: testBase.NewID(testString2)}
	require.Equal(t, FromID(id2), testAssetID2)
}

func TestReadClassificationID(t *testing.T) {
	classificationID := testBase.NewID("classificationID")
	immutableProperties := typesBase.NewProperties(typesBase.NewProperty(testBase.NewID("ID1"), typesBase.NewFact(typesBase.NewStringData("ImmutableData"))))
	assetID := NewAssetID(classificationID, immutableProperties)

	require.Equal(t, assetIDFromInterface(assetID).ClassificationID, ReadClassificationID(assetID))

}
