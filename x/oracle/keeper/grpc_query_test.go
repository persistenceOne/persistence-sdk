package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

func (s *KeeperTestSuite) TestQueryExchangeRate() {
	// Check that querying exchange rate on default genesis returns error.
	// Since no exchange rate is set.
	resp, err := s.queryClient.ExchangeRate(s.ctx.Context(), &types.QueryExchangeRateRequest{})
	s.Require().Error(err)
	s.Require().Nil(resp)

	// Set exchange rate for XPRT.
	s.app.OracleKeeper.SetExchangeRate(s.ctx, types.PersistenceSymbol, sdk.OneDec())

	resp, err = s.queryClient.ExchangeRate(s.ctx.Context(), &types.QueryExchangeRateRequest{Denom: types.PersistenceSymbol})
	s.Require().NoError(err)
	s.Require().Equal(resp.ExchangeRate, sdk.OneDec().String())
}

func (s *KeeperTestSuite) TestQueryAllExchangeRate() {
	// Check that querying all exchange rates on default genesis returns empty list.
	resp, err := s.queryClient.AllExchangeRates(s.ctx.Context(), &types.QueryAllExchangeRatesRequest{})
	s.Require().NoError(err)
	s.Require().Nil(resp.ExchangeRates)

	// Set exchange rate for XPRT.
	s.app.OracleKeeper.SetExchangeRate(s.ctx, types.PersistenceSymbol, sdk.OneDec())
	// Set exchange rate for ATOM.
	s.app.OracleKeeper.SetExchangeRate(s.ctx, types.AtomSymbol, sdk.OneDec())

	resp, err = s.queryClient.AllExchangeRates(s.ctx.Context(), &types.QueryAllExchangeRatesRequest{})
	s.Require().NoError(err)
	s.Require().Len(resp.ExchangeRates, 2)
}
