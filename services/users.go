package services

import (
	"errors"

	"go_rest_api/models"
	"go_rest_api/utils"
	"go_rest_api/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(user models.User) (*mongo.InsertOneResult, error) {
	_, err := repositories.GetUserByEmail(user.Email)

	if err == nil {
		return nil, errors.New("email already registered")
	}

	hashedPass, err := utils.GenerateHashFromPassword(user.Password)
	if err != nil{
		return nil, errors.New("error during password encrypt")
	}

	userToInsert := models.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Password: hashedPass,
	}

	registeredUser, err := repositories.RegisterUser(userToInsert)

	return registeredUser, err
}

func GetUserById(id primitive.ObjectID) (models.User, error){
	result, err := repositories.GetUserByID(id)

	if err != nil {
		return result, errors.New("user not found")
	}

	return result, nil
}

func UpdateUser(user models.EditingUser) (*mongo.UpdateResult, error){
	existingUser, err := repositories.GetUserByID(user.ID)

	if err != nil{
		return nil, errors.New("user not found")
	}

	insertedPass := user.OldPassword
	existingPass := existingUser.Password

	err = utils.ComparePasswords(insertedPass, existingPass)

	if err != nil {
		return nil, errors.New("invalid credentials for user update")
	}

	newHashedPass, err := utils.GenerateHashFromPassword(user.NewPassword) 

	if err != nil{
		return nil, errors.New("error during password encrypt")
	}
	
	userToUpdate := models.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Password: newHashedPass,
	}

	result, err := repositories.UpdateUser(userToUpdate)

	if err != nil {
		return nil, errors.New("error updating user")
	}

	return result, nil
}

func LoginUser(loginObject models.UserLogin) (models.User, error){
	existingUser, err := repositories.GetUserByEmail(loginObject.Email)

	if err != nil{
		return existingUser, errors.New("user not found")
	}

	insertedPass := loginObject.Password
	existingPass := existingUser.Password
	error := utils.ComparePasswords(insertedPass, existingPass)

	if error != nil{
		return existingUser, errors.New("invalid credentials")
	}

	return existingUser, nil
}