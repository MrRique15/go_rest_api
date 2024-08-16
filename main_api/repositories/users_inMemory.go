package repositories

import (
	"errors"
	"main_api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MemoryDBHandlerUsers struct {
	userCollection *[]models.User
}

func (dbh *MemoryDBHandlerUsers) InitiateCollection() {
	dbh.userCollection = &[]models.User{}
}

func (dbh MemoryDBHandlerUsers) findOneID(id primitive.ObjectID) (models.User, error) {
	for _, user := range *dbh.userCollection {
		if user.ID == id {
			return user, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

func (dbh MemoryDBHandlerUsers) findOneEmail(email string) (models.User, error) {
	for _, user := range *dbh.userCollection {
		if user.Email == email {
			return user, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

func (dbh MemoryDBHandlerUsers) insertOne(user models.User) (models.User, error) {
	newCollection := *dbh.userCollection

	newCollection = append(newCollection, user)

	newUser := newCollection[len(newCollection)-1]

	if newUser.Email != user.Email {
		return models.User{}, errors.New("error during user registration")
	}

	*dbh.userCollection = newCollection

	return user, nil
}

func (dbh MemoryDBHandlerUsers) updateOne(id primitive.ObjectID, user models.User) (models.User, error) {
	newDbCollection := *dbh.userCollection

	for index, foundUser := range *dbh.userCollection {
		if foundUser.ID == id {
			newDbCollection[index] = user
		}
	}

	updatedUser, err := dbh.findOneID(user.ID)

	if err != nil {
		return models.User{}, errors.New("error during user update")
	}

	return updatedUser, nil
}
