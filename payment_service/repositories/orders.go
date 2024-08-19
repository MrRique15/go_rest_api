package repositories

import (
	"errors"
	"log"

	"github.com/MrRique15/go_rest_api/payment_service/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/payment_service/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/payment_service/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrdersRepository struct {
	db DBHandlerOrders
}

type DBHandlerOrders interface {
	InitiateCollection()
	registerOrder(order models.Order) (models.Order, error)
	getOrderById(id primitive.ObjectID) (models.Order, error)
	updateOrder(id primitive.ObjectID, order models.Order) (models.Order, error)
}

func NewOrdersRepository(dbh DBHandlerOrders) *OrdersRepository {
	OrdersRepository := OrdersRepository{
		db: dbh,
	}

	OrdersRepository.db.InitiateCollection()

	return &OrdersRepository
}

// ------------------------------------------------- Functions ----------------------------
func (os OrdersRepository) GetOrderById(id primitive.ObjectID) (models.Order, error) {
	order, err := os.db.getOrderById(id)

	if err != nil {
		return models.Order{}, errors.New("order not found")
	}

	return order, nil
}

func (os OrdersRepository) RegisterOrder(newOrder models.Order) (models.Order, error) {
	order, err := os.db.registerOrder(newOrder)

	if err != nil {
		return models.Order{}, errors.New("error during order registration")
	}

	var OrderEvent = models.KafkaOrderEvent{
		Order: order,
		Event: kafka.OrdersSEC_KafkaEvents[0],
	}

	stringOrder, error := utils.OrderToJson(OrderEvent)

	if error != nil {
		log.Println("error during order transformation for kafka event")
	}

	if error == nil {
		KafkaProducer.SendKafkaEvent(kafka.KafkaTopics["order_sec"], stringOrder)
	}

	return order, nil
}

func (os OrdersRepository) UpdateOrder(order models.Order) (models.Order, error) {
	result, err := os.db.updateOrder(order.ID, order)

	if err != nil {
		return models.Order{}, errors.New("error during order update")
	}

	return result, nil
}
