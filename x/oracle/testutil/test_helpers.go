package testutil

import (
	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutils "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	ValidatorInitPower    = int64(10000000000)
	ValidatorInitTokens   = sdk.TokensFromConsensusPower(ValidatorInitPower, sdk.DefaultPowerReduction)
	ValidatorInitCoins    = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, ValidatorInitTokens))
	ValidatorAmountBonded = sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)
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
	stakingKeeper *stakingkeeper.Keeper,
	ctx sdk.Context,
	num int,
) (
	[]sdk.AccAddress,
	[]sdk.ValAddress,
	error,
) {
	accAddresses := make([]sdk.AccAddress, num)
	valAddresses := make([]sdk.ValAddress, num)
	stakingMsgServer := stakingkeeper.NewMsgServerImpl(stakingKeeper)

	valPubKeys := simtestutils.CreateTestPubKeys(num)

	for i := 0; i < num; i++ {
		var (
			valPubKey = valPubKeys[i]
			pubKey    = secp256k1.GenPrivKey().PubKey()
			accAddr   = sdk.AccAddress(pubKey.Address())
			valAddr   = sdk.ValAddress(pubKey.Address())
		)

		// fund the validator account with initial coins
		if err := FundAccount(bankKeeper, ctx, accAddr, ValidatorInitCoins); err != nil {
			return nil, nil, err
		}

		// create validator in staking keeper with amount bonded
		createValidatorMsg, err := newTestMsgCreateValidator(valAddr, valPubKey, ValidatorAmountBonded)
		if err != nil {
			return nil, nil, err
		}

		if _, err := stakingMsgServer.CreateValidator(ctx, createValidatorMsg); err != nil {
			return nil, nil, err
		}

		accAddresses[i] = accAddr
		valAddresses[i] = valAddr
	}

	// ensure that validators are updated
	staking.EndBlocker(ctx, stakingKeeper)

	return accAddresses, valAddresses, nil
}

func newTestMsgCreateValidator(address sdk.ValAddress, pubKey cryptotypes.PubKey, amt math.Int) (*stakingtypes.MsgCreateValidator, error) {
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	msg, err := stakingtypes.NewMsgCreateValidator(
		address, pubKey, sdk.NewCoin(sdk.DefaultBondDenom, amt),
		stakingtypes.Description{}, commission, math.OneInt(),
	)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
