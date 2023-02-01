package keeper_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	persistenceapp "github.com/persistenceOne/persistence-sdk/v2/simapp"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/keeper"
	"github.com/persistenceOne/persistence-sdk/v2/x/oracle/types"
)

const (
	exchangeRate string = persistenceapp.DisplayDenom
)

type IntegrationTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *persistenceapp.SimApp
	queryClient types.QueryClient
	msgServer   types.MsgServer
}

const (
	initialPower     = int64(10000000000)
	rewardPoolAmount = int64(5)
	testBalance      = int64(50000000)
)

func (s *IntegrationTestSuite) SetupTest() {
	app := persistenceapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{
		ChainID: fmt.Sprintf("test-chain-%s", tmrand.Str(4)),
		Height:  9,
	})

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, keeper.NewQuerier(app.OracleKeeper))

	sh := staking.NewHandler(app.StakingKeeper)
	amt := sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)

	// mint and send coins to validators
	s.Require().NoError(app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, initCoins))
	s.Require().NoError(app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, initCoins))
	s.Require().NoError(app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, initCoins))
	s.Require().NoError(app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr2, initCoins))

	_, err := sh(ctx, NewTestMsgCreateValidator(valAddr, valPubKey, amt))
	s.Require().NoError(err)
	_, err = sh(ctx, NewTestMsgCreateValidator(valAddr2, valPubKey2, amt))
	s.Require().NoError(err)

	staking.EndBlocker(ctx, app.StakingKeeper)

	s.app = app
	s.ctx = ctx
	s.queryClient = types.NewQueryClient(queryHelper)
	s.msgServer = keeper.NewMsgServerImpl(app.OracleKeeper)
}

// Test addresses
var (
	valPubKeys = simapp.CreateTestPubKeys(2)

	valPubKey = valPubKeys[0]
	pubKey    = secp256k1.GenPrivKey().PubKey()
	addr      = sdk.AccAddress(pubKey.Address())
	valAddr   = sdk.ValAddress(pubKey.Address())

	valPubKey2 = valPubKeys[1]
	pubKey2    = secp256k1.GenPrivKey().PubKey()
	addr2      = sdk.AccAddress(pubKey2.Address())
	valAddr2   = sdk.ValAddress(pubKey2.Address())

	initTokens = sdk.TokensFromConsensusPower(initialPower, sdk.DefaultPowerReduction)
	initCoins  = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, initTokens))
)

// NewTestMsgCreateValidator test msg creator
func NewTestMsgCreateValidator(address sdk.ValAddress, pubKey cryptotypes.PubKey, amt sdk.Int) *stakingtypes.MsgCreateValidator {
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
	msg, _ := stakingtypes.NewMsgCreateValidator(
		address, pubKey, sdk.NewCoin(sdk.DefaultBondDenom, amt),
		stakingtypes.Description{}, commission, sdk.OneInt(),
	)

	return msg
}

