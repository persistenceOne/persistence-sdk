package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"reflect"
	"testing"
)

func TestGetOwnableTotalSplitsValue(t *testing.T) {
	type args struct {
		collection helpers.Collection
		ownableID  types.ID
	}
	var tests []struct {
		name string
		args args
		want sdkTypes.Dec
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOwnableTotalSplitsValue(tt.args.collection, tt.args.ownableID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwnableTotalSplitsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
