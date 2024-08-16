package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
    ID       primitive.ObjectID `json:"id" validate:"required"`
    ClientID primitive.ObjectID `json:"client_id" validate:"required"`
    Price    float32            `json:"price" validate:"required"`
    Items    []Item             `json:"items" validate:"required"`
	Status   string             `json:"status" validate:"required"`
}

type NewOrder struct {
	ClientID string  `json:"client_id" validate:"required"`
	Price    float32 `json:"price" validate:"required"`
	Items    []Item  `json:"items" validate:"required"`
	Status   string  `json:"status" validate:"required"`
}

type UpdateOrder struct {
	ID       string  `json:"id" validate:"required"`
	ClientID string  `json:"client_id" validate:"required"`
	Price    float32 `json:"price" validate:"required"`
	Items    []Item  `json:"items" validate:"required"`
	Status   string  `json:"status" validate:"required"`
}

type Item struct {
	ProductID  primitive.ObjectID `json:"product_id" validate:"required"`
	Quantity   int                `json:"quantity" validate:"required"`
}