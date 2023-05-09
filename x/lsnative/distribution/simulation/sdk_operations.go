package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdkdistr "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simappparams "github.com/persistenceOne/persistence-sdk/v2/simapp/params"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/distribution/keeper"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/distribution/types"
	stakingkeeper "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/keeper"
)

// SdkWeightedOperations returns all the operations from the module with their respective weights
func SdkWeightedOperations(appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk types.StakingKeeper) simulation.WeightedOperations {
	var weightMsgSetWithdrawAddress int
	appParams.GetOrGenerate(cdc, OpWeightMsgSetWithdrawAddress, &weightMsgSetWithdrawAddress, nil,
		func(_ *rand.Rand) {
			weightMsgSetWithdrawAddress = simappparams.DefaultWeightMsgSetWithdrawAddress
		},
	)

	var weightMsgWithdrawDelegationReward int
	appParams.GetOrGenerate(cdc, OpWeightMsgWithdrawDelegationReward, &weightMsgWithdrawDelegationReward, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawDelegationReward = simappparams.DefaultWeightMsgWithdrawDelegationReward
		},
	)

	var weightMsgWithdrawValidatorCommission int
	appParams.GetOrGenerate(cdc, OpWeightMsgWithdrawValidatorCommission, &weightMsgWithdrawValidatorCommission, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawValidatorCommission = simappparams.DefaultWeightMsgWithdrawValidatorCommission
		},
	)

	var weightMsgFundCommunityPool int
	appParams.GetOrGenerate(cdc, OpWeightMsgFundCommunityPool, &weightMsgFundCommunityPool, nil,
		func(_ *rand.Rand) {
			weightMsgFundCommunityPool = simappparams.DefaultWeightMsgFundCommunityPool
		},
	)

	stakeKeeper := sk.(stakingkeeper.Keeper)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgSetWithdrawAddress,
			SimulateSdkMsgSetWithdrawAddress(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgWithdrawDelegationReward,
			SimulateSdkMsgWithdrawDelegatorReward(ak, bk, k, stakeKeeper),
		),
		simulation.NewWeightedOperation(
			weightMsgWithdrawValidatorCommission,
			SimulateSdkMsgWithdrawValidatorCommission(ak, bk, k, stakeKeeper),
		),
		simulation.NewWeightedOperation(
			weightMsgFundCommunityPool,
			SimulateSdkMsgFundCommunityPool(ak, bk, k, stakeKeeper),
		),
	}
}

// SimulateSdkMsgSetWithdrawAddress generates a MsgSetWithdrawAddress with random values.
func SimulateSdkMsgSetWithdrawAddress(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		if !k.GetWithdrawAddrEnabled(ctx) {
			return simtypes.NoOpMsg(sdkdistr.ModuleName, sdkdistr.TypeMsgSetWithdrawAddress, "withdrawal is not enabled"), nil, nil
		}

		simAccount, _ := simtypes.RandomAcc(r, accs)
		simToAccount, _ := simtypes.RandomAcc(r, accs)

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := sdkdistr.NewMsgSetWithdrawAddress(simAccount.Address, simToAccount.Address)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      sdkdistr.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateSdkMsgWithdrawDelegatorReward generates a MsgWithdrawDelegatorReward with random values.
func SimulateSdkMsgWithdrawDelegatorReward(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk stakingkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		delegations := sk.GetAllDelegatorDelegations(ctx, simAccount.Address)
		if len(delegations) == 0 {
			return simtypes.NoOpMsg(sdkdistr.ModuleName, sdkdistr.TypeMsgWithdrawDelegatorReward, "number of delegators equal 0"), nil, nil
		}

		delegation := delegations[r.Intn(len(delegations))]

		validator := sk.Validator(ctx, delegation.GetValidatorAddr())
		if validator == nil {
			return simtypes.NoOpMsg(sdkdistr.ModuleName, sdkdistr.TypeMsgWithdrawDelegatorReward, "validator is nil"), nil, fmt.Errorf("validator %s not found", delegation.GetValidatorAddr())
		}

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := sdkdistr.NewMsgWithdrawDelegatorReward(simAccount.Address, validator.GetOperator())

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      sdkdistr.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateSdkMsgWithdrawValidatorCommission generates a MsgWithdrawValidatorCommission with random values.
func SimulateSdkMsgWithdrawValidatorCommission(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk stakingkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		validator, ok := stakingkeeper.RandomValidator(r, sk, ctx)
		if !ok {
			return simtypes.NoOpMsg(sdkdistr.ModuleName, sdkdistr.TypeMsgWithdrawValidatorCommission, "random validator is not ok"), nil, nil
		}

		commission := k.GetValidatorAccumulatedCommission(ctx, validator.GetOperator())
		if commission.Commission.IsZero() {
			return simtypes.NoOpMsg(sdkdistr.ModuleName, sdkdistr.TypeMsgWithdrawValidatorCommission, "validator commission is zero"), nil, nil
		}

		simAccount, found := simtypes.FindAccount(accs, sdk.AccAddress(validator.GetOperator()))
		if !found {
			return simtypes.NoOpMsg(sdkdistr.ModuleName, sdkdistr.TypeMsgWithdrawValidatorCommission, "could not find account"), nil, fmt.Errorf("validator %s not found", validator.GetOperator())
		}

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := sdkdistr.NewMsgWithdrawValidatorCommission(validator.GetOperator())

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      sdkdistr.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateSdkMsgFundCommunityPool simulates MsgFundCommunityPool execution where
// a random account sends a random amount of its funds to the community pool.
func SimulateSdkMsgFundCommunityPool(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk stakingkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		funder, _ := simtypes.RandomAcc(r, accs)

		account := ak.GetAccount(ctx, funder.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		fundAmount := simtypes.RandSubsetCoins(r, spendable)
		if fundAmount.Empty() {
			return simtypes.NoOpMsg(sdkdistr.ModuleName, sdkdistr.TypeMsgFundCommunityPool, "fund amount is empty"), nil, nil
		}

		var (
			fees sdk.Coins
			err  error
		)

		coins, hasNeg := spendable.SafeSub(fundAmount...)
		if !hasNeg {
			fees, err = simtypes.RandomFees(r, ctx, coins)
			if err != nil {
				return simtypes.NoOpMsg(sdkdistr.ModuleName, sdkdistr.TypeMsgFundCommunityPool, "unable to generate fees"), nil, err
			}
		}

		msg := sdkdistr.NewMsgFundCommunityPool(fundAmount, funder.Address)

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:           nil,
			Msg:           msg,
			MsgType:       msg.Type(),
			Context:       ctx,
			SimAccount:    funder,
			AccountKeeper: ak,
			ModuleName:    sdkdistr.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}
