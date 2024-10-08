package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/MrRique15/go_rest_api/main_api/env"
	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/models"
	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandlerUsers struct {
	userCollection *mongo.Collection
}

func (dbh *MongoDBHandlerUsers) InitiateCollection() {
	dbh.userCollection = mongodb.GetCollection(MongoDB, "users", env.EnvMongoDatabase())
}

func (dbh MongoDBHandlerUsers) findOneID(id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := dbh.userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

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

	_, err := dbh.userCollection.InsertOne(ctx, bson.M{"name": user.Name, "email": user.Email, "password": user.Password})

	if err != nil {
		return models.User{}, errors.New("error during user registration")
	}

	return dbh.findOneEmail(user.Email)
}

func (dbh MongoDBHandlerUsers) updateOne(id primitive.ObjectID, user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}

	_, err := dbh.userCollection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": update})

	if err != nil {
		return user, errors.New("error during user update")
	}

	return dbh.findOneID(id)
}
