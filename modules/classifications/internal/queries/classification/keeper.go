/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package classification

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {

	classifications := queryKeeper.mapper.NewCollection(context).Fetch(key.FromID(queryRequestFromInterface(queryRequest).ClassificationID))
	classification := classifications.Get(key.FromID(queryRequestFromInterface(queryRequest).ClassificationID))
	if classification == nil {
		return newQueryResponse(nil, errors.EntityNotFound)
	}

	return newQueryResponse(classification, nil)
	//return newQueryResponse(queryKeeper.mapper.NewCollection(context).Fetch(key.FromID(queryRequestFromInterface(queryRequest).ClassificationID)), nil)
}

func (queryKeeper queryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	queryKeeper.mapper = mapper
	return queryKeeper
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}
