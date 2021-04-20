package transaction

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"reflect"
	"testing"
)

func TestRegisterCodec(t *testing.T) {
	type args struct {
		messagePrototype func() helpers.Message
	}
	var tests []struct {
		name string
		args args
		want *codec.Codec
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterCodec(tt.args.messagePrototype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterCodec() = %v, want %v", got, tt.want)
			}
		})
	}
}
