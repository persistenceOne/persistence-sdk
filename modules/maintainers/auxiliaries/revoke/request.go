/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package revoke

import (
	"github.com/asaskevich/govalidator"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

type auxiliaryRequest struct {
	FromID           test_types.ID `json:"fromID" valid:"required~required field fromID missing"`
	ToID             test_types.ID `json:"toID" valid:"required~required field toID missing"`
	ClassificationID test_types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
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

func NewAuxiliaryRequest(fromID test_types.ID, toID test_types.ID, classificationID test_types.ID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:           fromID,
		ToID:             toID,
		ClassificationID: classificationID,
	}
}
