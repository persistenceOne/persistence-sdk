package queuing

import (
	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/codec"
	"reflect"
	"testing"
)

func TestKafkaTopicConsumer(t *testing.T) {
	type args struct {
		topic     string
		consumers map[string]sarama.PartitionConsumer
		cdc       *codec.Codec
	}
	var tests []struct {
		name string
		args args
		want KafkaMsg
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KafkaTopicConsumer(tt.args.topic, tt.args.consumers, tt.args.cdc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KafkaTopicConsumer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConsumer(t *testing.T) {
	type args struct {
		kafkaPorts []string
	}
	var tests []struct {
		name string
		args args
		want sarama.Consumer
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConsumer(tt.args.kafkaPorts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConsumer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartitionConsumers(t *testing.T) {
	type args struct {
		consumer sarama.Consumer
		topic    string
	}
	var tests []struct {
		name string
		args args
		want sarama.PartitionConsumer
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartitionConsumers(tt.args.consumer, tt.args.topic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartitionConsumers() = %v, want %v", got, tt.want)
			}
		})
	}
}
