package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	simapp "github.com/persistenceOne/persistence-sdk/v2/simapp"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	atomExchangeRate        = sdk.MustNewDecFromStr("0.0000001")
	persistenceExchangeRate = sdk.MustNewDecFromStr("0.0000005")
)

func (s *KeeperTestSuite) TestOracleExportGenesis() {
	app := simapp.Setup(s.T(), false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	genesisState := oracle.ExportGenesis(ctx, app.OracleKeeper)
	params := genesisState.GetParams()
	s.Require().NotNil(params)

	expectedOracleParams := types.DefaultGenesisState().GetParams()
	s.Require().Equal(expectedOracleParams, params)
}

func (s *KeeperTestSuite) TestOracleInitGenesis() {
	app, ctx := s.app, s.ctx
	addr, valAddr, valAddr2 := s.accAddresses[0], s.valAddresses[0], s.valAddresses[1]

	currGenesisState := oracle.ExportGenesis(ctx, app.OracleKeeper)
	params := currGenesisState.GetParams()
	s.Require().NotNil(params)

	expectedOracleParams := types.DefaultGenesisState().GetParams()
	// On init genesis, default oracle information is set.
	// Confirm that the current genesis state param are default genesis state params.
	s.Require().Equal(expectedOracleParams, params)

	// new genesis state with different params and data
	newGenesisState := types.GenesisState{
		// only update params
		Params: types.Params{
			VotePeriod:               100,
			VoteThreshold:            sdk.MustNewDecFromStr("0.5"),
			RewardBand:               sdk.MustNewDecFromStr("0.05"),
			RewardDistributionWindow: 10000,
			AcceptList: types.DenomList{
				types.Denom{
					BaseDenom:   types.PersistenceDenom,
					SymbolDenom: types.PersistenceSymbol,
					Exponent:    8,
				},
				types.Denom{
					BaseDenom:   types.AtomDenom,
					SymbolDenom: types.AtomSymbol,
					Exponent:    7,
				},
			},
			SlashFraction:     sdk.MustNewDecFromStr("0.05"),
			SlashWindow:       10000,
			MinValidPerWindow: sdk.MustNewDecFromStr("0.6"),
		},
		ExchangeRates: types.ExchangeRateTuples{
			types.ExchangeRateTuple{
				Denom:        types.AtomSymbol,
				ExchangeRate: atomExchangeRate,
			},
			types.ExchangeRateTuple{
				Denom:        types.PersistenceSymbol,
				ExchangeRate: persistenceExchangeRate,
			},
		},
		MissCounters: []types.MissCounter{
			{
				ValidatorAddress: valAddr.String(),
				MissCounter:      1,
			},
		},
		FeederDelegations: []types.FeederDelegation{
			{
				ValidatorAddress: valAddr.String(),
				FeederAddress:    addr.String(),
			},
		},
		AggregateExchangeRatePrevotes: []types.AggregateExchangeRatePrevote{
			{
				Hash:  "hash1",
				Voter: valAddr.String(),
			},
			{
				Hash:  "hash2",
				Voter: valAddr2.String(),
			},
		},
		AggregateExchangeRateVotes: []types.AggregateExchangeRateVote{
			{
				ExchangeRateTuples: types.ExchangeRateTuples{
					types.ExchangeRateTuple{
						Denom:        types.AtomSymbol,
						ExchangeRate: atomExchangeRate,
					},
					types.ExchangeRateTuple{
						Denom:        types.PersistenceSymbol,
						ExchangeRate: persistenceExchangeRate,
					},
				},
				Voter: valAddr.String(),
			},
			{
				ExchangeRateTuples: types.ExchangeRateTuples{
					types.ExchangeRateTuple{
						Denom:        types.AtomSymbol,
						ExchangeRate: atomExchangeRate,
					},
					types.ExchangeRateTuple{
						Denom:        types.PersistenceSymbol,
						ExchangeRate: persistenceExchangeRate,
					},
				},
				Voter: valAddr2.String(),
			},
		},
	}

	// initialize the keeper with new genesis state and confirm that the params are set correctly
	oracle.InitGenesis(ctx, app.OracleKeeper, newGenesisState)
	newlyExportedState := oracle.ExportGenesis(ctx, app.OracleKeeper)

	s.Require().Equal(newGenesisState.GetParams(), newlyExportedState.GetParams())
	s.Require().Equal(len(newGenesisState.GetExchangeRates()), len(newlyExportedState.GetExchangeRates()))
	s.Require().Equal(len(newGenesisState.GetMissCounters()), len(newlyExportedState.GetMissCounters()))
	s.Require().Equal(len(newGenesisState.GetFeederDelegations()), len(newlyExportedState.GetFeederDelegations()))
	s.Require().EqualValues(len(newGenesisState.GetAggregateExchangeRatePrevotes()), len(newlyExportedState.GetAggregateExchangeRatePrevotes()))
	s.Require().Equal(len(newGenesisState.GetAggregateExchangeRateVotes()), len(newlyExportedState.GetAggregateExchangeRateVotes()))
}
