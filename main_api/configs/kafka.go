package configs

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func ConnectKafka() (*sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer([]string{EnvKafkaHost()}, config)
    if err != nil {
        log.Fatal("Failed to start Kafka producer:", err)
    }
    // defer producer.Close()

	if err != nil {
		log.Fatal("Failed to connect to Kafka:", err)
		return nil, err
	}

	fmt.Println("Connected to Kafka")

	return &producer, nil
}