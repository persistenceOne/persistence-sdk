package simulation

import (
	"errors"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	sdkslashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/persistenceOne/persistence-sdk/v2/simapp/helpers"
	simappparams "github.com/persistenceOne/persistence-sdk/v2/simapp/params"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/slashing/keeper"
	"github.com/persistenceOne/persistence-sdk/v2/x/lsnative/slashing/types"
	stakingkeeper "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/keeper"
)

// SdkWeightedOperations returns all the operations from the module with their respective weights
func SdkWeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper, sk types.StakingKeeper,
) simulation.WeightedOperations {
	var weightMsgUnjail int
	appParams.GetOrGenerate(cdc, OpWeightMsgUnjail, &weightMsgUnjail, nil,
		func(_ *rand.Rand) {
			weightMsgUnjail = simappparams.DefaultWeightMsgUnjail
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgUnjail,
			SimulateSdkMsgUnjail(ak, bk, k, sk.(stakingkeeper.Keeper)),
		),
	}
}

// SimulateSdkMsgUnjail generates a MsgUnjail with random values
func SimulateSdkMsgUnjail(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper, sk stakingkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		validator, ok := stakingkeeper.RandomValidator(r, sk, ctx)
		if !ok {
			return simtypes.NoOpMsg(sdkslashing.ModuleName, sdkslashing.TypeMsgUnjail, "validator is not ok"), nil, nil // skip
		}

		simAccount, found := simtypes.FindAccount(accs, sdk.AccAddress(validator.GetOperator()))
		if !found {
			return simtypes.NoOpMsg(sdkslashing.ModuleName, sdkslashing.TypeMsgUnjail, "unable to find account"), nil, nil // skip
		}

		if !validator.IsJailed() {
			// TODO: due to this condition this message is almost, if not always, skipped !
			return simtypes.NoOpMsg(sdkslashing.ModuleName, sdkslashing.TypeMsgUnjail, "validator is not jailed"), nil, nil
		}

		consAddr, err := validator.GetConsAddr()
		if err != nil {
			return simtypes.NoOpMsg(sdkslashing.ModuleName, sdkslashing.TypeMsgUnjail, "unable to get validator consensus key"), nil, err
		}
		info, found := k.GetValidatorSigningInfo(ctx, consAddr)
		if !found {
			return simtypes.NoOpMsg(sdkslashing.ModuleName, sdkslashing.TypeMsgUnjail, "unable to find validator signing info"), nil, nil // skip
		}

		selfDel := sk.Delegation(ctx, simAccount.Address, validator.GetOperator())
		if selfDel == nil {
			return simtypes.NoOpMsg(sdkslashing.ModuleName, sdkslashing.TypeMsgUnjail, "self delegation is nil"), nil, nil // skip
		}

		account := ak.GetAccount(ctx, sdk.AccAddress(validator.GetOperator()))
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(sdkslashing.ModuleName, sdkslashing.TypeMsgUnjail, "unable to generate fees"), nil, err
		}

		msg := sdkslashing.NewMsgUnjail(validator.GetOperator())

		txCfg := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenSignedMockTx(
			r,
			txCfg,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(sdkslashing.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		_, res, err := app.SimDeliver(txCfg.TxEncoder(), tx)

		// result should fail if:
		// - validator cannot be unjailed due to tombstone
		// - validator is still in jailed period
		// - self delegation too low
		if info.Tombstoned ||
			ctx.BlockHeader().Time.Before(info.JailedUntil) {
			if res != nil && err == nil {
				if info.Tombstoned {
					return simtypes.NewOperationMsg(msg, true, "", nil), nil, errors.New("validator should not have been unjailed if validator tombstoned")
				}
				if ctx.BlockHeader().Time.Before(info.JailedUntil) {
					return simtypes.NewOperationMsg(msg, true, "", nil), nil, errors.New("validator unjailed while validator still in jail period")
				}
			}
			// msg failed as expected
			return simtypes.NewOperationMsg(msg, false, "", nil), nil, nil
		}

		if err != nil {
			noop := simtypes.NoOpMsg(sdkslashing.ModuleName, msg.Type(), "unable to deliver tx")
			if res != nil {
				return noop, nil, errors.New(res.Log)
			}

			return noop, nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
