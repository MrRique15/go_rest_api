package configs

import (
    "fmt"
    "log"
    "time"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client  {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))
	
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}


func ConnectPrismaDB() {
    fmt.Println("Connected to PrismaDB")
}

//Client instance
var MongoDB *mongo.Client = ConnectMongoDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database(EnvMongoDatabase()).Collection(collectionName)
    return collection
}