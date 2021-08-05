/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package conform

import (
	base2 "github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	"testing"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Conform_Request(t *testing.T) {

	classificationID := base2.NewID("classificationID")
	mutableProperties := base.NewProperties(base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))))
	immutableProperties := base.NewProperties(base.NewProperty(base2.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, immutableProperties, mutableProperties)

	require.Equal(t, auxiliaryRequest{ClassificationID: classificationID, ImmutableProperties: immutableProperties, MutableProperties: mutableProperties}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
