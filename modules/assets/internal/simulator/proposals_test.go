package simulator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"
	"reflect"
	"testing"
)

func Test_simulateTextProposalContent(t *testing.T) {
	type args struct {
		r   *rand.Rand
		in1 sdk.Context
		in2 []simulation.Account
	}
	var tests []struct {
		name string
		args args
		want types.Content
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateTextProposalContent(tt.args.r, tt.args.in1, tt.args.in2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simulateTextProposalContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simulator_WeightedProposalContentList(t *testing.T) {
	var tests []struct {
		name string
		want []simulation.WeightedProposalContent
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := simulator{}
			if got := si.WeightedProposalContentList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeightedProposalContentList() = %v, want %v", got, tt.want)
			}
		})
	}
}
