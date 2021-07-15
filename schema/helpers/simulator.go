/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	//"github.com/cosmos/cosmos-sdk/x/simulation"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

type Simulator interface {
	RandomizedGenesisState(*module.SimulationState)
	WeightedOperations(simulationTypes.Params, *codec.LegacyAmino) []simulationTypes.WeightedOperation
	WeightedProposalContentList() []simulationTypes.WeightedProposalContent
	ParamChangeList(*rand.Rand) []simulation.ParamChange
}
