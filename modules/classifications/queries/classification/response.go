package classification

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
)

type queryResponse struct {
	Classifications mappers.Classifications `json:"classifications" valid:"required~required field classifications missing"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() helpers.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(classifications mappers.Classifications) helpers.QueryResponse {
	return queryResponse{Classifications: classifications}
}