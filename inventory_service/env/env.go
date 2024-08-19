package env

import (
	"os"
)

func EnvMongoURI() string {
    return os.Getenv("MONGOURI")
}

func EnvMongoDatabase() string {
    return os.Getenv("MONGODATABASE")
}

func EnvKafkaHost() string {
    return os.Getenv("KAFKAHOST")
}