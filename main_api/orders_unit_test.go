package main

import (
	"testing"

	"github.com/MrRique15/go_rest_api/main_api/models"
	"github.com/MrRique15/go_rest_api/main_api/repositories"
	"github.com/MrRique15/go_rest_api/main_api/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ordersRepository = repositories.NewOrdersRepository(&repositories.MemoryDBHandlerOrders{})
var ordersService = services.NewOrdersService(ordersRepository)

var originalCustomer = models.User{
	Name:     "test customer",
	ID:       primitive.NewObjectID(),
	Email:    "customer@test.com",
	Password: "123456",
}

var originalOrder = models.Order{
	ID:       primitive.NewObjectID(),
	ClientID: originalCustomer.ID,
	Price:    100.0,
	Items: []models.Item{
		{
			ProductID: primitive.NewObjectID(),
			Quantity:  1,
		},
	},
	Status: "pending",
}

func TestRegisterOrder(t *testing.T) {
	registeredUser, err := usersService.RegisterUser(originalCustomer)
	if err != nil {
		t.Errorf("Registering User failed with reason: `%s`", err)
	}

	if registeredUser.Email != originalCustomer.Email {
		t.Errorf("Registering User failed with reason: Registered User does not Match")
	}

	registeredOrder, err := ordersService.RegisterOrder(originalOrder)
	if err != nil {
		t.Errorf("Registering Order failed with reason: `%s`", err)
	}

	if registeredOrder.ID != originalOrder.ID {
		t.Errorf("Registering Order failed with reason: Registered Order does not Match")
	}
}

func TestGetOrderByID(t *testing.T) {
	gottenOrder, err := ordersService.GetOrderById(originalOrder.ID)
	if err != nil {
		t.Errorf("Getting Order by ID failed with reason: `%s`", err)
	}

	if gottenOrder.ID != originalOrder.ID {
		t.Errorf("Getting Order by ID failed with reason: Gotten Order does not Match")
	}
}

func TestUpdateOrder(t *testing.T) {
	updatingOrder := models.Order{
		ID:       originalOrder.ID,
		ClientID: originalOrder.ClientID,
		Price:    200.0,
		Items: []models.Item{
			{
				ProductID: primitive.NewObjectID(),
				Quantity:  2,
			},
		},
		Status: "completed",
	}

	updatedOrder, err := ordersService.UpdateOrder(updatingOrder)
	if err != nil {
		t.Errorf("Updating Order failed with reason: `%s`", err)
	}

	if updatedOrder.Price != updatedOrder.Price {
		t.Errorf("Updating Order failed with reason: Updated User does not Match")
	}
}
