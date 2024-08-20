package services

import (
	"errors"
	"fmt"

	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/main_api/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrdersService struct {
	ordersRepository   *repositories.OrdersRepository
	shippingRepository *repositories.ShippingRepository
}

func NewOrdersService(ord *repositories.OrdersRepository, ship *repositories.ShippingRepository) *OrdersService {
	ordersService := OrdersService{
		ordersRepository:   ord,
		shippingRepository: ship,
	}

	return &ordersService
}

func (os OrdersService) RegisterOrder(order models.Order) (models.Order, error) {
	_, err := os.ordersRepository.GetOrderById(order.ID)

	if err == nil {
		return models.Order{}, errors.New("order already registered")
	}

	orderToInsert := models.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Price:      order.Price,
		Items:      order.Items,
		Status:     order.Status,
	}

	registeredOrder, err := os.ordersRepository.RegisterOrder(orderToInsert)

	if err != nil {
		return models.Order{}, errors.New("error during order registration")
	}

	newShippingRegister := models.Shipping{
		ID:            0,
		ShippingType:  "standard",
		OrderID:       order.ID.String(),
		ShippingPrice: float64(order.Price) * 0.1,
	}

	_, err = os.shippingRepository.RegisterShipping(newShippingRegister)

	if err != nil {
		fmt.Println(err)
	}

	return registeredOrder, err
}

func (os OrdersService) UpdateOrder(order models.Order) (models.Order, error) {
	_, err := os.ordersRepository.GetOrderById(order.ID)

	if err != nil {
		return models.Order{}, errors.New("order not found")
	}

	orderToUpdate := models.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Price:      order.Price,
		Items:      order.Items,
		Status:     order.Status,
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
