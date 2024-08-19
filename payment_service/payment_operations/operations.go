package payment_operations

import (
	// "errors"

	"errors"

	"github.com/MrRique15/go_rest_api/payment_service/repositories"
	"github.com/MrRique15/go_rest_api/payment_service/pkg/shared/models"
)

var ordersRepository = repositories.NewOrdersRepository(&repositories.MongoDBHandlerOrders{})


func CheckPayment(order models.Order) error {
	// TODO: check payment in payment gateway
	// At this moment, without any payment gateway, we will just return nil to simulate a successful payment

	order, err := GetOrderById(order)

	if err != nil {
		return err
	}

	if order.Status != "reserved" {
		return errors.New("order is not reserved")
	}

	return nil
}

func RollbackPayment(order models.Order) error {
	// TODO: rollback payment in payment gateway
	// At this moment, without any payment gateway, we will just return nil to simulate a successful rollback
	
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