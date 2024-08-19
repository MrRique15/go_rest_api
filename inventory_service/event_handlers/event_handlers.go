package event_handlers

import (
	"errors"
	"fmt"
	"log"

	"github.com/MrRique15/go_rest_api/inventory_service/inventory_operations"
	"github.com/MrRique15/go_rest_api/inventory_service/producers"
	"github.com/MrRique15/go_rest_api/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/pkg/utils"
)

type EventHandler func(event models.KafkaOrderEvent) error

var eventHandlers = map[string]EventHandler{
	kafka.Inventory_KafkaEvents[0]:   handleVerifyInventory,
	kafka.Inventory_KafkaEvents[1]: handleRollbackInventory,
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

func handleVerifyInventory(event models.KafkaOrderEvent) error {
	if len(event.Order.Items) == 0 {
		inventory_operations.CancelOrder(event.Order)

		event.Event = kafka.OrdersSEC_KafkaEvents[3]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return errors.New("no items in order: " + event.Order.ID.String())
	}
	
	err := inventory_operations.CheckInventory(event.Order.Items)

	if err != nil {
		inventory_operations.CancelOrder(event.Order)
		
		event.Event = kafka.OrdersSEC_KafkaEvents[3]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return errors.New("insufficient stock for some items in order: " + event.Order.ID.String())
	}

	inventory_update_error := inventory_operations.RemoveStock(event.Order.Items)

	if inventory_update_error != nil {
		inventory_operations.CancelOrder(event.Order)
		
		event.Event = kafka.OrdersSEC_KafkaEvents[3]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return errors.New("error updating stock for some items in order: " + event.Order.ID.String())
	}

	updatedOrder, err := inventory_operations.UpdateOrderStatus(event.Order, "reserved")

	if err != nil {
		inventory_operations.CancelOrder(event.Order)

		event.Event = kafka.OrdersSEC_KafkaEvents[3]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return errors.New("error updating order status to reserved: " + event.Order.ID.String())
	}

	success_kafka_message := models.KafkaOrderEvent{
		Event: kafka.OrdersSEC_KafkaEvents[4],
		Order: updatedOrder,
	}

	sendKafkaEvent(kafka.KafkaTopics["order_sec"], success_kafka_message)

	return nil
}

func handleRollbackInventory(event models.KafkaOrderEvent) error {
	err := inventory_operations.RollbackStock(event.Order.Items)

	if err != nil {

		event.Event = kafka.OrdersSEC_KafkaEvents[1]
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], event)

		return errors.New("error rolling back stock for some items in order: " + event.Order.ID.String())
	}

	inventory_operations.CancelOrder(event.Order)

	event.Event = kafka.OrdersSEC_KafkaEvents[2]

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