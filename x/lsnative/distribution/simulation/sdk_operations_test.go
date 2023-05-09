package simulation_test

import (
	"math/rand"

	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"
	sdkdistr "github.com/cosmos/cosmos-sdk/x/distribution/types"
	simappparams "github.com/persistenceOne/persistence-sdk/v2/simapp/params"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/distribution/simulation"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/distribution/types"
	distrtypes "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/distribution/types"
	stakingtypes "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"
)

// TestSdkWeightedOperations tests the weights of the operations.
func (suite *SimTestSuite) TestSdkWeightedOperations() {
	cdc := suite.app.AppCodec()
	appParams := make(simtypes.AppParams)

	weightesOps := simulation.SdkWeightedOperations(appParams, cdc, suite.app.AccountKeeper,
		suite.app.BankKeeper, suite.app.DistrKeeper, suite.app.StakingKeeper,
	)

	// setup 3 accounts
	s := rand.NewSource(1)
	r := rand.New(s)
	accs := suite.getTestingAccounts(r, 3)

	expected := []struct {
		weight     int
		opMsgRoute string
		opMsgName  string
	}{
		{simappparams.DefaultWeightMsgSetWithdrawAddress, sdkdistr.ModuleName, sdkdistr.TypeMsgSetWithdrawAddress},
		{simappparams.DefaultWeightMsgWithdrawDelegationReward, sdkdistr.ModuleName, sdkdistr.TypeMsgWithdrawDelegatorReward},
		{simappparams.DefaultWeightMsgWithdrawValidatorCommission, sdkdistr.ModuleName, sdkdistr.TypeMsgWithdrawValidatorCommission},
		{simappparams.DefaultWeightMsgFundCommunityPool, sdkdistr.ModuleName, sdkdistr.TypeMsgFundCommunityPool},
	}

	for i, w := range weightesOps {
		operationMsg, _, _ := w.Op()(r, suite.app.BaseApp, suite.ctx, accs, "")
		// the following checks are very much dependent from the ordering of the output given
		// by WeightedOperations. if the ordering in WeightedOperations changes some tests
		// will fail
		suite.Require().Equal(expected[i].weight, w.Weight(), "weight should be the same")
		suite.Require().Equal(expected[i].opMsgRoute, operationMsg.Route, "route should be the same")
		suite.Require().Equal(expected[i].opMsgName, operationMsg.Name, "operation Msg name should be the same")
	}
}

// TestSimulateSdkMsgSetWithdrawAddress tests the normal scenario of a valid message of type TypeMsgSetWithdrawAddress.
// Abonormal scenarios, where the message is created by an errors, are not tested here.
func (suite *SimTestSuite) TestSimulateSdkMsgSetWithdrawAddress() {
	// setup 3 accounts
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 3)

	// begin a new block
	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})

	// execute operation
	op := simulation.SimulateSdkMsgSetWithdrawAddress(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.DistrKeeper)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, suite.ctx, accounts, "")
	suite.Require().NoError(err)

	var msg sdkdistr.MsgSetWithdrawAddress
	sdkdistr.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	suite.Require().True(operationMsg.OK)
	suite.Require().Equal("cosmos1ghekyjucln7y67ntx7cf27m9dpuxxemn4c8g4r", msg.DelegatorAddress)
	suite.Require().Equal("cosmos1p8wcgrjr4pjju90xg6u9cgq55dxwq8j7u4x9a0", msg.WithdrawAddress)
	suite.Require().Equal(sdkdistr.TypeMsgSetWithdrawAddress, msg.Type())
	suite.Require().Equal(sdkdistr.ModuleName, msg.Route())
	suite.Require().Len(futureOperations, 0)
}

// TestSimulateSdkMsgWithdrawDelegatorReward tests the normal scenario of a valid message
// of type TypeMsgWithdrawDelegatorReward.
// Abonormal scenarios, where the message is created by an errors, are not tested here.
func (suite *SimTestSuite) TestSimulateSdkMsgWithdrawDelegatorReward() {
	// setup 3 accounts
	s := rand.NewSource(4)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 3)

	// setup accounts[0] as validator
	validator0 := suite.getTestingValidator0(accounts)

	// setup delegation
	delTokens := suite.app.StakingKeeper.TokensFromConsensusPower(suite.ctx, 2)
	validator0, issuedShares := validator0.AddTokensFromDel(delTokens)
	delegator := accounts[1]
	delegation := stakingtypes.NewDelegation(delegator.Address, validator0.GetOperator(), issuedShares, false)
	suite.app.StakingKeeper.SetDelegation(suite.ctx, delegation)
	suite.app.DistrKeeper.SetDelegatorStartingInfo(suite.ctx, validator0.GetOperator(), delegator.Address, distrtypes.NewDelegatorStartingInfo(2, sdk.OneDec(), 200))

	suite.setupValidatorRewards(validator0.GetOperator())

	// begin a new block
	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})

	// execute operation
	op := simulation.SimulateSdkMsgWithdrawDelegatorReward(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.DistrKeeper, suite.app.StakingKeeper)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, suite.ctx, accounts, "")
	suite.Require().NoError(err)

	var msg sdkdistr.MsgWithdrawDelegatorReward
	sdkdistr.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	suite.Require().True(operationMsg.OK)
	suite.Require().Equal("cosmosvaloper1l4s054098kk9hmr5753c6k3m2kw65h686d3mhr", msg.ValidatorAddress)
	suite.Require().Equal("cosmos1d6u7zhjwmsucs678d7qn95uqajd4ucl9jcjt26", msg.DelegatorAddress)
	suite.Require().Equal(sdkdistr.TypeMsgWithdrawDelegatorReward, msg.Type())
	suite.Require().Equal(sdkdistr.ModuleName, msg.Route())
	suite.Require().Len(futureOperations, 0)
}

