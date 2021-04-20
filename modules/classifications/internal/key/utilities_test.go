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

func Test_classificationIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	var tests []struct {
		name string
		args args
		want classificationID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := classificationIDFromInterface(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("classificationIDFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readClassificationID(t *testing.T) {
	type args struct {
		classificationIDString string
	}
	var tests []struct {
		name string
		args args
		want types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readClassificationID(tt.args.classificationIDString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readClassificationID() = %v, want %v", got, tt.want)
			}
		})
	}
}
