package conform

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"reflect"
	"testing"
)

func Test_auxiliaryKeeperMock_Help(t *testing.T) {
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		in0     types.Context
		request helpers.AuxiliaryRequest
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   helpers.AuxiliaryResponse
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryKeeper := auxiliaryKeeperMock{
				mapper: tt.fields.mapper,
			}
			if got := auxiliaryKeeper.Help(tt.args.in0, tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Help() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryKeeperMock_Initialize(t *testing.T) {
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		mapper helpers.Mapper
		in1    helpers.Parameters
		in2    []interface{}
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := auxiliaryKeeperMock{
				mapper: tt.fields.mapper,
			}
			if got := au.Initialize(tt.args.mapper, tt.args.in1, tt.args.in2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keeperPrototypeMock(t *testing.T) {
	var tests []struct {
		name string
		want helpers.AuxiliaryKeeper
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototypeMock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototypeMock() = %v, want %v", got, tt.want)
			}
		})
	}
}
