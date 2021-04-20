package simulator

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"reflect"
	"testing"
)

func Test_simulateMsg(t *testing.T) {
	var tests []struct {
		name string
		want simulation.Operation
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateMsg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simulateMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simulator_WeightedOperations(t *testing.T) {
	type args struct {
		appParams simulation.AppParams
		codec     *codec.Codec
	}
	var tests []struct {
		name string
		args args
		want simulation.WeightedOperations
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := simulator{}
			if got := si.WeightedOperations(tt.args.appParams, tt.args.codec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeightedOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
