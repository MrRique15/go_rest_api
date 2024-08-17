package order_operations

import (
	"fmt"
	repositories "stock_service/data_base"
	"stock_service/models"
)

var prodcutsRepository = repositories.NewProductsRepository(&repositories.MongoDBHandlerProducts{})

func RemoveStock(item models.Item) error {
	existingItem, err := prodcutsRepository.GetProductById(item.ProductID)

	if err != nil {
		return err
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
			}
		}

		// TODO: Add payment processing status for order

	case "order.event.updated":

		break

	case "order.event.deleted":

		break

	default:
		break
	}
}
