package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"reflect"
	"testing"
)

func TestCustomEncoder(t *testing.T) {
	type args struct {
		moduleList []helpers.Module
	}
	var tests []struct {
		name string
		args args
		want wasm.CustomEncoder
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CustomEncoder(tt.args.moduleList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
