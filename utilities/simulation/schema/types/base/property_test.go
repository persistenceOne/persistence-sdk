package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomProperty(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want types.Property
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomProperty(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}
