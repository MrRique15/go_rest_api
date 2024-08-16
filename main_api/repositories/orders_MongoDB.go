package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/MrRique15/go_rest_api/main_api/configs"
	"github.com/MrRique15/go_rest_api/main_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandlerOrders struct {
	ordersCollection *mongo.Collection
}

func (dbh *MongoDBHandlerOrders) InitiateCollection() {
	dbh.ordersCollection = configs.GetCollection(configs.MongoDB, "orders")
}

func (dbh MongoDBHandlerOrders) getOrderById(id primitive.ObjectID) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order models.Order
	err := dbh.ordersCollection.FindOne(ctx, bson.M{"id": id}).Decode(&order)

	if err != nil {
		return models.Order{}, errors.New("order not found")
	}

	return order, nil
}

func (dbh MongoDBHandlerOrders) registerOrder(order models.Order) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := dbh.ordersCollection.InsertOne(ctx, order)

	if err != nil {
		return models.Order{}, errors.New("error during order registration")
	}

	return dbh.getOrderById(order.ID)
}

func (dbh MongoDBHandlerOrders) updateOrder(id primitive.ObjectID, order models.Order) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"client_id": order.ClientID,
		"price":     order.Price,
		"items":     order.Items,
		"status":    order.Status,
	}

	_, err := dbh.ordersCollection.UpdateOne(ctx, bson.M{"id": order.ID}, bson.M{"$set": update})

	if err != nil {
		return models.Order{}, errors.New("error during order update")
	}

	return dbh.getOrderById(id)
}
