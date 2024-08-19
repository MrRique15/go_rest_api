package event_handlers

import (
	"fmt"
	"log"

	"github.com/MrRique15/go_rest_api/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/pkg/utils"
	"github.com/MrRique15/go_rest_api/saga_execution_controller/producers"
)

func sendKafkaEvent(topic string, event models.KafkaOrderEvent) error {
	stringOrder, error := utils.OrderToJson(event)

	if error != nil {
		log.Println("error during order transformation for kafka event")
		return error
	}
	
	producers.KafkaProducer.SendKafkaEvent(topic, stringOrder)

	return nil
}

type EventHandler func(event models.KafkaOrderEvent) error

var eventHandlers = map[string]EventHandler{
	"event.order.process": handleProcessOrder,

	"event.inventory.rollback_failed": handleFailedInventoryRollback,
	"event.inventory.rollback_success": handleSuccessInventoryRollback,
	"event.inventory.reservation_failed": handleFailedInventoryReservation,
	"event.inventory.reservation_succeeded": handleSucceededInventoryReservation,


}

func handleProcessOrder(event models.KafkaOrderEvent) error {
	eventMessage := models.KafkaOrderEvent{
		Event: kafka.Inventory_KafkaEvents[0],
		Order: event.Order,
	}

	err := sendKafkaEvent(kafka.KafkaTopics["inventory"], eventMessage)

	if err != nil {
		log.Panic("error sending message to inventory_sec: ", err)
	}

	return nil
}

func handleFailedInventoryRollback(event models.KafkaOrderEvent) error {
	fmt.Println("Rollback failed for order: ", event.Order.ID)
	return nil
}

func handleSuccessInventoryRollback(event models.KafkaOrderEvent) error {
	fmt.Println("Rollback succeeded for order: ", event.Order.ID)
	return nil
}

func handleFailedInventoryReservation(event models.KafkaOrderEvent) error {
	fmt.Println("Reservation failed for order: ", event.Order.ID)
	return nil
}

func handleSucceededInventoryReservation(event models.KafkaOrderEvent) error {
	// TODO: Trigger event to Payment service
	return nil
}

func HandleEvent(event models.KafkaOrderEvent) {
	handler := eventHandlers[event.Event]
	err := handler(event)

	if err != nil {
		fmt.Println("Error handling event: ", err)
	}
}