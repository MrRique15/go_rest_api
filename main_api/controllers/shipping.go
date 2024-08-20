package controllers

import (
	"net/http"

	"github.com/MrRique15/go_rest_api/main_api/repositories"
	"github.com/MrRique15/go_rest_api/main_api/responses"
	"github.com/MrRique15/go_rest_api/main_api/services"

	"github.com/gin-gonic/gin"
)

var shippingRepository = repositories.NewShippingRepository(&repositories.PostgresDBHandlerShipping{})
var shippingService = services.NewShippingService(shippingRepository)

func ListShipping(c *gin.Context) {
	shippingList, err := shippingService.ListShipping()

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ShippingResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.ShippingResponse{Status: http.StatusCreated, Message: "success", Data: &gin.H{"data": shippingList}})
}