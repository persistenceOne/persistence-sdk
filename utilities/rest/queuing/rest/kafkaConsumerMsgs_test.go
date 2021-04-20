package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/utilities/rest/queuing"
	"testing"
)

func TestKafkaConsumerMessages(t *testing.T) {
	type args struct {
		cliCtx     context.CLIContext
		kafkaState queuing.KafkaState
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
