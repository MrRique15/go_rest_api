package repositories

import (
	"github.com/MrRique15/go_rest_api/shipping_service/env"
	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/kafka"
	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/mongodb"
	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/postgresdb"
)

var MongoDB = mongodb.ConnectMongoDB(env.EnvMongoURI())
var ShippingPostgresDB = postgresdb.ConnectPostgresDB(env.EnvShippingDatabase(), "shipping")
var KafkaProducer *kafka.KafkaProducer

func StartKafkaProducer() {
	KafkaProducer = kafka.NewKafkaProducer(env.EnvKafkaHost())
}


