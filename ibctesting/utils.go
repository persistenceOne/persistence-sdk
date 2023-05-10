package ibctesting

import (
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	sdkstaking "github.com/cosmos/cosmos-sdk/x/staking/types"
	stakingtypes "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"
)

// ApplyValSetChanges takes in tmtypes.ValidatorSet and []abci.ValidatorUpdate and will return a new tmtypes.ValidatorSet which has the
// provided validator updates applied to the provided validator set.
func ApplyValSetChanges(t *testing.T, valSet *tmtypes.ValidatorSet, valUpdates []abci.ValidatorUpdate) *tmtypes.ValidatorSet {
	updates, err := tmtypes.PB2TM.ValidatorUpdates(valUpdates)
	require.NoError(t, err)

	// must copy since validator set will mutate with UpdateWithChangeSet
	newVals := valSet.Copy()
	err = newVals.UpdateWithChangeSet(updates)
	require.NoError(t, err)

	return newVals
}

// SdkValidatorsToValidators converts sdk types into our custom staking types.
func SdkValidatorsToValidators(sdkVals sdkstaking.Validators) stakingtypes.Validators {
	vals := make([]stakingtypes.Validator, 0, len(sdkVals))

	for _, val := range sdkVals {
		vals = append(vals, stakingtypes.Validator(stakingtypes.Validator{
			OperatorAddress: val.OperatorAddress,
			ConsensusPubkey: val.ConsensusPubkey,
			Jailed:          val.Jailed,
			Status:          val.Status,
			Tokens:          val.Tokens,
			DelegatorShares: val.DelegatorShares,
			Description:     stakingtypes.Description(val.Description),
			UnbondingHeight: val.UnbondingHeight,
			UnbondingTime:   val.UnbondingTime,
			Commission: stakingtypes.Commission{
				CommissionRates: stakingtypes.CommissionRates(val.Commission.CommissionRates),
				UpdateTime:      val.Commission.UpdateTime,
			},
			MinSelfDelegation: val.MinSelfDelegation,
		}))
	}

	return vals
}
