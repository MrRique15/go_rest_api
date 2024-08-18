package services

import (
	"errors"

	"github.com/MrRique15/go_rest_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/main_api/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsService struct {
	productsRepository *repositories.ProductsRepository
}

func NewProductsService(ord *repositories.ProductsRepository) *ProductsService {
	productsService := ProductsService{
		productsRepository: ord,
	}

	return &productsService
}

func (os ProductsService) RegisterProduct(product models.Product) (models.Product, error) {
	_, err := os.productsRepository.GetProductById(product.ID)

	if err == nil {
		return models.Product{}, errors.New("product already registered")
	}

	productToInsert := models.Product{
		ID:    product.ID,
		Name:  product.Name,
		Stock: product.Stock,
	}

	registeredProduct, err := os.productsRepository.RegisterProduct(productToInsert)

	return registeredProduct, err
}

func (os ProductsService) UpdateProduct(product models.Product) (models.Product, error) {
	_, err := os.productsRepository.GetProductById(product.ID)

	if err != nil {
		return models.Product{}, errors.New("product not found")
	}

	productToUpdate := models.Product{
		ID:    product.ID,
		Name:  product.Name,
		Stock: product.Stock,
	}

	updatedProduct, err := os.productsRepository.UpdateProduct(productToUpdate)

	return updatedProduct, err
}

func (os ProductsService) GetProductById(id primitive.ObjectID) (models.Product, error) {

	product, err := os.productsRepository.GetProductById(id)

	if err != nil {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}
