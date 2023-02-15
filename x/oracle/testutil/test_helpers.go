package testutil

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// FundAccount is a utility function that funds an account by minting and
// sending the coins to the address. This should be used for testing purposes
// only!
func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

// FundModuleAccount is a utility function that funds a module account by
// minting and sending the coins to the address. This should be used for testing
// purposes only!
func FundModuleAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, recipientMod, amounts)
}

// StakingAddValidators generates N validators and adds them to staking keeper as bonded validators.
func StakingAddValidators(
	bankKeeper bankkeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
	ctx sdk.Context,
	num int,
) (
	accAddresses []sdk.AccAddress,
	valAddresses []sdk.ValAddress,
	err error,
) {
	stakingHandler := staking.NewHandler(stakingKeeper)

	var (
		valPubKeys   = simapp.CreateTestPubKeys(num)
		initPower    = int64(10000000000)
		initTokens   = sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction)
		initCoins    = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, initTokens))
		amountBonded = sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)
	)

	for i := 0; i < num; i++ {
		var (
			valPubKey = valPubKeys[i]
			pubKey    = secp256k1.GenPrivKey().PubKey()
			accAddr   = sdk.AccAddress(pubKey.Address())
			valAddr   = sdk.ValAddress(pubKey.Address())
		)

		// fund the validator account with initial coins
		orPanic(FundAccount(bankKeeper, ctx, accAddr, initCoins))

		// create validator in staking keeper with amount bonded
		createValidatorMsg := newTestMsgCreateValidator(valAddr, valPubKey, amountBonded)
		_, err := stakingHandler(ctx, createValidatorMsg)
		orPanic(err)

		accAddresses = append(accAddresses, accAddr)
		valAddresses = append(valAddresses, valAddr)
	}

	// ensure that validators are updated
	staking.EndBlocker(ctx, stakingKeeper)
	return
}

func newTestMsgCreateValidator(address sdk.ValAddress, pubKey cryptotypes.PubKey, amt sdk.Int) *stakingtypes.MsgCreateValidator {
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	msg, err := stakingtypes.NewMsgCreateValidator(
		address, pubKey, sdk.NewCoin(sdk.DefaultBondDenom, amt),
		stakingtypes.Description{}, commission, sdk.OneInt(),
	)
	orPanic(err)

	return msg
}

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}
