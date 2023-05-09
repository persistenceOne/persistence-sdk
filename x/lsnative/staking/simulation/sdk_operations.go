package simulation

import (
	"fmt"
	"math/rand"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	sdkstaking "github.com/cosmos/cosmos-sdk/x/staking/types"
	simappparams "github.com/persistenceOne/persistence-sdk/v2/simapp/params"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/keeper"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"
)

// SdkWeightedOperations returns all the operations from the module with their respective weights
func SdkWeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgCreateValidator             int
		weightMsgEditValidator               int
		weightMsgDelegate                    int
		weightMsgUndelegate                  int
		weightMsgBeginRedelegate             int
		weightMsgCancelUnbondingDelegation   int
		weightMsgTokenizeShares              int
		weightMsgRedeemTokensforShares       int
		weightMsgTransferTokenizeShareRecord int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateValidator, &weightMsgCreateValidator, nil,
		func(_ *rand.Rand) {
			weightMsgCreateValidator = simappparams.DefaultWeightMsgCreateValidator
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgEditValidator, &weightMsgEditValidator, nil,
		func(_ *rand.Rand) {
			weightMsgEditValidator = simappparams.DefaultWeightMsgEditValidator
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgDelegate, &weightMsgDelegate, nil,
		func(_ *rand.Rand) {
			weightMsgDelegate = simappparams.DefaultWeightMsgDelegate
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUndelegate, &weightMsgUndelegate, nil,
		func(_ *rand.Rand) {
			weightMsgUndelegate = simappparams.DefaultWeightMsgUndelegate
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBeginRedelegate, &weightMsgBeginRedelegate, nil,
		func(_ *rand.Rand) {
			weightMsgBeginRedelegate = simappparams.DefaultWeightMsgBeginRedelegate
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCancelUnbondingDelegation, &weightMsgCancelUnbondingDelegation, nil,
		func(_ *rand.Rand) {
			weightMsgCancelUnbondingDelegation = simappparams.DefaultWeightMsgCancelUnbondingDelegation
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgTokenizeShares, &weightMsgTokenizeShares, nil,
		func(_ *rand.Rand) {
			weightMsgTokenizeShares = simappparams.DefaultWeightMsgTokenizeShares
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRedeemTokensforShares, &weightMsgRedeemTokensforShares, nil,
		func(_ *rand.Rand) {
			weightMsgRedeemTokensforShares = simappparams.DefaultWeightMsgRedeemTokensforShares
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgTransferTokenizeShareRecord, &weightMsgTransferTokenizeShareRecord, nil,
		func(_ *rand.Rand) {
			weightMsgTransferTokenizeShareRecord = simappparams.DefaultWeightMsgTransferTokenizeShareRecord
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateValidator,
			SimulateSdkMsgCreateValidator(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgEditValidator,
			SimulateSdkMsgEditValidator(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgDelegate,
			SimulateSdkMsgDelegate(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUndelegate,
			SimulateSdkMsgUndelegate(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgBeginRedelegate,
			SimulateSdkMsgBeginRedelegate(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgCancelUnbondingDelegation,
			SimulateSdkMsgCancelUnbondingDelegate(ak, bk, k),
		),
	}
}

// SimulateSdkMsgCreateValidator generates a MsgCreateValidator with random values
func SimulateSdkMsgCreateValidator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		address := sdk.ValAddress(simAccount.Address)

		// ensure the validator doesn't exist already
		_, found := k.GetLiquidValidator(ctx, address)
		if found {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCreateValidator, "validator already exists"), nil, nil
		}

		denom := k.GetParams(ctx).BondDenom

		balance := bk.GetBalance(ctx, simAccount.Address, denom).Amount
		if !balance.IsPositive() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCreateValidator, "balance is negative"), nil, nil
		}

		amount, err := simtypes.RandPositiveInt(r, balance)
		if err != nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCreateValidator, "unable to generate positive amount"), nil, err
		}

		selfDelegation := sdk.NewCoin(denom, amount)

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		var fees sdk.Coins

		coins, hasNeg := spendable.SafeSub(selfDelegation)
		if !hasNeg {
			fees, err = simtypes.RandomFees(r, ctx, coins)
			if err != nil {
				return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCreateValidator, "unable to generate fees"), nil, err
			}
		}

		description := sdkstaking.NewDescription(
			simtypes.RandStringOfLength(r, 10),
			simtypes.RandStringOfLength(r, 10),
			simtypes.RandStringOfLength(r, 10),
			simtypes.RandStringOfLength(r, 10),
			simtypes.RandStringOfLength(r, 10),
		)

		maxCommission := sdk.NewDecWithPrec(int64(simtypes.RandIntBetween(r, 0, 100)), 2)
		commission := sdkstaking.NewCommissionRates(
			simtypes.RandomDecAmount(r, maxCommission),
			maxCommission,
			simtypes.RandomDecAmount(r, maxCommission),
		)

		msg, err := sdkstaking.NewMsgCreateValidator(address, simAccount.ConsKey.PubKey(), selfDelegation, description, commission, sdk.OneInt())
		if err != nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, msg.Type(), "unable to create CreateValidator message"), nil, err
		}

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:           nil,
			Msg:           msg,
			MsgType:       msg.Type(),
			Context:       ctx,
			SimAccount:    simAccount,
			AccountKeeper: ak,
			ModuleName:    sdkstaking.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}

// SimulateSdkMsgEditValidator generates a MsgEditValidator with random values
func SimulateSdkMsgEditValidator(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		if len(k.GetAllValidators(ctx)) == 0 {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgEditValidator, "number of validators equal zero"), nil, nil
		}

		val, ok := keeper.RandomValidator(r, k, ctx)
		if !ok {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgEditValidator, "unable to pick a validator"), nil, nil
		}

		address := val.GetOperator()

		newCommissionRate := simtypes.RandomDecAmount(r, val.Commission.MaxRate)

		if err := val.Commission.ValidateNewRate(newCommissionRate, ctx.BlockHeader().Time); err != nil {
			// skip as the commission is invalid
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgEditValidator, "invalid commission rate"), nil, nil
		}

		simAccount, found := simtypes.FindAccount(accs, sdk.AccAddress(val.GetOperator()))
		if !found {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgEditValidator, "unable to find account"), nil, fmt.Errorf("validator %s not found", val.GetOperator())
		}

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		description := sdkstaking.NewDescription(
			simtypes.RandStringOfLength(r, 10),
			simtypes.RandStringOfLength(r, 10),
			simtypes.RandStringOfLength(r, 10),
			simtypes.RandStringOfLength(r, 10),
			simtypes.RandStringOfLength(r, 10),
		)

		minSelfDelegation := math.OneInt()
		msg := sdkstaking.NewMsgEditValidator(address, description, &newCommissionRate, &minSelfDelegation)

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
			ModuleName:      sdkstaking.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateSdkMsgDelegate generates a MsgDelegate with random values
