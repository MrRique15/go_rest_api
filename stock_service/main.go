package main

import (
	"log"

	"stock_service/configs"
	"stock_service/consumers"
)

func main() {

	order_consumer, err := configs.ConnectKafka()
	configs.ConnectMongoDB()

	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v", err)
	}

	consumers.OrdersConsumer(order_consumer)
}
