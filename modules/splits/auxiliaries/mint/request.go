/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mint

import (
	"fmt"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"

	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryRequest struct {
	OwnerID   test_types.ID `json:"ownerID" valid:"required~required field ownerID missing"`
	OwnableID test_types.ID `json:"ownableID" valid:"required~required field ownableID missing"`
	Value     sdkTypes.Dec  `json:"value" valid:"required~required field value missing"`
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

func NewAuxiliaryRequest(ownerID fmt.Stringer, ownableID fmt.Stringer, value sdkTypes.Dec) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		OwnerID:   test_types.NewID(ownerID.String()),
		OwnableID: test_types.NewID(ownableID.String()),
		Value:     value,
	}
}
