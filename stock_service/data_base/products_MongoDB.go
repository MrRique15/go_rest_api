package repositories

import (
	"context"
	"errors"
	"time"

	"stock_service/configs"
	"stock_service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandlerProducts struct {
	productsCollection *mongo.Collection
}

func (dbh *MongoDBHandlerProducts) InitiateCollection() {
	dbh.productsCollection = configs.GetCollection(configs.MongoDB, "products")
}

func (dbh MongoDBHandlerProducts) getProductById(id primitive.ObjectID) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product models.Product
	err := dbh.productsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)

	if err != nil {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}

func (dbh MongoDBHandlerProducts) updateProduct(id primitive.ObjectID, product models.Product) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"name":  product.Name,
		"stock": product.Stock,
	}

	_, err := dbh.productsCollection.UpdateOne(ctx, bson.M{"_id": product.ID}, bson.M{"$set": update})

	if err != nil {
		return models.Product{}, errors.New("error during product update")
	}

	return dbh.getProductById(id)
}
