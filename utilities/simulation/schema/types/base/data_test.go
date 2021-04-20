package base

import (
	"math/rand"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"reflect"
	"testing"
)

func TestGenerateRandomData(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want types.Data
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomData(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomData() = %v, want %v", got, tt.want)
			}
		})
	}
}
