package repositories

import (
	"errors"

	"github.com/MrRique15/go_rest_api/main_api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MemoryDBHandlerOrders struct {
	ordersCollection *[]models.Order
}

func (dbh *MemoryDBHandlerOrders) InitiateCollection() {
	dbh.ordersCollection = &[]models.Order{}
}

func (dbh MemoryDBHandlerOrders) getOrderById(id primitive.ObjectID) (models.Order, error) {
	for _, order := range *dbh.ordersCollection {
		if order.ID == id {
			return order, nil
		}
	}

	return models.Order{}, errors.New("order not found")
}

func (dbh MemoryDBHandlerOrders) registerOrder(order models.Order) (models.Order, error) {
	newCollection := *dbh.ordersCollection

	newCollection = append(newCollection, order)

	newOrder := newCollection[len(newCollection)-1]

	if newOrder.ID != order.ID {
		return models.Order{}, errors.New("error during order registration")
	}

	*dbh.ordersCollection = newCollection

	return order, nil
}

func (dbh MemoryDBHandlerOrders) updateOrder(id primitive.ObjectID, order models.Order) (models.Order, error) {
	newDbCollection := *dbh.ordersCollection

	for index, foundOrder := range *dbh.ordersCollection {
		if foundOrder.ID == id {
			newDbCollection[index] = order
		}
	}

	updatedOrder, err := dbh.getOrderById(order.ID)

	if err != nil {
		return models.Order{}, errors.New("error during order update")
	}

	return updatedOrder, nil
}
