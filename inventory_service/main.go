package main

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/MrRique15/go_rest_api/inventory_service/consumers"
	"github.com/MrRique15/go_rest_api/inventory_service/env"
	"github.com/MrRique15/go_rest_api/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/pkg/shared/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	order_consumer, err := kafka.StartKafkaConsumer(env.EnvKafkaHost())
	mongodb.ConnectMongoDB(env.EnvMongoURI())

	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v", err)
	}

	consumers.OrdersConsumer(order_consumer, "inventory", sarama.OffsetNewest)
}
