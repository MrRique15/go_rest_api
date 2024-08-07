package repositories

import (
	"context"
	"errors"
	"go_rest_api/configs"
	"go_rest_api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func GetUserByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err := userCollection.FindOne(ctx, bson.M{"id": id}).Decode(&user)

	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func RegisterUser(newUser models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection.InsertOne(ctx, newUser)

	if err != nil {
		return result, errors.New("error during user registration")
	}

	return result, nil
}

func UpdateUser(user models.User)(*mongo.UpdateResult, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    update := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}
    result, err := userCollection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": update})

	if err != nil{
		return result, errors.New("error during user update")
	}

	return result, nil
}