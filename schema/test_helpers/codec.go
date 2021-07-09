package test_helpers

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

func NewTestInterfaceRegistry() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	RegisterInterfaces(registry)
	return registry
}

func RegisterInterfaces(registry types.InterfaceRegistry) {

	registry.RegisterInterface("helpers.Transactions", (*helpers.Transactions)(nil))
	registry.RegisterImplementations(
		(*helpers.Transactions)(nil),
		&Transactions{},
		&Transaction{})
}
