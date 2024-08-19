package consumers

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/MrRique15/go_rest_api/shipping_service/event_handlers"
	"github.com/MrRique15/go_rest_api/pkg/utils"
)

func ShippingConsumer(consumer sarama.Consumer, topic string, offset int64) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, offset)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}

	fmt.Println("Shipping Service Listening to messages in topic: ", topic)
	for msg := range partitionConsumer.Messages() {
		fmt.Println("Received message: ", string(msg.Value))

		event := utils.ByteToInterface(msg.Value)
		if err != nil {
			log.Fatalf("Error unmarshalling message: %v", err)
		}

		event_handlers.HandleEvent(event)
	}

	select {}
}
