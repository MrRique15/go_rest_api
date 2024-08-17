package consumers

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func OrdersConsumer(consumer sarama.Consumer) {
	// create a consumer, conect to kafka topic and stay listening for messages
	for {
		// Choose the topic and partition
		topic := "order"
		partition := int32(0)

		// Start reading messages
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Fatalf("Error consuming partition: %v", err)
		}

		// Loop to listen to messages
		fmt.Println("Listening to messages...")
		for msg := range partitionConsumer.Messages() {
			fmt.Printf("Received: %s\n", string(msg.Value))
		}
	}
}