package shipping_operations

import (
	// "errors"

	"errors"

	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/shipping_service/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ordersRepository = repositories.NewOrdersRepository(&repositories.MongoDBHandlerOrders{})
var shippingRepository = repositories.NewShippingRepository(&repositories.PostgresDBHandlerShipping{})

func StartShipping(order models.Order) error {
	order, err := GetOrderById(order)

	if err != nil {
		return err
	}

	if order.Status != "paid" {
		return errors.New("order is not paid")
	}

	_, err = GetShippingByOrderId(order.ID)

	if err != nil {
		return errors.New("shipping register not found for this order")
	}

	return nil
}

func RollbackShipping(order models.Order) error {
	// TODO: rollback shipping in shipping gateway
	// At this moment, without any shipping gateway, we will just return nil to simulate a successful rollback
	
	return nil
}

// ------------------------------------------
// --- Orders Operations ---
// ------------------------------------------
func UpdateOrderStatus(order models.Order, status string) (models.Order, error) {
	order.Status = status

	updatedOrder, err := ordersRepository.UpdateOrder(order)

	if err != nil {
		return models.Order{}, err
	}

	return updatedOrder, nil
}

func GetOrderById(order models.Order) (models.Order, error) {
	order, err := ordersRepository.GetOrderById(order.ID)

	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// ------------------------------------------
// --- Shipping Operations ---
// ------------------------------------------
func GetShippingByOrderId(orderId primitive.ObjectID) (models.Shipping, error) {
	shipping, err := shippingRepository.GetShippingByOrderId(orderId)

	if err != nil {
		return models.Shipping{}, err
	}

	return shipping, nil
}