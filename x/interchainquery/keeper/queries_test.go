package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/persistenceOne/persistence-sdk/x/interchainquery/keeper"
)

func (s *KeeperTestSuite) TestQuery() {
	bondedQuery := stakingtypes.QueryValidatorsRequest{Status: stakingtypes.BondStatusBonded}
	bz1, err := bondedQuery.Marshal()
	s.NoError(err)

	query := s.GetSimApp(s.chainA).InterchainQueryKeeper.NewQuery(
		s.chainA.GetContext(),
		"",
		s.path.EndpointB.ConnectionID,
		s.chainB.ChainID,
		"cosmos.staking.v1beta1.Query/Validators",
		bz1,
		sdk.NewInt(200),
		"",
		0,
	)

	// set the query
	s.GetSimApp(s.chainA).InterchainQueryKeeper.SetQuery(s.chainA.GetContext(), *query)

	// get the stored query
	id := keeper.GenerateQueryHash(query.ConnectionId, query.ChainId, query.QueryType, query.Request, "")
	query1, found := s.GetSimApp(s.chainA).InterchainQueryKeeper.GetQuery(s.chainA.GetContext(), id)
	s.True(found)
	s.Equal(s.path.EndpointB.ConnectionID, query1.ConnectionId)
	s.Equal(s.chainB.ChainID, query1.ChainId)
	s.Equal("cosmos.staking.v1beta1.Query/Validators", query1.QueryType)
	s.Equal(sdk.NewInt(200), query1.Period)
	s.Equal(uint64(0), query1.Ttl)
	s.Equal("", query1.CallbackId)

	// get all the queries
	queries := s.GetSimApp(s.chainA).InterchainQueryKeeper.AllQueries(s.chainA.GetContext())
	s.Len(queries, 1)

	// delete the query
	s.GetSimApp(s.chainA).InterchainQueryKeeper.DeleteQuery(s.chainA.GetContext(), id)

	// get query
	_, found2 := s.GetSimApp(s.chainA).InterchainQueryKeeper.GetQuery(s.chainA.GetContext(), id)
	s.False(found2)

	queries1 := s.GetSimApp(s.chainA).InterchainQueryKeeper.AllQueries(s.chainA.GetContext())
	s.Len(queries1, 0)
}
