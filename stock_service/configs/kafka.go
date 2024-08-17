package configs

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func ConnectKafka() (sarama.Consumer, error) {
	// Configura o consumidor
    config := sarama.NewConfig()
    config.Consumer.Return.Errors = true

    // Conecta ao Kafka
    consumer, err := sarama.NewConsumer([]string{EnvKafkaHost()}, config)
    if err != nil {
        log.Fatalf("Erro ao criar consumidor: %v", err)
    }
    // defer consumer.Close()

	fmt.Println("Conectado ao Kafka")
	return consumer, nil
}