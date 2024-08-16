package main

import (
    "log"

	"payment_service/configs"
	"payment_service/consumers"
)

func main() {
    
	order_consumer, err := configs.ConnectKafka()

	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v", err)
	}

	consumers.OrdersConsumer(order_consumer)
}