func (s *IntegrationTestSuite) TestSetFeederDelegation() {
	app, ctx := s.app, s.ctx

	feederAddr := sdk.AccAddress([]byte("addr________________"))
	feederAcc := app.AccountKeeper.NewAccountWithAddress(ctx, feederAddr)
	app.AccountKeeper.SetAccount(ctx, feederAcc)

	err := s.app.OracleKeeper.ValidateFeeder(ctx, addr, valAddr)
	s.Require().NoError(err)
	err = s.app.OracleKeeper.ValidateFeeder(ctx, feederAddr, valAddr)
	s.Require().Error(err)

	s.app.OracleKeeper.SetFeederDelegation(ctx, valAddr, feederAddr)

	err = s.app.OracleKeeper.ValidateFeeder(ctx, addr, valAddr)
	s.Require().Error(err)
	err = s.app.OracleKeeper.ValidateFeeder(ctx, feederAddr, valAddr)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestGetFeederDelegation() {
	app, ctx := s.app, s.ctx

	feederAddr := sdk.AccAddress([]byte("addr________________"))
	feederAcc := app.AccountKeeper.NewAccountWithAddress(ctx, feederAddr)
	app.AccountKeeper.SetAccount(ctx, feederAcc)

	s.app.OracleKeeper.SetFeederDelegation(ctx, valAddr, feederAddr)
	resp, err := app.OracleKeeper.GetFeederDelegation(ctx, valAddr)
	s.Require().NoError(err)
	s.Require().Equal(resp, feederAddr)
}

func (s *IntegrationTestSuite) TestMissCounter() {
	app, ctx := s.app, s.ctx
	num, err := rand.Int(rand.Reader, new(big.Int).SetInt64(int64(100)))
	s.Require().NoError(err)

	missCounter := num.Uint64()

	s.Require().Equal(app.OracleKeeper.GetMissCounter(ctx, valAddr), uint64(0))
	app.OracleKeeper.SetMissCounter(ctx, valAddr, missCounter)
	s.Require().Equal(app.OracleKeeper.GetMissCounter(ctx, valAddr), missCounter)

	app.OracleKeeper.DeleteMissCounter(ctx, valAddr)
	s.Require().Equal(app.OracleKeeper.GetMissCounter(ctx, valAddr), uint64(0))
}

func (s *IntegrationTestSuite) TestAggregateExchangeRatePrevote() {
	app, ctx := s.app, s.ctx

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

func (s *IntegrationTestSuite) TestAggregateExchangeRatePrevoteError() {
	app, ctx := s.app, s.ctx

	_, err := app.OracleKeeper.GetAggregateExchangeRatePrevote(ctx, valAddr)
	s.Require().Errorf(err, types.ErrNoAggregatePrevote.Error())
}

func (s *IntegrationTestSuite) TestAggregateExchangeRateVote() {
	app, ctx := s.app, s.ctx

	var tuples types.ExchangeRateTuples
	tuples = append(tuples, types.ExchangeRateTuple{
		Denom:        exchangeRate,
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

func (s *IntegrationTestSuite) TestAggregateExchangeRateVoteError() {
	app, ctx := s.app, s.ctx

	_, err := app.OracleKeeper.GetAggregateExchangeRateVote(ctx, valAddr)
	s.Require().Errorf(err, types.ErrNoAggregateVote.Error())
}

func (s *IntegrationTestSuite) TestSetExchangeRateWithEvent() {
	app, ctx := s.app, s.ctx
	app.OracleKeeper.SetExchangeRateWithEvent(ctx, exchangeRate, sdk.OneDec())
	rate, err := app.OracleKeeper.GetExchangeRate(ctx, exchangeRate)
	s.Require().NoError(err)
	s.Require().Equal(rate, sdk.OneDec())
}

func (s *IntegrationTestSuite) TestGetExchangeRate_InvalidDenom() {
	app, ctx := s.app, s.ctx

	_, err := app.OracleKeeper.GetExchangeRate(ctx, "uxyz")
	s.Require().Error(err)
}

func (s *IntegrationTestSuite) TestGetExchangeRate_NotSet() {
	app, ctx := s.app, s.ctx

	_, err := app.OracleKeeper.GetExchangeRate(ctx, exchangeRate)
	s.Require().Error(err)
}

func (s *IntegrationTestSuite) TestGetExchangeRate_Valid() {
	app, ctx := s.app, s.ctx

	app.OracleKeeper.SetExchangeRate(ctx, exchangeRate, sdk.OneDec())
	rate, err := app.OracleKeeper.GetExchangeRate(ctx, exchangeRate)
	s.Require().NoError(err)
	s.Require().Equal(rate, sdk.OneDec())

	app.OracleKeeper.SetExchangeRate(ctx, strings.ToLower(exchangeRate), sdk.OneDec())
	rate, err = app.OracleKeeper.GetExchangeRate(ctx, exchangeRate)
	s.Require().NoError(err)
	s.Require().Equal(rate, sdk.OneDec())
}

func (s *IntegrationTestSuite) TestDeleteExchangeRate() {
	app, ctx := s.app, s.ctx

	app.OracleKeeper.SetExchangeRate(ctx, exchangeRate, sdk.OneDec())
	app.OracleKeeper.DeleteExchangeRate(ctx, exchangeRate)
	_, err := app.OracleKeeper.GetExchangeRate(ctx, exchangeRate)
	s.Require().Error(err)
}

func (s *IntegrationTestSuite) balanceSetup() {
	// Prepare account balance
	givingAmt := sdk.NewCoins(sdk.NewInt64Coin(types.PersistenceDenom, testBalance))
	err := s.app.BankKeeper.MintCoins(s.ctx, minttypes.ModuleName, givingAmt)
	s.Require().NoError(err)

	err = s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, minttypes.ModuleName, addr, givingAmt)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestFundRewardPool() {
	s.balanceSetup()
	app, ctx := s.app, s.ctx

	// Fund reward pool form account
	coins := sdk.NewCoins(sdk.NewInt64Coin(types.PersistenceDenom, rewardPoolAmount))
	err := app.OracleKeeper.FundRewardPool(ctx, addr, coins)
	s.Require().NoError(err)

	moduleAddr := app.AccountKeeper.GetModuleAddress(types.ModuleName)
	balance := app.BankKeeper.GetAllBalances(ctx, moduleAddr)
	denomAmount := balance.AmountOf(types.PersistenceDenom)
	s.Require().Equal(denomAmount.Int64(), rewardPoolAmount)
}

func (s *IntegrationTestSuite) TestGetRewardPoolBalance() {
	s.balanceSetup()
	app, ctx := s.app, s.ctx

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
	suite.Run(t, new(IntegrationTestSuite))
}
