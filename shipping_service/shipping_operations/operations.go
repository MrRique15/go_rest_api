package shipping_operations

import (
	// "errors"

	"errors"

	"github.com/MrRique15/go_rest_api/main_api/repositories"
	"github.com/MrRique15/go_rest_api/pkg/shared/models"
)

var ordersRepository = repositories.NewOrdersRepository(&repositories.MongoDBHandlerOrders{})


func StartShipping(order models.Order) error {
	order, err := GetOrderById(order)

	if err != nil {
		return err
	}

	if order.Status != "paid" {
		return errors.New("order is not paid")
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