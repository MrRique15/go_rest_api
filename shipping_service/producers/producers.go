package producers

import (
	"github.com/MrRique15/go_rest_api/shipping_service/env"
	"github.com/MrRique15/go_rest_api/pkg/shared/kafka"
)

var KafkaProducer *kafka.KafkaProducer

func StartKafkaProducer() {
	KafkaProducer = kafka.NewKafkaProducer(env.EnvKafkaHost())
}