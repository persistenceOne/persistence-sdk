/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package super

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	ClassificationID  protoTypes.ID    `json:"classificationID" valid:"required~required field classificationID missing"`
	IdentityID        protoTypes.ID    `json:"identityID" valid:"required~required field identityID missing"`
	MutableProperties types.Properties `json:"mutableProperties" valid:"required~required field mutableProperties missing"`
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

func NewAuxiliaryRequest(classificationID protoTypes.ID, identityID protoTypes.ID, mutableProperties types.Properties) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID:  classificationID,
		IdentityID:        identityID,
		MutableProperties: mutableProperties,
	}
}
