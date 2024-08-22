package event_handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/MrRique15/go_rest_api/payment_service/payment_operations"
	"github.com/MrRique15/go_rest_api/payment_service/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/payment_service/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/payment_service/pkg/utils"
	"github.com/MrRique15/go_rest_api/payment_service/producers"
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
		
		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[5],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "PAYMENT_SERVICE",
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

	updated_order, err := payment_operations.UpdateOrderStatus(event.Order, "paid")

	if err != nil {
		fmt.Println("Error updating order status for order: ", event.Order.ID)

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[5],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "PAYMENT_SERVICE",
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
		Event: kafka.OrdersSEC_KafkaEvents[8],
		Order: updated_order,
		CreatedAt: time.Now().String(),
		Source: "PAYMENT_SERVICE",
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

func handleRollbackPayment(event models.KafkaOrderEvent) error {
	err := payment_operations.RollbackPayment(event.Order)

	if err != nil {
		fmt.Println("Error rolling back payment for order: ", event.Order.ID)

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[6],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "PAYMENT_SERVICE",
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

	updatedOrder, err := payment_operations.UpdateOrderStatus(event.Order, "refounded")

	if err != nil {
		fmt.Println("Error updating order status for order: ", event.Order.ID)

		eventMessage := models.KafkaOrderEvent{
			Event: kafka.OrdersSEC_KafkaEvents[6],
			Order: event.Order,
			CreatedAt: time.Now().String(),
			Source: "PAYMENT_SERVICE",
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
		Event: kafka.OrdersSEC_KafkaEvents[7],
		Order: updatedOrder,
		CreatedAt: time.Now().String(),
		Source: "PAYMENT_SERVICE",
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