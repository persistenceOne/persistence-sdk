package keeper_test

import (
	"crypto/rand"
	"math/big"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	persistenceapp "github.com/persistenceOne/persistence-sdk/v2/simapp"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/testutil"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

const (
	exchangeRate string = persistenceapp.DisplayDenom
)

type KeeperTestSuite struct {
	suite.Suite

	valAccAddresses []sdk.AccAddress
	valAddresses    []sdk.ValAddress

	ctx sdk.Context
	app *persistenceapp.SimApp
}

const (
	rewardPoolAmount     = int64(5)
	testBalance          = int64(50000000)
	initialValidatorsNum = 2
)

func (s *KeeperTestSuite) SetupTest() {
	app := persistenceapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{
		Height: 100,
		Time:   tmtime.Now(),
	})

	var err error
	s.valAccAddresses, s.valAddresses, err = testutil.StakingAddValidators(
		app.BankKeeper,
		app.StakingKeeper,
		ctx,
		initialValidatorsNum,
	)
	s.Require().NoError(err)

	s.app = app
	s.ctx = ctx
}

func (s *KeeperTestSuite) TestSetFeederDelegation() {
	app, ctx := s.app, s.ctx
	addr, valAddr := s.valAccAddresses[0], s.valAddresses[0]

	feederAddr := sdk.AccAddress([]byte("addr________________"))
	feederAcc := app.AccountKeeper.NewAccountWithAddress(ctx, feederAddr)
	app.AccountKeeper.SetAccount(ctx, feederAcc)

	err := s.app.OracleKeeper.ValidateFeeder(ctx, valAddr, addr)
	s.Require().NoError(err)

	err = s.app.OracleKeeper.ValidateFeeder(ctx, valAddr, feederAddr)
	s.Require().ErrorContains(err, types.ErrNoVotingPermission.Error())

	s.app.OracleKeeper.SetFeederDelegation(ctx, valAddr, feederAddr)

	err = s.app.OracleKeeper.ValidateFeeder(ctx, valAddr, addr)
	s.Require().ErrorContains(err, types.ErrNoVotingPermission.Error())

	err = s.app.OracleKeeper.ValidateFeeder(ctx, valAddr, feederAddr)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) TestGetFeederDelegation() {
	app, ctx := s.app, s.ctx
	valAddr := s.valAddresses[0]

	feederAddr := sdk.AccAddress([]byte("addr________________"))
	feederAcc := app.AccountKeeper.NewAccountWithAddress(ctx, feederAddr)
	app.AccountKeeper.SetAccount(ctx, feederAcc)

	s.app.OracleKeeper.SetFeederDelegation(ctx, valAddr, feederAddr)
	resp, err := app.OracleKeeper.GetFeederDelegation(ctx, valAddr)
	s.Require().NoError(err)
	s.Require().Equal(resp, feederAddr)
}

func (s *KeeperTestSuite) TestMissCounter() {
	app, ctx := s.app, s.ctx
	valAddr := s.valAddresses[0]

	num, err := rand.Int(rand.Reader, new(big.Int).SetInt64(int64(100)))
	s.Require().NoError(err)

	missCounter := num.Uint64()

	s.Require().Equal(app.OracleKeeper.GetMissCounter(ctx, valAddr), uint64(0))
	app.OracleKeeper.SetMissCounter(ctx, valAddr, missCounter)
	s.Require().Equal(app.OracleKeeper.GetMissCounter(ctx, valAddr), missCounter)

	app.OracleKeeper.DeleteMissCounter(ctx, valAddr)
	s.Require().Equal(app.OracleKeeper.GetMissCounter(ctx, valAddr), uint64(0))
}

func (s *KeeperTestSuite) TestAggregateExchangeRatePrevote() {
	app, ctx := s.app, s.ctx
	addr, valAddr := s.valAccAddresses[0], s.valAddresses[0]

	prevote := types.AggregateExchangeRatePrevote{
		Hash:        "hash",
		Voter:       addr.String(),
		SubmitBlock: 0,
	}
	app.OracleKeeper.SetAggregateExchangeRatePrevote(ctx, valAddr, prevote)

	_, err := app.OracleKeeper.GetAggregateExchangeRatePrevote(ctx, valAddr)
	s.Require().NoError(err)

	app.OracleKeeper.DeleteAggregateExchangeRatePrevote(ctx, valAddr)

	_, err = app.OracleKeeper.GetAggregateExchangeRatePrevote(ctx, valAddr)
	s.Require().Error(err)
}

func (s *KeeperTestSuite) TestAggregateExchangeRatePrevoteError() {
	app, ctx := s.app, s.ctx
	valAddr := s.valAddresses[0]

	_, err := app.OracleKeeper.GetAggregateExchangeRatePrevote(ctx, valAddr)
	s.Require().Errorf(err, types.ErrNoAggregatePrevote.Error())
}

