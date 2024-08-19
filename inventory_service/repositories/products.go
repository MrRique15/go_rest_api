package repositories

import (
	"errors"

	"github.com/MrRique15/go_rest_api/inventory_service/pkg/shared/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsRepository struct {
	db DBHandlerProducts
}

type DBHandlerProducts interface {
	InitiateCollection()
	registerProduct(product models.Product) (models.Product, error)
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

func (pr ProductsRepository) RegisterProduct(newProduct models.Product) (models.Product, error) {
	product, err := pr.db.registerProduct(newProduct)

	if err != nil {
		return models.Product{}, errors.New("error during product registration")
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
