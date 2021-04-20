package queuing

import (
	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/codec"
	"reflect"
	"testing"
)

func TestKafkaProducerDeliverMessage(t *testing.T) {
	type args struct {
		msg      KafkaMsg
		topic    string
		producer sarama.SyncProducer
		cdc      *codec.Codec
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := KafkaProducerDeliverMessage(tt.args.msg, tt.args.topic, tt.args.producer, tt.args.cdc); (err != nil) != tt.wantErr {
				t.Errorf("KafkaProducerDeliverMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewProducer(t *testing.T) {
	type args struct {
		kafkaPorts []string
	}
	var tests []struct {
		name string
		args args
		want sarama.SyncProducer
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProducer(tt.args.kafkaPorts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendToKafka(t *testing.T) {
	type args struct {
		msg        KafkaMsg
		kafkaState KafkaState
		cdc        *codec.Codec
	}
	var tests []struct {
		name string
		args args
		want []byte
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SendToKafka(tt.args.msg, tt.args.kafkaState, tt.args.cdc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendToKafka() = %v, want %v", got, tt.want)
			}
		})
	}
}
