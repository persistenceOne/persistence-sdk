package queuing

import (
	"github.com/Shopify/sarama"
	"reflect"
	"testing"
)

func TestKafkaAdmin(t *testing.T) {
	type args struct {
		kafkaPorts []string
	}
	var tests []struct {
		name string
		args args
		want sarama.ClusterAdmin
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KafkaAdmin(tt.args.kafkaPorts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KafkaAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopicsInit(t *testing.T) {
	type args struct {
		admin sarama.ClusterAdmin
		topic string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
