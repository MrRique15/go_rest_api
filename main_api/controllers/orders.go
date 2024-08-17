package controllers

import (
	"fmt"
	"net/http"

	"main_api/models"
	"main_api/repositories"
	"main_api/responses"
	"main_api/services"

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

	customer_id, err := primitive.ObjectIDFromHex(order.CustomerID)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "invalid customer_id"}})
		return
	}

	_, error := usersService.GetUserById(customer_id)

	if error != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "customer not found"}})
		return
	}

	newOrder := models.Order{
		ID:       primitive.NewObjectID(),
		CustomerID: customer_id,
		Price:    order.Price,
		Items:    order.Items,
		Status:   order.Status,
	}

	result, err := ordersService.RegisterOrder(newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	//print result in JSON format with keys and values
	fmt.Printf("%+v\n", result)

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
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "invalid customer_id"}})
		return
	}

	customer_id, err := primitive.ObjectIDFromHex(order.CustomerID)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "invalid customer_id"}})
		return
	}

	_, error_customer := usersService.GetUserById(customer_id)

	if error_customer != nil {
		c.JSON(http.StatusBadRequest, responses.OrderResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "customer not found"}})
		return
	}

	updatingOrder := models.Order{
		ID:       order_id,
		CustomerID: customer_id,
		Price:    order.Price,
		Items:    order.Items,
		Status:   order.Status,
	}

	_, error_update := ordersService.UpdateOrder(updatingOrder)
	if error_update != nil {
		c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err}})
		return
	}

	gottenOrder, err := ordersService.GetOrderById(order_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	result := models.Order{
		ID:       gottenOrder.ID,
		CustomerID: gottenOrder.CustomerID,
		Price:    gottenOrder.Price,
		Items:    gottenOrder.Items,
		Status:   gottenOrder.Status,
	}

	c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: &gin.H{"data": result}})
}
