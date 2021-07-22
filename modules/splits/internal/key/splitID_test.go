/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_SplitID_Methods(t *testing.T) {

	ownerID := testBase.NewID("ownerID")
	ownableID := testBase.NewID("ownableID")

	testSplitID := NewSplitID(ownerID, ownableID).(splitID)
	testSplitID2 := NewSplitID(testBase.NewID(""), testBase.NewID("")).(splitID)
	require.NotPanics(t, func() {
		require.Equal(t, strings.Join([]string{ownerID.String(), ownableID.String()}, constants.SecondOrderCompositeIDSeparator), testSplitID.String())
		require.Equal(t, true, testSplitID.Equals(testSplitID))
		require.Equal(t, false, testSplitID.Equals(testSplitID2))
		require.Equal(t, false, testSplitID.IsPartial())
		require.Equal(t, true, testSplitID2.IsPartial())

		require.Equal(t, true, testSplitID.Matches(testSplitID))
		require.Equal(t, false, testSplitID.Matches(testSplitID2))
		require.Equal(t, false, testSplitID.Matches(nil))
		require.Equal(t, testSplitID, FromID(testSplitID))
		require.Equal(t, testSplitID2, FromID(testBase.NewID("")))
		require.Equal(t, splitID{OwnerID: testBase.NewID("ID1"), OwnableID: testBase.NewID("ID2")}, readSplitID("ID1*ID2"))
	})
}
