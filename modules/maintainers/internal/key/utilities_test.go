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
	var tests []struct {
		name string
		args args
		want types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadClassificationID(tt.args.assetID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadClassificationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadIdentityID(t *testing.T) {
	type args struct {
		assetID types.ID
	}
	var tests []struct {
		name string
		args args
		want types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadIdentityID(tt.args.assetID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadIdentityID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainerIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	var tests []struct {
		name string
		args args
		want maintainerID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maintainerIDFromInterface(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maintainerIDFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readMaintainerID(t *testing.T) {
	type args struct {
		maintainerIDString string
	}
	var tests []struct {
		name string
		args args
		want types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readMaintainerID(tt.args.maintainerIDString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readMaintainerID() = %v, want %v", got, tt.want)
			}
		})
	}
}
