/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package maintain

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	ClassificationID     protoTypes.ID    `json:"classificationID" valid:"required~required field classificationID missing"`
	IdentityID           protoTypes.ID    `json:"identityID" valid:"required~required field identityID missing"`
	MaintainedProperties types.Properties `json:"maintainedProperties" valid:"required~required field maintainedProperties missing"`
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

func NewAuxiliaryRequest(classificationID protoTypes.ID, identityID protoTypes.ID, maintainedProperties types.Properties) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID:     classificationID,
		IdentityID:           identityID,
		MaintainedProperties: maintainedProperties,
	}
}
