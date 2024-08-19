package inventory_operations

import (
	"errors"
	"fmt"

	"github.com/MrRique15/go_rest_api/main_api/repositories"
	"github.com/MrRique15/go_rest_api/pkg/shared/models"
)

var prodcutsRepository = repositories.NewProductsRepository(&repositories.MongoDBHandlerProducts{})
var ordersRepository = repositories.NewOrdersRepository(&repositories.MongoDBHandlerOrders{})


func CheckInventory(items []models.Item) error {
	for _, item := range items {
		existingItem, err := prodcutsRepository.GetProductById(item.ProductID)

		if err != nil {
			return err
		}

		if existingItem.Stock < item.Quantity {
			return errors.New("insufficient stock for item: " + existingItem.Name)
		}
	}

	return nil
}

func RemoveStock(items []models.Item) error {
	for _, item := range items {
		existingItem, err := prodcutsRepository.GetProductById(item.ProductID)

		if err != nil {
			return err
		}

		existingItem.Stock -= item.Quantity

		updated_product, update_error := prodcutsRepository.UpdateProduct(existingItem)

		if update_error != nil {
			return update_error
		}

		fmt.Println("\n---------Stock Update---------\nProduct: ", updated_product.Name, "\nRemaining Stock: ", updated_product.Stock)
	}

	return nil
}

func RollbackStock(items []models.Item) error {
	for _, item := range items {
		existingItem, err := prodcutsRepository.GetProductById(item.ProductID)

		if err != nil {
			return err
		}

		existingItem.Stock += item.Quantity

		updated_product, update_error := prodcutsRepository.UpdateProduct(existingItem)

		if update_error != nil {
			return update_error
		}

		fmt.Println("\n---------Stock Rollback---------\nProduct: ", updated_product.Name, "\nRemaining Stock: ", updated_product.Stock)
	}

	return nil
}

// -------------------------------------
// Orders Operations
// -------------------------------------
func CancelOrder(order models.Order) error {
	order.Status = "cancelled"

	_, err := ordersRepository.UpdateOrder(order)

	if err != nil {
		return err
	}

	// TODO: send email to customer
	return nil
}

func UpdateOrderStatus(order models.Order, status string) (models.Order, error) {
	order.Status = status

	updatedOrder, err := ordersRepository.UpdateOrder(order)

	if err != nil {
		return models.Order{}, err
	}

	return updatedOrder, nil
}