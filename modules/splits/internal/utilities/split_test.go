package utilities

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"reflect"
	"testing"
)

func TestAddSplits(t *testing.T) {
	type args struct {
		splits    helpers.Collection
		ownerID   types.ID
		ownableID types.ID
		value     sdkTypes.Dec
	}
	var tests []struct {
		name    string
		args    args
		want    helpers.Collection
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddSplits(tt.args.splits, tt.args.ownerID, tt.args.ownableID, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddSplits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSplits() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtractSplits(t *testing.T) {
	type args struct {
		splits    helpers.Collection
		ownerID   types.ID
		ownableID types.ID
		value     sdkTypes.Dec
	}
	var tests []struct {
		name    string
		args    args
		want    helpers.Collection
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubtractSplits(tt.args.splits, tt.args.ownerID, tt.args.ownableID, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubtractSplits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubtractSplits() got = %v, want %v", got, tt.want)
			}
		})
	}
}
