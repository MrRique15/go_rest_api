package repositories

import (
	"errors"

	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/models"
)

type MemoryDBHandlerShipping struct {
	shippingDatabase *[]models.Shipping
}

func (dbh *MemoryDBHandlerShipping) InitiateCollection() {
	dbh.shippingDatabase = &[]models.Shipping{}
}

func (dbh MemoryDBHandlerShipping) getShippingById(id int) (models.Shipping, error) {
	for _, shipping := range *dbh.shippingDatabase {
		if shipping.ID == id {
			return shipping, nil
		}
	}

	return models.Shipping{}, errors.New("shipping register not found")
}

func (dbh MemoryDBHandlerShipping) registerShipping(newShipping models.Shipping) (models.Shipping, error) {
	newShipping.ID = len(*dbh.shippingDatabase) + 1

	newCollection := *dbh.shippingDatabase

	newCollection = append(newCollection, newShipping)

	newShippingRegister := newCollection[len(newCollection)-1]

	if newShippingRegister.ID != newShipping.ID {
		return models.Shipping{}, errors.New("error during shipping registration")
	}

	*dbh.shippingDatabase = newCollection

	return newShipping, nil
}

func (dbh MemoryDBHandlerShipping) getShippingByOrderId(orderId string) (models.Shipping, error) {
	for _, shipping := range *dbh.shippingDatabase {
		if shipping.OrderID == orderId {
			return shipping, nil
		}
	}

	return models.Shipping{}, errors.New("shipping register not found")
}

func (dbh MemoryDBHandlerShipping) listShipping() ([]models.Shipping, error) {
	return *dbh.shippingDatabase, nil
}