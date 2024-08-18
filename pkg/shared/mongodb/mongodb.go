package mongodb

import (
    "fmt"
    "log"
    "time"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(uri string) *mongo.Client  {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	
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

func GetCollection(client *mongo.Client, collectionName string, dataBase string) *mongo.Collection {
    collection := client.Database(dataBase).Collection(collectionName)
    return collection
}