package controllers

import (
	"net/http"

	"github.com/MrRique15/go_rest_api/main_api/models"
	"github.com/MrRique15/go_rest_api/main_api/repositories"
	"github.com/MrRique15/go_rest_api/main_api/responses"
	"github.com/MrRique15/go_rest_api/main_api/services"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validateOrder = validator.New()

var ordersRepository = repositories.NewOrdersRepository(&repositories.MongoDBHandlerOrders{})
var ordersService = services.NewOrdersService(ordersRepository)

func NewOrder(c *gin.Context) {
	var order models.NewOrder

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	if validationErr := validateOrder.Struct(&order); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": validationErr.Error()}})
		return
	}

	client_id, err := primitive.ObjectIDFromHex(order.ClientID)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "invalid client_id"}})
		return
	}

	_, error := usersService.GetUserById(client_id)

	if error != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "customer not found"}})
		return
	}

	newOrder := models.Order{
		ID:       primitive.NewObjectID(),
		ClientID: client_id,
		Price:    order.Price,
		Items:    order.Items,
		Status:   order.Status,
	}

	result, err := ordersService.RegisterOrder(newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.OrderResponse{Status: http.StatusCreated, Message: "success", Data: &gin.H{"data": result}})
}

func UpdateOrder(c *gin.Context) {
	var order models.UpdateOrder

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	if validationErr := validateOrder.Struct(&order); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": validationErr.Error()}})
		return
	}

	order_id, error := primitive.ObjectIDFromHex(order.ID)

	if error != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "invalid client_id"}})
		return
	}

	client_id, err := primitive.ObjectIDFromHex(order.ClientID)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "invalid client_id"}})
		return
	}

	_, error_customer := usersService.GetUserById(client_id)

	if error_customer != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "customer not found"}})
		return
	}

	updatingOrder := models.Order{
		ID:       order_id,
		ClientID: client_id,
		Price:    order.Price,
		Items:    order.Items,
		Status:   order.Status,
	}

	_, error_update := ordersService.UpdateOrder(updatingOrder)
	if error_update != nil {
		c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	gottenOrder, err := ordersService.GetOrderById(order_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	result := models.Order{
		ID:       gottenOrder.ID,
		ClientID: gottenOrder.ClientID,
		Price:    gottenOrder.Price,
		Items:    gottenOrder.Items,
		Status:   gottenOrder.Status,
	}

	c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: &gin.H{"data": result}})
}
