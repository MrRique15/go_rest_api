package models

type Shipping struct {
	ID int `json:"_id"`
	ShippingType string `json:"shipping_type"`
	OrderID string `json:"order_id"`
	ShippingPrice float64 `json:"shipping_price"`
}