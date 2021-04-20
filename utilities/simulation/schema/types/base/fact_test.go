package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomFact(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want types.Fact
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomFact(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomFact() = %v, want %v", got, tt.want)
			}
		})
	}
}