// TestSimulateSdkMsgWithdrawValidatorCommission tests the normal scenario of a valid message
// of type TypeMsgWithdrawValidatorCommission.
// Abonormal scenarios, where the message is created by an errors, are not tested here.
func (suite *SimTestSuite) TestSimulateSdkMsgWithdrawValidatorCommission() {
	suite.testSimulateSdkMsgWithdrawValidatorCommission("atoken")
	suite.testSimulateSdkMsgWithdrawValidatorCommission("tokenxxx")
}

// all the checks in this function should not fail if we change the tokenName
func (suite *SimTestSuite) testSimulateSdkMsgWithdrawValidatorCommission(tokenName string) {
	// setup 3 accounts
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 3)

	// setup accounts[0] as validator
	validator0 := suite.getTestingValidator0(accounts)

	// set module account coins
	distrAcc := suite.app.DistrKeeper.GetDistributionAccount(suite.ctx)
	suite.Require().NoError(testutil.FundModuleAccount(suite.app.BankKeeper, suite.ctx, distrAcc.GetName(), sdk.NewCoins(
		sdk.NewCoin(tokenName, sdk.NewInt(10)),
		sdk.NewCoin("stake", sdk.NewInt(5)),
	)))
	suite.app.AccountKeeper.SetModuleAccount(suite.ctx, distrAcc)

	// set outstanding rewards
	valCommission := sdk.NewDecCoins(
		sdk.NewDecCoinFromDec(tokenName, sdk.NewDec(5).Quo(sdk.NewDec(2))),
		sdk.NewDecCoinFromDec("stake", sdk.NewDec(1).Quo(sdk.NewDec(1))),
	)

	suite.app.DistrKeeper.SetValidatorOutstandingRewards(suite.ctx, validator0.GetOperator(), types.ValidatorOutstandingRewards{Rewards: valCommission})
	suite.app.DistrKeeper.SetValidatorOutstandingRewards(suite.ctx, suite.genesisVals[0].GetOperator(), types.ValidatorOutstandingRewards{Rewards: valCommission})

	// setup validator accumulated commission
	suite.app.DistrKeeper.SetValidatorAccumulatedCommission(suite.ctx, validator0.GetOperator(), types.ValidatorAccumulatedCommission{Commission: valCommission})
	suite.app.DistrKeeper.SetValidatorAccumulatedCommission(suite.ctx, suite.genesisVals[0].GetOperator(), types.ValidatorAccumulatedCommission{Commission: valCommission})

	// begin a new block
	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})

	// execute operation
	op := simulation.SimulateSdkMsgWithdrawValidatorCommission(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.DistrKeeper, suite.app.StakingKeeper)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, suite.ctx, accounts, "")
	if !operationMsg.OK {
		suite.Require().Equal("could not find account", operationMsg.Comment)
	} else {
		suite.Require().NoError(err)

		var msg sdkdistr.MsgWithdrawValidatorCommission
		sdkdistr.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

		suite.Require().True(operationMsg.OK)
		suite.Require().Equal("cosmosvaloper1tnh2q55v8wyygtt9srz5safamzdengsn9dsd7z", msg.ValidatorAddress)
		suite.Require().Equal(sdkdistr.TypeMsgWithdrawValidatorCommission, msg.Type())
		suite.Require().Equal(sdkdistr.ModuleName, msg.Route())
		suite.Require().Len(futureOperations, 0)
	}
}

// TestSimulateSdkMsgFundCommunityPool tests the normal scenario of a valid message of type TypeMsgFundCommunityPool.
// Abonormal scenarios, where the message is created by an errors, are not tested here.
func (suite *SimTestSuite) TestSimulateSdkMsgFundCommunityPool() {
	// setup 3 accounts
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 3)

	// begin a new block
	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})

	// execute operation
	op := simulation.SimulateSdkMsgFundCommunityPool(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.DistrKeeper, suite.app.StakingKeeper)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, suite.ctx, accounts, "")
	suite.Require().NoError(err)

	var msg sdkdistr.MsgFundCommunityPool
	sdkdistr.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	suite.Require().True(operationMsg.OK)
	suite.Require().Equal("4896096stake", msg.Amount.String())
	suite.Require().Equal("cosmos1ghekyjucln7y67ntx7cf27m9dpuxxemn4c8g4r", msg.Depositor)
	suite.Require().Equal(sdkdistr.TypeMsgFundCommunityPool, msg.Type())
	suite.Require().Equal(sdkdistr.ModuleName, msg.Route())
	suite.Require().Len(futureOperations, 0)
}
