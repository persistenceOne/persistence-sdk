package v046

import "github.com/persistenceOne/persistence-sdk/v2/x/lsnative/staking/types"

// MigrateJSON accepts exported v0.43 x/stakinng genesis state and migrates it to
// v0.46 x/staking genesis state. The migration includes:
//
// - Add MinCommissionRate & ExemptionFactor params.
func MigrateJSON(oldState types.GenesisState) (types.GenesisState, error) {
	oldState.Params.MinCommissionRate = types.DefaultMinCommissionRate
	oldState.Params.ExemptionFactor = types.DefaultExemptionFactor

	return oldState, nil
}
