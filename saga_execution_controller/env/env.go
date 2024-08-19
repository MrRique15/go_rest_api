package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGOURI")
}

func EnvMongoDatabase() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGODATABASE")
}

func EnvKafkaHost() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    return os.Getenv("KAFKAHOST")
}