func SimulateSdkMsgDelegate(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		denom := k.GetParams(ctx).BondDenom

		if len(k.GetAllValidators(ctx)) == 0 {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgDelegate, "number of validators equal zero"), nil, nil
		}

		simAccount, _ := simtypes.RandomAcc(r, accs)
		val, ok := keeper.RandomValidator(r, k, ctx)
		if !ok {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgDelegate, "unable to pick a validator"), nil, nil
		}

		if val.InvalidExRate() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgDelegate, "validator's invalid echange rate"), nil, nil
		}

		amount := bk.GetBalance(ctx, simAccount.Address, denom).Amount
		if !amount.IsPositive() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgDelegate, "balance is negative"), nil, nil
		}

		amount, err := simtypes.RandPositiveInt(r, amount)
		if err != nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgDelegate, "unable to generate positive amount"), nil, err
		}

		bondAmt := sdk.NewCoin(denom, amount)

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		var fees sdk.Coins

		coins, hasNeg := spendable.SafeSub(bondAmt)
		if !hasNeg {
			fees, err = simtypes.RandomFees(r, ctx, coins)
			if err != nil {
				return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgDelegate, "unable to generate fees"), nil, err
			}
		}

		msg := sdkstaking.NewMsgDelegate(simAccount.Address, val.GetOperator(), bondAmt)

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:           nil,
			Msg:           msg,
			MsgType:       msg.Type(),
			Context:       ctx,
			SimAccount:    simAccount,
			AccountKeeper: ak,
			ModuleName:    sdkstaking.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}

// SimulateSdkMsgUndelegate generates a MsgUndelegate with random values
func SimulateSdkMsgUndelegate(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// get random validator
		validator, ok := keeper.RandomValidator(r, k, ctx)
		if !ok {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgUndelegate, "validator is not ok"), nil, nil
		}

		valAddr := validator.GetOperator()
		delegations := k.GetValidatorDelegations(ctx, validator.GetOperator())
		if delegations == nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgUndelegate, "keeper does have any delegation entries"), nil, nil
		}

		// get random delegator from validator
		delegation := delegations[r.Intn(len(delegations))]
		delAddr := delegation.GetDelegatorAddr()

		if k.HasMaxUnbondingDelegationEntries(ctx, delAddr, valAddr) {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgUndelegate, "keeper does have a max unbonding delegation entries"), nil, nil
		}

		totalBond := validator.TokensFromShares(delegation.GetShares()).TruncateInt()
		if !totalBond.IsPositive() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgUndelegate, "total bond is negative"), nil, nil
		}

		unbondAmt, err := simtypes.RandPositiveInt(r, totalBond)
		if err != nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgUndelegate, "invalid unbond amount"), nil, err
		}

		if unbondAmt.IsZero() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgUndelegate, "unbond amount is zero"), nil, nil
		}

		msg := sdkstaking.NewMsgUndelegate(
			delAddr, valAddr, sdk.NewCoin(k.BondDenom(ctx), unbondAmt),
		)

		// need to retrieve the simulation account associated with delegation to retrieve PrivKey
		var simAccount simtypes.Account

		for _, simAcc := range accs {
			if simAcc.Address.Equals(delAddr) {
				simAccount = simAcc
				break
			}
		}
		// if simaccount.PrivKey == nil, delegation address does not exist in accs. Return error
		if simAccount.PrivKey == nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, msg.Type(), "account private key is nil"), nil, fmt.Errorf("delegation addr: %s does not exist in simulation accounts", delAddr)
		}

		account := ak.GetAccount(ctx, delAddr)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

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
			ModuleName:      sdkstaking.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateSdkMsgCancelUnbondingDelegate generates a MsgCancelUnbondingDelegate with random values
