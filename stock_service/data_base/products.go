package repositories

import (
	"errors"

	"stock_service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsRepository struct {
	db DBHandlerProducts
}

type DBHandlerProducts interface {
	InitiateCollection()
	getProductById(id primitive.ObjectID) (models.Product, error)
	updateProduct(id primitive.ObjectID, product models.Product) (models.Product, error)
}

func NewProductsRepository(dbh DBHandlerProducts) *ProductsRepository {
	ProductsRepository := ProductsRepository{
		db: dbh,
	}

	ProductsRepository.db.InitiateCollection()

	return &ProductsRepository
}

// ------------------------------------------------- Functions ----------------------------
func (pr ProductsRepository) GetProductById(id primitive.ObjectID) (models.Product, error) {
	product, err := pr.db.getProductById(id)

	if err != nil {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}

func (pr ProductsRepository) UpdateProduct(product models.Product) (models.Product, error) {
	result, err := pr.db.updateProduct(product.ID, product)

	if err != nil {
		return models.Product{}, errors.New("error during product update")
	}

	return result, nil
}
