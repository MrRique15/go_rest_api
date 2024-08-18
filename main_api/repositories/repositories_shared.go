package repositories

import (
	"github.com/MrRique15/go_rest_api/main_api/env"
	"github.com/MrRique15/go_rest_api/pkg/shared/kafka"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoDB mongo.Client
var KafkaProducer *kafka.KafkaProducer

func SetMongoDB(mongo_client *mongo.Client) {
	MongoDB = *mongo_client
}

func StartKafkaProducer() {
	KafkaProducer = kafka.NewKafkaProducer(env.EnvKafkaHost())
}


