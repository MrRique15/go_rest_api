package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type KafkaOrderEvent struct {
	Event string `json:"event"`
	Order Order  `json:"order"`
}

type Order struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	CustomerID primitive.ObjectID `json:"customer_id" validate:"required"`
	Price      float32            `json:"price" validate:"required"`
	Items      []Item             `json:"items" validate:"required"`
	Status     string             `json:"status" validate:"required"`
}

type Item struct {
	ProductID primitive.ObjectID `json:"product_id" validate:"required"`
	Quantity  int                `json:"quantity" validate:"required"`
}