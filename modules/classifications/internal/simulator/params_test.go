package simulator

import (
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"
	"reflect"
	"testing"
)

func Test_simulator_ParamChangeList(t *testing.T) {
	type args struct {
		in0 *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want []simulation.ParamChange
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := simulator{}
			if got := si.ParamChangeList(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParamChangeList() = %v, want %v", got, tt.want)
			}
		})
	}
}
