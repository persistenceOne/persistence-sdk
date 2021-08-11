/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"math/rand"
)

type Simulator interface {
	RandomizedGenesisState(*module.SimulationState)
	WeightedOperations(simulationTypes.AppParams, codec.JSONMarshaler) []simulationTypes.WeightedOperation
	WeightedProposalContentList() []simulationTypes.WeightedProposalContent
	ParamChangeList(*rand.Rand) []simulationTypes.ParamChange
}
