package event_handlers

import (
	"fmt"
	"log"

	"github.com/MrRique15/go_rest_api/payment_service/payment_operations"
	"github.com/MrRique15/go_rest_api/payment_service/producers"
	"github.com/MrRique15/go_rest_api/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/pkg/utils"
)

type EventHandler func(event models.KafkaOrderEvent) error

var eventHandlers = map[string]EventHandler{
	kafka.Payment_KafkaEvents[0]: handleVerifyerifyPayment,
	kafka.Payment_KafkaEvents[1]: handleRollbackPayment,
}

func sendKafkaEvent(topic string, event models.KafkaOrderEvent) error {
	stringOrder, error := utils.OrderToJson(event)

	if error != nil {
		log.Println("error during order transformation for kafka event")
		return error
	}
	
	producers.KafkaProducer.SendKafkaEvent(topic, stringOrder)

	return nil
}

func handleVerifyerifyPayment(event models.KafkaOrderEvent) error {
	err := payment_operations.CheckPayment(event.Order)

	if err != nil {
		fmt.Println("Error checking payment for order: ", event.Order.ID)

		event.Event = kafka.OrdersSEC_KafkaEvents[5]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return err
	}

	payment_operations.UpdateOrderStatus(event.Order, "paid")

	event.Event = kafka.OrdersSEC_KafkaEvents[8]
	sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

	return nil
}

func handleRollbackPayment(event models.KafkaOrderEvent) error {
	err := payment_operations.RollbackPayment(event.Order)

	if err != nil {
		fmt.Println("Error rolling back payment for order: ", event.Order.ID)

		event.Event = kafka.OrdersSEC_KafkaEvents[6]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return err
	}

	payment_operations.UpdateOrderStatus(event.Order, "refounded")

	event.Event = kafka.OrdersSEC_KafkaEvents[7]
	sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

	return nil
}

func HandleEvent(event models.KafkaOrderEvent) {
	handler := eventHandlers[event.Event]
	err := handler(event)

	if err != nil {
		fmt.Println("Error handling event: ", err)
	}
}