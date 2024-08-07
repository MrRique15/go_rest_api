package repositories

import (
	"time"
	"errors"
	"context"
	"go_rest_api/configs"
	"go_rest_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoDBHandler struct {
	userCollection *mongo.Collection
}

func (dbh *MongoDBHandler) InitiateCollection() {
	dbh.userCollection = configs.GetCollection(configs.DB, "users")
}

func (dbh MongoDBHandler) findOneID(id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := dbh.userCollection.FindOne(ctx, bson.M{"id": id}).Decode(&user)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (dbh MongoDBHandler) findOneEmail(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := dbh.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (dbh MongoDBHandler) insertOne(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := dbh.userCollection.InsertOne(ctx, user)

	if err != nil {
		return models.User{}, errors.New("error during user registration")
	}

	return dbh.findOneID(user.ID)
}

func (dbh MongoDBHandler) updateOne(id primitive.ObjectID, user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}

	_, err := dbh.userCollection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": update})

	if err != nil {
		return user, errors.New("error during user update")
	}

	return dbh.findOneID(id)
}