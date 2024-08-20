package services

import (
	"errors"

	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/main_api/repositories"
)

type ShippingService struct {
	shippingRepository *repositories.ShippingRepository
}

func NewShippingService(ship *repositories.ShippingRepository) *ShippingService {
	shippingService := ShippingService{
		shippingRepository: ship,
	}

	return &shippingService
}

func (ss ShippingService) ListShipping() ([]models.Shipping, error) {
	shippingList, err := ss.shippingRepository.ListShipping()

	if err != nil {
		return []models.Shipping{}, errors.New("error during shipping list")
	}

	return shippingList, nil
}