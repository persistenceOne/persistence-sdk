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

func TestReadClassificationID(t *testing.T) {
	type args struct {
		assetID types.ID
	}
	tests := []struct {
		name string
		args args
		want types.ID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadClassificationID(tt.args.assetID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadClassificationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assetIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want assetID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assetIDFromInterface(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assetIDFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readAssetID(t *testing.T) {
	type args struct {
		assetIDString string
	}
	tests := []struct {
		name string
		args args
		want types.ID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readAssetID(tt.args.assetIDString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readAssetID() = %v, want %v", got, tt.want)
			}
		})
	}
}
