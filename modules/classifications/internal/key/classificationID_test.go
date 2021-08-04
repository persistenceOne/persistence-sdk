/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	"strings"
	"testing"

	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"
)

func Test_ClassificationID_Methods(t *testing.T) {
	chainID := testBase.NewID("chainID")
	immutableProperties := base.NewProperties(base.NewProperty(testBase.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))
	mutableProperties := base.NewProperties(base.NewProperty(testBase.NewID("ID2"), base.NewFact(base.NewStringData("MutableData"))))

	testClassificationID := NewClassificationID(chainID, immutableProperties, mutableProperties).(classificationID)
	require.NotPanics(t, func() {
		require.Equal(t, classificationID{ChainID: chainID, HashID: testBase.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().String()))}, testClassificationID)
		require.Equal(t, strings.Join([]string{chainID.String(), base.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().String())).String()}, constants.IDSeparator), testClassificationID.String())
		require.Equal(t, false, testClassificationID.Matches(classificationID{ChainID: testBase.NewID("chainID"), HashID: testBase.NewID("hashID")}))
		require.Equal(t, false, testClassificationID.Matches(nil))
		require.Equal(t, false, testClassificationID.Equals(testBase.NewID("id")))
		require.Equal(t, true, testClassificationID.Equals(testClassificationID))
		require.Equal(t, false, testClassificationID.IsPartial())
		require.Equal(t, true, classificationID{ChainID: chainID, HashID: testBase.NewID("")}.IsPartial())
		require.Equal(t, testClassificationID, FromID(testClassificationID))
		require.Equal(t, classificationID{ChainID: testBase.NewID(""), HashID: testBase.NewID("")}, FromID(testBase.NewID("tempID")))
		require.Equal(t, testClassificationID, readClassificationID(testClassificationID.String()))
	})

}
