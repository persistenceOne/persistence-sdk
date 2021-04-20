package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomID(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomID(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomIDWithDec(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomIDWithDec(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomIDWithDec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomIDWithInt64(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	var tests []struct {
		name string
		args args
		want types.ID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomIDWithInt64(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomIDWithInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
