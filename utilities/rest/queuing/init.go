package queuing

import (
	"github.com/cosmos/cosmos-sdk/client"
	"time"
)

func InitializeKafka(nodeList []string, cliContext client.Context) {
	KafkaState = *NewKafkaState(nodeList)
	if KafkaState.IsEnabled {
		go func() {
			for {
				kafkaConsumerMessages(cliContext)
				time.Sleep(sleepRoutine)
			}
		}()
	}
}
