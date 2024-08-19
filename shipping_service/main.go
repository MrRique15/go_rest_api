package main

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/MrRique15/go_rest_api/shipping_service/consumers"
	"github.com/MrRique15/go_rest_api/shipping_service/env"
	"github.com/MrRique15/go_rest_api/shipping_service/producers"
	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/mongodb"
)

func main() {
	payment_consumer, err := kafka.StartKafkaConsumer(env.EnvKafkaHost())
	producers.StartKafkaProducer()
	mongodb.ConnectMongoDB(env.EnvMongoURI())

	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v", err)
	}

	consumers.ShippingConsumer(payment_consumer, kafka.KafkaTopics["shipping"], sarama.OffsetNewest)
}
