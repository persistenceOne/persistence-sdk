package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"reflect"
	"testing"
)

func TestFromID(t *testing.T) {
	type args struct {
		id types.ID
	}
	var tests []struct {
		name string
		args args
		want helpers.Key
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	var tests []struct {
		name string
		args args
		want identityID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := identityIDFromInterface(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("identityIDFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readIdentityID(t *testing.T) {
	type args struct {
		identityIDString string
	}
	var tests []struct {
		name string
		args args
		want types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readIdentityID(tt.args.identityIDString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readIdentityID() = %v, want %v", got, tt.want)
			}
		})
	}
}