func (s *KeeperTestSuite) TestAggregateExchangeRateVote() {
	app, ctx := s.app, s.ctx
	addr, valAddr := s.valAccAddresses[0], s.valAddresses[0]

	var tuples types.ExchangeRateTuples
	tuples = append(tuples, types.ExchangeRateTuple{
		Denom:        types.PersistenceDenom,
		ExchangeRate: sdk.ZeroDec(),
	})

	vote := types.AggregateExchangeRateVote{
		ExchangeRateTuples: tuples,
		Voter:              addr.String(),
	}
	app.OracleKeeper.SetAggregateExchangeRateVote(ctx, valAddr, vote)

	_, err := app.OracleKeeper.GetAggregateExchangeRateVote(ctx, valAddr)
	s.Require().NoError(err)

	app.OracleKeeper.DeleteAggregateExchangeRateVote(ctx, valAddr)

	_, err = app.OracleKeeper.GetAggregateExchangeRateVote(ctx, valAddr)
	s.Require().Error(err)
}

func (s *KeeperTestSuite) TestAggregateExchangeRateVoteError() {
	app, ctx := s.app, s.ctx
	valAddr := s.valAddresses[0]

	_, err := app.OracleKeeper.GetAggregateExchangeRateVote(ctx, valAddr)
	s.Require().Errorf(err, types.ErrNoAggregateVote.Error())
}

func (s *KeeperTestSuite) TestSetExchangeRateWithEvent() {
	app, ctx := s.app, s.ctx
	app.OracleKeeper.SetExchangeRateWithEvent(ctx, types.PersistenceDenom, sdk.OneDec())
	rate, err := app.OracleKeeper.GetExchangeRate(ctx, types.PersistenceDenom)
	s.Require().NoError(err)
	s.Require().Equal(rate, sdk.OneDec())
}

func (s *KeeperTestSuite) TestGetExchangeRate_UnknownDenom() {
	app, ctx := s.app, s.ctx

	_, err := app.OracleKeeper.GetExchangeRate(ctx, "uxyz")
	s.Require().ErrorContains(err, types.ErrUnknownDenom.Error())
}

func (s *KeeperTestSuite) TestGetExchangeRate_NotSet() {
	app, ctx := s.app, s.ctx

	_, err := app.OracleKeeper.GetExchangeRate(ctx, types.PersistenceDenom)
	s.Require().Error(err)
}

func (s *KeeperTestSuite) TestGetExchangeRate_Valid() {
	app, ctx := s.app, s.ctx

	app.OracleKeeper.SetExchangeRate(ctx, types.PersistenceDenom, sdk.OneDec())
	rate, err := app.OracleKeeper.GetExchangeRate(ctx, types.PersistenceDenom)
	s.Require().NoError(err)
	s.Require().Equal(rate, sdk.OneDec())

	app.OracleKeeper.SetExchangeRate(ctx, strings.ToLower(types.PersistenceDenom), sdk.OneDec())
	rate, err = app.OracleKeeper.GetExchangeRate(ctx, types.PersistenceDenom)
	s.Require().NoError(err)
	s.Require().Equal(rate, sdk.OneDec())
}

func (s *KeeperTestSuite) TestDeleteExchangeRate() {
	app, ctx := s.app, s.ctx

	app.OracleKeeper.SetExchangeRate(ctx, types.PersistenceDenom, sdk.OneDec())
	app.OracleKeeper.DeleteExchangeRate(ctx, types.PersistenceDenom)
	_, err := app.OracleKeeper.GetExchangeRate(ctx, types.PersistenceDenom)
	s.Require().Error(err)
}

func (s *KeeperTestSuite) balanceSetup() {
	app, ctx := s.app, s.ctx
	addr := s.valAccAddresses[0]

	// Prepare account balance
	givingAmt := sdk.NewCoins(sdk.NewInt64Coin(types.PersistenceDenom, testBalance))
	err := app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, givingAmt)
	s.Require().NoError(err)

	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, givingAmt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) TestFundRewardPool() {
	s.balanceSetup()
	app, ctx := s.app, s.ctx
	addr := s.valAccAddresses[0]

	// Fund reward pool form account
	coins := sdk.NewCoins(sdk.NewInt64Coin(types.PersistenceDenom, rewardPoolAmount))
	err := app.OracleKeeper.FundRewardPool(ctx, addr, coins)
	s.Require().NoError(err)

	moduleAddr := app.AccountKeeper.GetModuleAddress(types.ModuleName)
	balance := app.BankKeeper.GetAllBalances(ctx, moduleAddr)
	denomAmount := balance.AmountOf(types.PersistenceDenom)
	s.Require().Equal(denomAmount.Int64(), rewardPoolAmount)
}

func (s *KeeperTestSuite) TestGetRewardPoolBalance() {
	s.balanceSetup()
	app, ctx := s.app, s.ctx
	addr := s.valAccAddresses[0]

	// Fund reward pool form account
	coins := sdk.NewCoins(sdk.NewInt64Coin(types.PersistenceDenom, rewardPoolAmount))
	err := app.OracleKeeper.FundRewardPool(ctx, addr, coins)
	s.Require().NoError(err)

	moduleAddr := app.AccountKeeper.GetModuleAddress(types.ModuleName)
	balance := app.OracleKeeper.GetRewardPoolBalance(ctx, moduleAddr)
	denomAmount := balance.AmountOf(types.PersistenceDenom)
	s.Require().Equal(denomAmount.Int64(), rewardPoolAmount)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
