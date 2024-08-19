package repositories

import (
	"github.com/MrRique15/go_rest_api/main_api/env"
	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/mongodb"
)

var MongoDB = mongodb.ConnectMongoDB(env.EnvMongoURI())
var KafkaProducer *kafka.KafkaProducer

func StartKafkaProducer() {
	KafkaProducer = kafka.NewKafkaProducer(env.EnvKafkaHost())
}


