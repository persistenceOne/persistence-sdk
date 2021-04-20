package module

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"reflect"
	"testing"
)

func TestRegisterCodec(t *testing.T) {
	type args struct {
		keyPrototype      func() helpers.Key
		mappablePrototype func() helpers.Mappable
	}
	var tests []struct {
		name string
		args args
		want *codec.Codec
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterCodec(tt.args.keyPrototype, tt.args.mappablePrototype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterCodec() = %v, want %v", got, tt.want)
			}
		})
	}
}
