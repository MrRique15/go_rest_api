package repositories

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/MrRique15/go_rest_api/main_api/models"
	"github.com/MrRique15/go_rest_api/main_api/producers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var kafkaProducer = producers.NewKafkaProducer()

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

func orderToJson(order models.Order) (string, error) {
	jsonOrder, error := json.Marshal(order)

	if error != nil {
		return "", error
	}

	return string(jsonOrder), nil
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

	stringOrder, error := orderToJson(order)

	if error != nil {
		log.Println("error during order transformation for kafka event")
	}

	// send kafka event
	if error == nil {
		kafkaProducer.SendKafkaEvent("order", stringOrder)
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
