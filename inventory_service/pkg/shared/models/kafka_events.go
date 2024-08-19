package models

type KafkaOrderEvent struct {
	Event string `json:"event"`
	Order Order  `json:"order"`
}