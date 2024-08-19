package producers

import (
	"github.com/MrRique15/go_rest_api/saga_execution_controller/env"
	"github.com/MrRique15/go_rest_api/saga_execution_controller/pkg/shared/kafka"
)

var KafkaProducer *kafka.KafkaProducer

func StartKafkaProducer() {
	KafkaProducer = kafka.NewKafkaProducer(env.EnvKafkaHost())
}