package types

import (
	"crypto/rand"
	"math/big"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkstaking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
)

var (
	DenomPersistence = Denom{
		BaseDenom:   PersistenceDenom,
		SymbolDenom: PersistenceSymbol,
		Exponent:    6,
	}

	DenomAtom = Denom{
		BaseDenom:   AtomDenom,
		SymbolDenom: AtomSymbol,
		Exponent:    6,
	}
)

// MockStakingKeeper implements the StakingKeeper interface.
type MockStakingKeeper struct {
	validators []MockValidator
}

func NewMockStakingKeeper(validators []MockValidator) MockStakingKeeper {
	return MockStakingKeeper{
		validators: validators,
	}
}

func (sk MockStakingKeeper) Validators() []MockValidator {
	return sk.validators
}

func (sk MockStakingKeeper) Validator(_ sdk.Context, address sdk.ValAddress) sdkstaking.ValidatorI {
	for _, validator := range sk.validators {
		if validator.GetOperator().Equals(address) {
			return validator
		}
	}

	return nil
}

func (MockStakingKeeper) TotalBondedTokens(sdk.Context) sdk.Int {
	return sdk.ZeroInt()
}

func (MockStakingKeeper) GetBondedValidatorsByPower(sdk.Context) []sdkstaking.Validator {
	return nil
}

func (MockStakingKeeper) ValidatorsPowerStoreIterator(sdk.Context) sdk.Iterator {
	return sdk.KVStoreReversePrefixIterator(nil, nil)
}

func (sk MockStakingKeeper) GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64) {
	return sk.Validator(ctx, operator).GetConsensusPower(sdk.DefaultPowerReduction)
}

func (MockStakingKeeper) MaxValidators(sdk.Context) uint32 {
	return 100
}

func (MockStakingKeeper) PowerReduction(sdk.Context) (res sdk.Int) {
	return sdk.DefaultPowerReduction
}

func (MockStakingKeeper) Slash(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec) sdk.Int {
	return sdk.ZeroInt()
}

func (MockStakingKeeper) Jail(sdk.Context, sdk.ConsAddress) {}

// MockValidator implements the ValidatorI interface.
type MockValidator struct {
	power    int64
	operator sdk.ValAddress
}

func NewMockValidator(valAddr sdk.ValAddress, power int64) MockValidator {
	return MockValidator{
		power:    power,
		operator: valAddr,
	}
}

func (MockValidator) IsJailed() bool {
	return false
}

func (MockValidator) GetMoniker() string {
	return ""
}

func (MockValidator) GetStatus() sdkstaking.BondStatus {
	return sdkstaking.Bonded
}

func (MockValidator) IsBonded() bool {
	return true
}

func (MockValidator) IsUnbonded() bool {
	return false
}

func (MockValidator) IsUnbonding() bool {
	return false
}

func (v MockValidator) GetOperator() sdk.ValAddress {
	return v.operator
}

func (MockValidator) ConsPubKey() (cryptotypes.PubKey, error) {
	return nil, nil
}

func (MockValidator) TmConsPublicKey() (tmprotocrypto.PublicKey, error) {
	return tmprotocrypto.PublicKey{}, nil
}

func (MockValidator) GetConsAddr() (sdk.ConsAddress, error) {
	return nil, nil
}

func (v MockValidator) GetTokens() sdk.Int {
	return sdk.TokensFromConsensusPower(v.power, sdk.DefaultPowerReduction)
}

func (v MockValidator) GetBondedTokens() sdk.Int {
	return sdk.TokensFromConsensusPower(v.power, sdk.DefaultPowerReduction)
}

func (v MockValidator) GetConsensusPower(sdk.Int) int64 {
	return v.power
}

func (v *MockValidator) SetConsensusPower(power int64) {
	v.power = power
}

func (MockValidator) GetCommission() sdk.Dec {
	return sdk.ZeroDec()
}

func (MockValidator) GetMinSelfDelegation() sdk.Int {
	return sdk.OneInt()
}

func (v MockValidator) GetDelegatorShares() sdk.Dec {
	return sdk.NewDec(v.power)
}

func (MockValidator) TokensFromShares(sdk.Dec) sdk.Dec {
	return sdk.ZeroDec()
}

func (MockValidator) TokensFromSharesTruncated(sdk.Dec) sdk.Dec {
	return sdk.ZeroDec()
}

func (MockValidator) TokensFromSharesRoundUp(sdk.Dec) sdk.Dec {
	return sdk.ZeroDec()
}

func (MockValidator) SharesFromTokens(sdk.Int) (sdk.Dec, error) {
	return sdk.ZeroDec(), nil
}

func (MockValidator) SharesFromTokensTruncated(sdk.Int) (sdk.Dec, error) {
	return sdk.ZeroDec(), nil
}

// generateRandomTestCase
func generateRandomTestCase() ([]sdk.ValAddress, MockStakingKeeper) {
	var (
		valValAddrs    []sdk.ValAddress
		mockValidators []MockValidator
	)

	randNum, _ := rand.Int(rand.Reader, big.NewInt(10000))
	numInputs := 10 + int(randNum.Int64()%100)

	for i := 0; i < numInputs; i++ {
		pubKey := secp256k1.GenPrivKey().PubKey()
		valValAddr := sdk.ValAddress(pubKey.Address())
		valValAddrs = append(valValAddrs, valValAddr)

		randomPower, _ := rand.Int(rand.Reader, big.NewInt(10000))
		power := randomPower.Int64()%1000 + 1

		mockValidator := NewMockValidator(valValAddr, power)
		mockValidators = append(mockValidators, mockValidator)
	}

	return valValAddrs, NewMockStakingKeeper(mockValidators)
}

// generateRandomValAddr returns N random validator addresses.
func generateRandomValAddr(quantity int) []sdk.ValAddress {
	var validatorAddrs []sdk.ValAddress

	for i := 0; i < quantity; i++ {
		pubKey := secp256k1.GenPrivKey().PubKey()
		valAddr := sdk.ValAddress(pubKey.Address())
		validatorAddrs = append(validatorAddrs, valAddr)
	}

	return validatorAddrs
}

// stringWithCharset generates a new string with the size of "length" param
// repeating every character of charset, if charset is empty uses "abcd"
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)

	if len(charset) == 0 {
		charset = "abcd"
	}

	for i := 0; i < length; i++ {
		for j := 0; j < len(charset); j++ {
			b[i] = charset[j]
			i++

			if len(b) == length {
				return string(b)
			}
		}
	}

	return string(b)
}
