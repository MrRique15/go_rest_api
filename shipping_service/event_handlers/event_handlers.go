package event_handlers

import (
	"fmt"
	"log"

	"github.com/MrRique15/go_rest_api/shipping_service/shipping_operations"
	"github.com/MrRique15/go_rest_api/shipping_service/producers"
	"github.com/MrRique15/go_rest_api/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/pkg/utils"
)

type EventHandler func(event models.KafkaOrderEvent) error

var eventHandlers = map[string]EventHandler{
	kafka.Shipping_KafkaEvents[0]: handleStartShipping,
	kafka.Shipping_KafkaEvents[1]: handleRollbackShipping,
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

func handleStartShipping(event models.KafkaOrderEvent) error {
	err := shipping_operations.StartShipping(event.Order)

	if err != nil {
		fmt.Println("error starting shipping for order: ", event.Order.ID)

		event.Event = kafka.OrdersSEC_KafkaEvents[9]

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return err
	}

	updated_order, err := shipping_operations.UpdateOrderStatus(event.Order, "shipped")

	if err != nil {
		fmt.Println("error updating order status for order: ", event.Order.ID)

		event.Event = kafka.OrdersSEC_KafkaEvents[9]

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return err
	}

	kafka_shipping_event := models.KafkaOrderEvent{
		Event: kafka.OrdersSEC_KafkaEvents[10],
		Order: updated_order,
	}

	sendKafkaEvent(kafka.KafkaTopics["order_sec"], kafka_shipping_event)

	return nil
}

func handleRollbackShipping(event models.KafkaOrderEvent) error {
	err := shipping_operations.RollbackShipping(event.Order)

	if err != nil {
		fmt.Println("error rolling back shipping for order: ", event.Order.ID)

		event.Event = kafka.OrdersSEC_KafkaEvents[11]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return err
	}

	updated_order, err := shipping_operations.UpdateOrderStatus(event.Order, "on_hold")

	if err != nil {
		fmt.Println("error updating order status for order: ", event.Order.ID)

		event.Event = kafka.OrdersSEC_KafkaEvents[11]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return err
	}

	kafka_shipping_event := models.KafkaOrderEvent{
		Event: kafka.OrdersSEC_KafkaEvents[12],
		Order: updated_order,
	}

	sendKafkaEvent(kafka.KafkaTopics["order_sec"], kafka_shipping_event)

	return nil
}

func HandleEvent(event models.KafkaOrderEvent) {
	handler := eventHandlers[event.Event]
	err := handler(event)

	if err != nil {
		fmt.Println("Error handling event: ", err)
	}
}