package controllers

import (
	"net/http"

	"main_api/models"
	"main_api/repositories"
	"main_api/responses"
	"main_api/services"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validateProduct = validator.New()

var productsRepository = repositories.NewProductsRepository(&repositories.MongoDBHandlerProducts{})
var productsService = services.NewProductsService(productsRepository)

func NewProduct(c *gin.Context) {
	var product models.NewProduct

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	if validationErr := validateProduct.Struct(&product); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": validationErr.Error()}})
		return
	}

	newProduct := models.Product{
		ID:       primitive.NewObjectID(),
		Name:     product.Name,
		Stock:    product.Stock,
	}

	result, err := productsService.RegisterProduct(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.ProductResponse{Status: http.StatusCreated, Message: "success", Data: &gin.H{"data": result}})
}

func UpdateProduct(c *gin.Context) {
	var product models.UpdateProduct

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	if validationErr := validateProduct.Struct(&product); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": validationErr.Error()}})
		return
	}

	product_id, err := primitive.ObjectIDFromHex(product.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "invalid product_id"}})
		return
	}

	updatingProduct := models.Product{
		ID:       product_id,
		Name:     product.Name,
		Stock:    product.Stock,
	}

	_, error_update := productsService.UpdateProduct(updatingProduct)
	if error_update != nil {
		c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": error_update}})
		return
	}

	gottenProduct, err := productsService.GetProductById(product_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	result := models.Product{
		ID:       gottenProduct.ID,
		Name:     gottenProduct.Name,
		Stock:    gottenProduct.Stock,
	}

	c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: &gin.H{"data": result}})
}
