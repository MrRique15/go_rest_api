package main

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/MrRique15/go_rest_api/saga_execution_controller/consumers"
	"github.com/MrRique15/go_rest_api/saga_execution_controller/env"
	"github.com/MrRique15/go_rest_api/saga_execution_controller/producers"
	"github.com/MrRique15/go_rest_api/pkg/shared/kafka"
)

func main() {
	orders_sec_consumer, err := kafka.StartKafkaConsumer(env.EnvKafkaHost())
	producers.StartKafkaProducer()

	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v", err)
	}

	consumers.OrdersSECConsumer(orders_sec_consumer, kafka.KafkaTopics["order_sec"], sarama.OffsetNewest)
}
