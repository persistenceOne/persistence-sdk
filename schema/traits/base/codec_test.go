package base

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"testing"
)

func TestRegisterCodec(t *testing.T) {
	type args struct {
		codec *codec.Codec
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
