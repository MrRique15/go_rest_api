package consumers

import (
	"fmt"
	"log"
	"stock_service/order_operations"
	"stock_service/utils"

	"github.com/IBM/sarama"
)

func OrdersConsumer(consumer sarama.Consumer) {
	// create a consumer, conect to kafka topic and stay listening for messages
	for {
		// Choose the topic and partition
		topic := "order"
		partition := int32(0)

		// Start reading messages
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Error consuming partition: %v", err)
		}

		// Loop to listen to messages
		fmt.Println("Listening to messages...")
		for msg := range partitionConsumer.Messages() {
			order_operations.ProcessOrder(utils.ByteToInterface(msg.Value))
		}
	}
}
