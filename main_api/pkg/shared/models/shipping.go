package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Shipping struct {
	ID primitive.ObjectID `json:"_id"`
	ShippingType string `json:"shipping_type"`
	OrderID primitive.ObjectID `json:"order_id"`
	ShippingPrice float64 `json:"shipping_price"`
}