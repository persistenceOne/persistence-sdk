package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomProperties(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want types.Properties
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomProperties(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}
