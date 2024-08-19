package event_handlers

import (
	"fmt"
	"log"

	"github.com/MrRique15/go_rest_api/saga_execution_controller/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/saga_execution_controller/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/saga_execution_controller/pkg/utils"
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
	kafka.OrdersSEC_KafkaEvents[0]: handleProcessOrder,

	kafka.OrdersSEC_KafkaEvents[1]: handleFailedInventoryRollback,
	kafka.OrdersSEC_KafkaEvents[2]: handleSucceededInventoryRollback,
	kafka.OrdersSEC_KafkaEvents[3]: handleFailedInventoryReservation,
	kafka.OrdersSEC_KafkaEvents[4]: handleSucceededInventoryReservation,

	kafka.OrdersSEC_KafkaEvents[5]: handleFailedPaymentVerification,
	kafka.OrdersSEC_KafkaEvents[6]: handleFailedPaymentRollback,
	kafka.OrdersSEC_KafkaEvents[7]: handleSucceededPaymentRollback,
	kafka.OrdersSEC_KafkaEvents[8]: handleSucceededPaymentVerification,

	kafka.OrdersSEC_KafkaEvents[9]: handleFailedShippingStart,
	kafka.OrdersSEC_KafkaEvents[10]: handleSucceededShippingStart,
	kafka.OrdersSEC_KafkaEvents[11]: handleFailedShippingRollback,
	kafka.OrdersSEC_KafkaEvents[12]: handleSucceededShippingRollback,
}

func handleProcessOrder(event models.KafkaOrderEvent) error {
	eventMessage := models.KafkaOrderEvent{
		Event: kafka.Inventory_KafkaEvents[0],
		Order: event.Order,
	}

	err := sendKafkaEvent(kafka.KafkaTopics["inventory"], eventMessage)

	if err != nil {
		log.Panic("error sending message to inventory topic: ", err)
	}

	return nil
}

// ------------------------------------------
// --- Inventory Handlers ---
// ------------------------------------------
func handleFailedInventoryRollback(event models.KafkaOrderEvent) error {
	fmt.Println("Rollback failed for order: ", event.Order.ID)
	return nil
}

func handleSucceededInventoryRollback(event models.KafkaOrderEvent) error {
	fmt.Println("\n\n----------\nOrder Rollback Completed and Cancelled: ", event.Order.ID, "\n----------")
	return nil
}

func handleFailedInventoryReservation(event models.KafkaOrderEvent) error {
	fmt.Println("Reservation failed for order: ", event.Order.ID)
	return nil
}

func handleSucceededInventoryReservation(event models.KafkaOrderEvent) error {
	event.Event = kafka.Payment_KafkaEvents[0]
	err := sendKafkaEvent(kafka.KafkaTopics["payment"], event)

	if err != nil {
		log.Panic("error sending message to payment topic: ", err)
	}

	return nil
}

// ------------------------------------------
// --- Payment Handlers ---
// ------------------------------------------
func handleFailedPaymentVerification(event models.KafkaOrderEvent) error {
	event.Event = kafka.Inventory_KafkaEvents[1]
	err := sendKafkaEvent(kafka.KafkaTopics["inventory"], event)

	if err != nil {
		log.Panic("error sending message to inventory topic: ", err)
	}

	return nil
}

func handleFailedPaymentRollback(event models.KafkaOrderEvent) error {
	fmt.Println("Shipping Rollback failed for order: ", event.Order.ID)
	event.Event = kafka.Inventory_KafkaEvents[1]
	err := sendKafkaEvent(kafka.KafkaTopics["inventory"], event)

	if err != nil {
		log.Panic("error sending message to inventory topic: ", err)
	}

	return nil
}

func handleSucceededPaymentRollback(event models.KafkaOrderEvent) error {
	event.Event = kafka.Inventory_KafkaEvents[1]
	err := sendKafkaEvent(kafka.KafkaTopics["inventory"], event)

	if err != nil {
		log.Panic("error sending message to inventory topic: ", err)
	}

	return nil
}

func handleSucceededPaymentVerification(event models.KafkaOrderEvent) error {
	event.Event = kafka.Shipping_KafkaEvents[0]
	err := sendKafkaEvent(kafka.KafkaTopics["shipping"], event)

	if err != nil {
		log.Panic("error sending message to shipping topic : ", err)
	}

	return nil
}

// ------------------------------------------
// --- Shipping Handlers ---
// ------------------------------------------
func handleFailedShippingStart(event models.KafkaOrderEvent) error {
	event.Event = kafka.Payment_KafkaEvents[1]
	err := sendKafkaEvent(kafka.KafkaTopics["payment"], event)

	if err != nil {
		log.Panic("error sending message to payment topic: ", err)
	}

	return nil
}

func handleSucceededShippingStart(event models.KafkaOrderEvent) error {
	fmt.Println("\n\n----------\nOrder Processing Completed: ", event.Order.ID, "\n----------")
	return nil
}

func handleFailedShippingRollback(event models.KafkaOrderEvent) error {
	fmt.Println("Shipping Rollback failed for order: ", event.Order.ID)
	event.Event = kafka.Payment_KafkaEvents[1]
	err := sendKafkaEvent(kafka.KafkaTopics["payment"], event)

	if err != nil {
		log.Panic("error sending message to payment topic: ", err)
	}

	return nil
}

func handleSucceededShippingRollback(event models.KafkaOrderEvent) error {
	event.Event = kafka.Payment_KafkaEvents[1]
	err := sendKafkaEvent(kafka.KafkaTopics["payment"], event)

	if err != nil {
		log.Panic("error sending message to payment topic: ", err)
	}

	return nil
}

func HandleEvent(event models.KafkaOrderEvent) {
	handler := eventHandlers[event.Event]
	err := handler(event)

	if err != nil {
		fmt.Println("Error handling event: ", err)
	}
}