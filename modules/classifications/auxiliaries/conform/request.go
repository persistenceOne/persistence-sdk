/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package conform

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	ClassificationID    protoTypes.ID    `json:"classificationID" valid:"required~required field classificationID missing"`
	ImmutableProperties types.Properties `json:"immutableProperties"`
	MutableProperties   types.Properties `json:"mutableProperties"`
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(auxiliaryRequest)
	return Error
}

func auxiliaryRequestFromInterface(request helpers.AuxiliaryRequest) auxiliaryRequest {
	switch value := request.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(classificationID protoTypes.ID, immutableProperties types.Properties, mutableProperties types.Properties) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID:    classificationID,
		ImmutableProperties: immutableProperties,
		MutableProperties:   mutableProperties,
	}
}
