package utils

import (
	"encoding/json"
	"stock_service/models"
)

func ByteToInterface(data []byte) models.KafkaOrderEvent {
	var event models.KafkaOrderEvent
	json.Unmarshal(data, &event)
	return event
}