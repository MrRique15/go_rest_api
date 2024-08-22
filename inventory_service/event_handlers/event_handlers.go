package event_handlers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/MrRique15/go_rest_api/inventory_service/inventory_operations"
	"github.com/MrRique15/go_rest_api/inventory_service/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/inventory_service/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/inventory_service/pkg/utils"
	"github.com/MrRique15/go_rest_api/inventory_service/producers"
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
		
		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[3],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "INVENTORY_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}
		
		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return errors.New("no items in order: " + event.Order.ID.String())
	}
	
	err := inventory_operations.CheckInventory(event.Order.Items)

	if err != nil {
		inventory_operations.CancelOrder(event.Order)
		
		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[3],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "INVENTORY_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return errors.New("insufficient stock for some items in order: " + event.Order.ID.String())
	}

	inventory_update_error := inventory_operations.RemoveStock(event.Order.Items)

	if inventory_update_error != nil {
		inventory_operations.CancelOrder(event.Order)

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[3],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "INVENTORY_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return errors.New("error updating stock for some items in order: " + event.Order.ID.String())
	}

	updatedOrder, err := inventory_operations.UpdateOrderStatus(event.Order, "reserved")

	if err != nil {
		inventory_operations.CancelOrder(event.Order)

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[3],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "INVENTORY_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return errors.New("error updating order status to reserved: " + event.Order.ID.String())
	}

	eventMessage := models.KafkaOrderEvent{
		Event: kafka.OrdersSEC_KafkaEvents[4],
		Order: updatedOrder,
		CreatedAt: time.Now().String(),
		Source: "INVENTORY_SERVICE",
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

func handleRollbackInventory(event models.KafkaOrderEvent) error {
	err := inventory_operations.RollbackStock(event.Order.Items)

	if err != nil {
		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[1],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "INVENTORY_SERVICE",
			Status: "PENDING",
			EventHistory: append(event.EventHistory, models.KafkaEventHistory{
				Source: event.Source,
				Status: "SUCCESS",
				CreatedAt: event.CreatedAt,
				Event: event.Event,
			}),
		}

		sendKafkaEvent(kafka.KafkaTopics["order_sec"], eventMessage)

		return errors.New("error rolling back stock for some items in order: " + event.Order.ID.String())
	}

	inventory_operations.CancelOrder(event.Order)

	eventMessage := models.KafkaOrderEvent{
		Event: kafka.OrdersSEC_KafkaEvents[2],
		Order: event.Order,
		CreatedAt: time.Now().String(),
		Source: "INVENTORY_SERVICE",
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