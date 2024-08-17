package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
    ID    primitive.ObjectID `bson:"_id" json:"_id"`
    Name  string             `json:"name" validate:"required"`
	Stock int                `json:"stock" validate:"required"`
}

type UpdateProduct struct {
	ID    string `json:"_id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}