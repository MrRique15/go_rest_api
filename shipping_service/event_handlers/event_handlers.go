package event_handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/shipping_service/pkg/utils"
	"github.com/MrRique15/go_rest_api/shipping_service/producers"
	"github.com/MrRique15/go_rest_api/shipping_service/shipping_operations"
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

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[9],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "SHIPPING_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}
		
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return err
	}

	updated_order, err := shipping_operations.UpdateOrderStatus(event.Order, "shipped")

	if err != nil {
		fmt.Println("error updating order status for order: ", event.Order.ID)

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[9],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "SHIPPING_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return err
	}

	eventMessage := models.KafkaOrderEvent{
		Event: kafka.OrdersSEC_KafkaEvents[10],
		Order: updated_order,
		CreatedAt: time.Now().String(),
		Source: "SHIPPING_SERVICE",
		Status: "PENDING",
		EventHistory: append(event.EventHistory, models.KafkaEventHistory{
			Source: event.Source,
			Status: "SUCCESS",
			CreatedAt: event.CreatedAt,
			Event: event.Event,
		}),
	}

	sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

	return nil
}

func handleRollbackShipping(event models.KafkaOrderEvent) error {
	err := shipping_operations.RollbackShipping(event.Order)

	if err != nil {
		fmt.Println("error rolling back shipping for order: ", event.Order.ID)

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[11],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "SHIPPING_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return err
	}

	updated_order, err := shipping_operations.UpdateOrderStatus(event.Order, "on_hold")

	if err != nil {
		fmt.Println("error updating order status for order: ", event.Order.ID)

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[11],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "SHIPPING_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return err
	}

	eventMessage := models.KafkaOrderEvent{
		Event: kafka.OrdersSEC_KafkaEvents[12],
		Order: updated_order,
		CreatedAt: time.Now().String(),
		Source: "SHIPPING_SERVICE",
		Status: "PENDING",
		EventHistory: append(event.EventHistory, models.KafkaEventHistory{
			Source: event.Source,
			Status: "SUCCESS",
			CreatedAt: event.CreatedAt,
			Event: event.Event,
		}),
	}

	sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

	return nil
}

func HandleEvent(event models.KafkaOrderEvent) {
	handler := eventHandlers[event.Event]
	err := handler(event)

	if err != nil {
		fmt.Println("Error handling event: ", err)
	}
}