func SimulateSdkMsgCancelUnbondingDelegate(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		if len(k.GetAllValidators(ctx)) == 0 {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgDelegate, "number of validators equal zero"), nil, nil
		}
		// get random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		// get random validator
		validator, ok := keeper.RandomValidator(r, k, ctx)
		if !ok {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCancelUnbondingDelegation, "validator is not ok"), nil, nil
		}

		if validator.IsJailed() || validator.InvalidExRate() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCancelUnbondingDelegation, "validator is jailed"), nil, nil
		}

		valAddr := validator.GetOperator()
		unbondingDelegation, found := k.GetUnbondingDelegation(ctx, simAccount.Address, valAddr)
		if !found {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCancelUnbondingDelegation, "account does have any unbonding delegation"), nil, nil
		}

		// get random unbonding delegation entry at block height
		unbondingDelegationEntry := unbondingDelegation.Entries[r.Intn(len(unbondingDelegation.Entries))]

		if unbondingDelegationEntry.CompletionTime.Before(ctx.BlockTime()) {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCancelUnbondingDelegation, "unbonding delegation is already processed"), nil, nil
		}

		if !unbondingDelegationEntry.Balance.IsPositive() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCancelUnbondingDelegation, "delegator receiving balance is negative"), nil, nil
		}

		cancelBondAmt := simtypes.RandomAmount(r, unbondingDelegationEntry.Balance)

		if cancelBondAmt.IsZero() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgCancelUnbondingDelegation, "cancelBondAmt amount is zero"), nil, nil
		}

		msg := sdkstaking.NewMsgCancelUnbondingDelegation(
			simAccount.Address, valAddr, unbondingDelegationEntry.CreationHeight, sdk.NewCoin(k.BondDenom(ctx), cancelBondAmt),
		)

		spendable := bk.SpendableCoins(ctx, simAccount.Address)

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
			ModuleName:      sdkstaking.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateSdkMsgBeginRedelegate generates a MsgBeginRedelegate with random values
func SimulateSdkMsgBeginRedelegate(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// get random source validator
		srcVal, ok := keeper.RandomValidator(r, k, ctx)
		if !ok {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "unable to pick validator"), nil, nil
		}

		srcAddr := srcVal.GetOperator()
		delegations := k.GetValidatorDelegations(ctx, srcAddr)
		if delegations == nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "keeper does have any delegation entries"), nil, nil
		}

		// get random delegator from src validator
		delegation := delegations[r.Intn(len(delegations))]
		delAddr := delegation.GetDelegatorAddr()

		if k.HasReceivingRedelegation(ctx, delAddr, srcAddr) {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "receveing redelegation is not allowed"), nil, nil // skip
		}

		// get random destination validator
		destVal, ok := keeper.RandomValidator(r, k, ctx)
		if !ok {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "unable to pick validator"), nil, nil
		}

		destAddr := destVal.GetOperator()
		if srcAddr.Equals(destAddr) || destVal.InvalidExRate() || k.HasMaxRedelegationEntries(ctx, delAddr, srcAddr, destAddr) {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "checks failed"), nil, nil
		}

		totalBond := srcVal.TokensFromShares(delegation.GetShares()).TruncateInt()
		if !totalBond.IsPositive() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "total bond is negative"), nil, nil
		}

		redAmt, err := simtypes.RandPositiveInt(r, totalBond)
		if err != nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "unable to generate positive amount"), nil, err
		}

		if redAmt.IsZero() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "amount is zero"), nil, nil
		}

		// check if the shares truncate to zero
		shares, err := srcVal.SharesFromTokens(redAmt)
		if err != nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "invalid shares"), nil, err
		}

		if srcVal.TokensFromShares(shares).TruncateInt().IsZero() {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "shares truncate to zero"), nil, nil // skip
		}

		// need to retrieve the simulation account associated with delegation to retrieve PrivKey
		var simAccount simtypes.Account

		for _, simAcc := range accs {
			if simAcc.Address.Equals(delAddr) {
				simAccount = simAcc
				break
			}
		}

		// if simaccount.PrivKey == nil, delegation address does not exist in accs. Return error
		if simAccount.PrivKey == nil {
			return simtypes.NoOpMsg(sdkstaking.ModuleName, sdkstaking.TypeMsgBeginRedelegate, "account private key is nil"), nil, fmt.Errorf("delegation addr: %s does not exist in simulation accounts", delAddr)
		}

		account := ak.GetAccount(ctx, delAddr)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := sdkstaking.NewMsgBeginRedelegate(
			delAddr, srcAddr, destAddr,
			sdk.NewCoin(k.BondDenom(ctx), redAmt),
		)

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
			ModuleName:      sdkstaking.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
