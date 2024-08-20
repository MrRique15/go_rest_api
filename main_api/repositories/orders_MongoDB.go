package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MrRique15/go_rest_api/main_api/env"
	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"database/sql"
)

type MongoDBHandlerOrders struct {
	ordersCollection *mongo.Collection
	shippingDatabase *sql.DB
}

func (dbh *MongoDBHandlerOrders) InitiateCollection() {
	dbh.ordersCollection = mongodb.GetCollection(MongoDB, "orders", env.EnvMongoDatabase())
	// TODO: Implement shipping repository in separate file later
	postgress_database, err := sql.Open("postgres", env.EnvShippingDatabase())

	if err != nil {
		panic("Error during shipping database connection")
	}

	dbh.shippingDatabase = postgress_database
}

func (dbh MongoDBHandlerOrders) getOrderById(id primitive.ObjectID) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order models.Order
	err := dbh.ordersCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

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

	newShippingRegister := models.Shipping{
		ID:           primitive.NewObjectID(),
		ShippingType: "standard",
		OrderID:      order.ID,
		ShippingPrice: float64(order.Price) * 0.1,
	}
	
	sql_query := `INSERT INTO shipping (_id, shipping_type, order_id, shipping_price) VALUES ($1, $2, $3, $4) RETURNING order_id`

	var order_id primitive.ObjectID

	error_sql := dbh.shippingDatabase.QueryRow(sql_query, newShippingRegister.ID, newShippingRegister.ShippingType, newShippingRegister.OrderID, newShippingRegister.ShippingPrice).Scan(&order_id)

	if error_sql != nil {
		fmt.Println("Error during shipping registration: ", error_sql)
	}

	return dbh.getOrderById(order.ID)
}

func (dbh MongoDBHandlerOrders) updateOrder(id primitive.ObjectID, order models.Order) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"customer_id": order.CustomerID,
		"price":     order.Price,
		"items":     order.Items,
		"status":    order.Status,
	}

	_, err := dbh.ordersCollection.UpdateOne(ctx, bson.M{"_id": order.ID}, bson.M{"$set": update})

	if err != nil {
		return models.Order{}, errors.New("error during order update")
	}

	return dbh.getOrderById(id)
}
