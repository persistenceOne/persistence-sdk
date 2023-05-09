package simulation_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdkslashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	simappparams "github.com/persistenceOne/persistence-sdk/v2/simapp/params"
	distrtypes "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/distribution/types"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/slashing/simulation"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/slashing/types"
	stakingtypes "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"
)

// TestSdkWeightedOperations tests the weights of the operations.
func TestSdkWeightedOperations(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)
	app, ctx, accs := createTestApp(t, false, r, 3)
	ctx.WithChainID("test-chain")

	cdc := app.AppCodec()
	appParams := make(simtypes.AppParams)

	expected := []struct {
		weight     int
		opMsgRoute string
		opMsgName  string
	}{{simappparams.DefaultWeightMsgUnjail, sdkslashing.ModuleName, sdkslashing.TypeMsgUnjail}}

	weightesOps := simulation.SdkWeightedOperations(appParams, cdc, app.AccountKeeper, app.BankKeeper, app.SlashingKeeper, app.StakingKeeper)
	for i, w := range weightesOps {
		operationMsg, _, _ := w.Op()(r, app.BaseApp, ctx, accs, ctx.ChainID())
		// the following checks are very much dependent from the ordering of the output given
		// by WeightedOperations. if the ordering in WeightedOperations changes some tests
		// will fail
		require.Equal(t, expected[i].weight, w.Weight(), "weight should be the same")
		require.Equal(t, expected[i].opMsgRoute, operationMsg.Route, "route should be the same")
		require.Equal(t, expected[i].opMsgName, operationMsg.Name, "operation Msg name should be the same")
	}
}

// TestSimulateSdkMsgUnjail tests the normal scenario of a valid message of type sdkslashing.MsgUnjail.
// Abonormal scenarios, where the message is created by an errors, are not tested here.
func TestSimulateSdkMsgUnjail(t *testing.T) {
	// setup 3 accounts
	s := rand.NewSource(5)
	r := rand.New(s)
	app, ctx, accounts := createTestApp(t, false, r, 3)
	blockTime := time.Now().UTC()
	ctx = ctx.WithBlockTime(blockTime)

	// remove genesis validator account
	accounts = accounts[1:]

	// setup accounts[0] as validator0
	validator0 := getTestingValidator0(t, app, ctx, accounts)

	// setup validator0 by consensus address
	app.StakingKeeper.SetValidatorByConsAddr(ctx, validator0)
	val0ConsAddress, err := validator0.GetConsAddr()
	require.NoError(t, err)
	info := types.NewValidatorSigningInfo(val0ConsAddress, int64(4), int64(3),
		time.Unix(2, 0), false, int64(10))
	app.SlashingKeeper.SetValidatorSigningInfo(ctx, val0ConsAddress, info)

	// put validator0 in jail
	app.StakingKeeper.Jail(ctx, val0ConsAddress)

	// setup self delegation
	delTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 2)
	validator0, issuedShares := validator0.AddTokensFromDel(delTokens)
	val0AccAddress, err := sdk.ValAddressFromBech32(validator0.OperatorAddress)
	require.NoError(t, err)
	selfDelegation := stakingtypes.NewDelegation(val0AccAddress.Bytes(), validator0.GetOperator(), issuedShares, false)
	app.StakingKeeper.SetDelegation(ctx, selfDelegation)
	app.DistrKeeper.SetDelegatorStartingInfo(ctx, validator0.GetOperator(), val0AccAddress.Bytes(), distrtypes.NewDelegatorStartingInfo(2, sdk.OneDec(), 200))

	// begin a new block
	app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: app.LastBlockHeight() + 1, AppHash: app.LastCommitID().Hash, Time: blockTime}})

	// execute operation
	op := simulation.SimulateSdkMsgUnjail(app.AccountKeeper, app.BankKeeper, app.SlashingKeeper, app.StakingKeeper)
	operationMsg, futureOperations, err := op(r, app.BaseApp, ctx, accounts, "")
	require.NoError(t, err)

	var msg sdkslashing.MsgUnjail
	sdkslashing.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	require.True(t, operationMsg.OK)
	require.Equal(t, sdkslashing.TypeMsgUnjail, msg.Type())
	require.Equal(t, "cosmosvaloper17s94pzwhsn4ah25tec27w70n65h5t2scgxzkv2", msg.ValidatorAddr)
	require.Len(t, futureOperations, 0)
}
