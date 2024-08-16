package repositories

import (
	"context"
	"errors"
	"time"

	"main_api/configs"
	"main_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandlerUsers struct {
	userCollection *mongo.Collection
}

func (dbh *MongoDBHandlerUsers) InitiateCollection() {
	dbh.userCollection = configs.GetCollection(configs.MongoDB, "users")
}

func (dbh MongoDBHandlerUsers) findOneID(id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := dbh.userCollection.FindOne(ctx, bson.M{"id": id}).Decode(&user)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (dbh MongoDBHandlerUsers) findOneEmail(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := dbh.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (dbh MongoDBHandlerUsers) insertOne(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := dbh.userCollection.InsertOne(ctx, user)

	if err != nil {
		return models.User{}, errors.New("error during user registration")
	}

	return dbh.findOneID(user.ID)
}

func (dbh MongoDBHandlerUsers) updateOne(id primitive.ObjectID, user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}

	_, err := dbh.userCollection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": update})

	if err != nil {
		return user, errors.New("error during user update")
	}

	return dbh.findOneID(id)
}
