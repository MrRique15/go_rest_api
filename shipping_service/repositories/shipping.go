package repositories

import (
	// "errors"

	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShippingRepository struct {
	db DBHandlerShipping
}

type DBHandlerShipping interface {
	InitiateCollection()
	listShipping() ([]models.Shipping, error)
	getShippingById(id int) (models.Shipping, error)
	getShippingByOrderId(orderId string) (models.Shipping, error)
	registerShipping(newShipping models.Shipping) (models.Shipping, error)
}

func NewShippingRepository(dbh DBHandlerShipping) *ShippingRepository {
	ShippingRepository := ShippingRepository{
		db: dbh,
	}

	ShippingRepository.db.InitiateCollection()

	return &ShippingRepository
}

// ------------------------------------------------- Functions ----------------------------
func (sr ShippingRepository) RegisterShipping(newShipping models.Shipping) (models.Shipping, error) {
	shipping, err := sr.db.registerShipping(newShipping)

	if err != nil {
		return models.Shipping{}, err
	}

	return shipping, nil
}

func (sr ShippingRepository) GetShippingById(id int) (models.Shipping, error) {
	shipping, err := sr.db.getShippingById(id)

	if err != nil {
		return models.Shipping{}, err
	}

	return shipping, nil
}

func (sr ShippingRepository) GetShippingByOrderId(orderId primitive.ObjectID) (models.Shipping, error) {
	shipping, err := sr.db.getShippingByOrderId(orderId.String())

	if err != nil {
		return models.Shipping{}, err
	}

	return shipping, nil
}

func (sr ShippingRepository) ListShipping() ([]models.Shipping, error) {
	result, err := sr.db.listShipping()

	if err != nil {
		return []models.Shipping{}, err
	}

	return result, nil
}
