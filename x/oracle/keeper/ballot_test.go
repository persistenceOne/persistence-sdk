package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

func (s *KeeperTestSuite) TestBallot_OrganizeBallotByDenom() {
	valAddr := s.valAddresses[0]
	s.app.OracleKeeper.SetExchangeRate(s.ctx, types.PersistenceDenom, sdk.OneDec())

	// Empty Map
	claimMap := make(map[string]types.Claim)
	res := s.app.OracleKeeper.OrganizeBallotByDenom(s.ctx, claimMap)
	s.Require().Empty(res)

	s.app.OracleKeeper.SetAggregateExchangeRateVote(
		s.ctx, valAddr, types.AggregateExchangeRateVote{
			ExchangeRateTuples: types.ExchangeRateTuples{
				types.ExchangeRateTuple{
					Denom:        "XPRT",
					ExchangeRate: sdk.OneDec(),
				},
			},
			Voter: valAddr.String(),
		},
	)

	claimMap[valAddr.String()] = types.Claim{
		Power:     1,
		Weight:    1,
		WinCount:  1,
		Recipient: valAddr,
	}
	res = s.app.OracleKeeper.OrganizeBallotByDenom(s.ctx, claimMap)
	s.Require().Equal([]types.BallotDenom{
		{
			Ballot: types.ExchangeRateBallot{types.NewVoteForTally(sdk.OneDec(), "XPRT", valAddr, 1)},
			Denom:  "XPRT",
		},
	}, res)
}

func (s *KeeperTestSuite) TestBallot_ClearBallots() {
	addr, valAddr := s.accAddresses[0], s.valAddresses[0]

	prevote := types.AggregateExchangeRatePrevote{
		Hash:        "hash",
		Voter:       addr.String(),
		SubmitBlock: 0,
	}
	s.app.OracleKeeper.SetAggregateExchangeRatePrevote(s.ctx, valAddr, prevote)
	prevoteRes, err := s.app.OracleKeeper.GetAggregateExchangeRatePrevote(s.ctx, valAddr)
	s.Require().NoError(err)
	s.Require().Equal(prevoteRes, prevote)

	var tuples types.ExchangeRateTuples
	tuples = append(tuples, types.ExchangeRateTuple{
		Denom:        "XPRT",
		ExchangeRate: sdk.ZeroDec(),
	})
	vote := types.AggregateExchangeRateVote{
		ExchangeRateTuples: tuples,
		Voter:              addr.String(),
	}
	s.app.OracleKeeper.SetAggregateExchangeRateVote(s.ctx, valAddr, vote)
	voteRes, err := s.app.OracleKeeper.GetAggregateExchangeRateVote(s.ctx, valAddr)
	s.Require().NoError(err)
	s.Require().Equal(voteRes, vote)

	s.app.OracleKeeper.ClearVotes(s.ctx, 0)
	_, err = s.app.OracleKeeper.GetAggregateExchangeRatePrevote(s.ctx, valAddr)
	s.Require().Error(err)
	_, err = s.app.OracleKeeper.GetAggregateExchangeRateVote(s.ctx, valAddr)
	s.Require().Error(err)
}
