package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"testing"
)

func TestRegisterXPRTConcrete(t *testing.T) {
	type args struct {
		codec      *codec.Codec
		moduleName string
		o          interface{}
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
