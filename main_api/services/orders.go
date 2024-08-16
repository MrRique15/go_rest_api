package services

import (
	"errors"

	"github.com/MrRique15/go_rest_api/main_api/models"
	"github.com/MrRique15/go_rest_api/main_api/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrdersService struct {
	ordersRepository *repositories.OrdersRepository
}

func NewOrdersService(ord *repositories.OrdersRepository) *OrdersService {
	ordersService := OrdersService{
		ordersRepository: ord,
	}

	return &ordersService
}

func (os OrdersService) RegisterOrder(order models.Order) (models.Order, error) {
	_, err := os.ordersRepository.GetOrderById(order.ID)

	if err == nil {
		return models.Order{}, errors.New("order already registered")
	}

	orderToInsert := models.Order{
		ID:       order.ID,
		ClientID: order.ClientID,
		Price:    order.Price,
		Items:    order.Items,
		Status:   order.Status,
	}

	registeredOrder, err := os.ordersRepository.RegisterOrder(orderToInsert)

	return registeredOrder, err
}

func (os OrdersService) UpdateOrder(order models.Order) (models.Order, error) {
	_, err := os.ordersRepository.GetOrderById(order.ID)

	if err != nil {
		return models.Order{}, errors.New("order not found")
	}

	orderToUpdate := models.Order{
		ID:       order.ID,
		ClientID: order.ClientID,
		Price:    order.Price,
		Items:    order.Items,
		Status:   order.Status,
	}

	updatedOrder, err := os.ordersRepository.UpdateOrder(orderToUpdate)

	return updatedOrder, err
}

func (os OrdersService) GetOrderById(id primitive.ObjectID) (models.Order, error) {

	order, err := os.ordersRepository.GetOrderById(id)

	if err != nil {
		return models.Order{}, errors.New("order not found")
	}

	return order, nil
}
