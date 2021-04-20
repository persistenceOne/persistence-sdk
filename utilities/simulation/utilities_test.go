package simulation

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomAddresses(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want []sdkTypes.AccAddress
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomAddresses(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomBool(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomBool(tt.args.r); got != tt.want {
				t.Errorf("RandomBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
