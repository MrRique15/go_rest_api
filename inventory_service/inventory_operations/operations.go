package inventory_operations

import (
	"errors"
	"fmt"

	"github.com/MrRique15/go_rest_api/main_api/repositories"
	"github.com/MrRique15/go_rest_api/pkg/shared/models"
)

var prodcutsRepository = repositories.NewProductsRepository(&repositories.MongoDBHandlerProducts{})

func RemoveStock(item models.Item) error {
	existingItem, err := prodcutsRepository.GetProductById(item.ProductID)

	if err != nil {
		return err
	}

	if existingItem.Stock < item.Quantity {
		return errors.New("insufficient stock for item: " + existingItem.Name)
	}

	existingItem.Stock -= item.Quantity

	updated_product, update_error := prodcutsRepository.UpdateProduct(existingItem)

	if update_error != nil {
		return update_error
	}

	fmt.Println("\n----------------------------\nProduct: ", updated_product.Name, "\nRemaining Stock: ", updated_product.Stock)

	return nil
}

func ProcessOrder(completeEvent models.KafkaOrderEvent) {
	switch completeEvent.Event {
	case "order.event.created":
		// Remove stock for each item
		for _, item := range completeEvent.Order.Items {
			err := RemoveStock(item)

			if err != nil {
				fmt.Println("Error removing stock for item: ", item)
				// TODO: Add rollback event for order in SAGA pattern
			}
		}

		// TODO: Add event for order processed in SAGA pattern

	case "order.event.updated":

		break

	case "order.event.deleted":

		break

	default:
		break
	}
}
