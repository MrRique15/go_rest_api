package consumers

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/MrRique15/go_rest_api/inventory_service/inventory_operations"
	"github.com/MrRique15/go_rest_api/pkg/utils"
)

func OrdersConsumer(consumer sarama.Consumer, topic string, offset int64) {
	// create a consumer, conect to kafka topic and stay listening for messages
	for {
		partition := int32(0)

		partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
		if err != nil {
			log.Fatalf("Error consuming partition: %v", err)
		}

		// Loop to listen to messages
		fmt.Println("Listening to messages...")
		for msg := range partitionConsumer.Messages() {
			inventory_operations.ProcessOrder(utils.ByteToInterface(msg.Value))
		}
	}
}
