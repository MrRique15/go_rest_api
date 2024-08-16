package services

import (
	"errors"

	"github.com/MrRique15/go_rest_api/main_api/models"
	"github.com/MrRique15/go_rest_api/main_api/repositories"
	"github.com/MrRique15/go_rest_api/main_api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usr *repositories.UsersRepository) *UsersService {
	usersService := UsersService{
		usersRepository: usr,
	}

	return &usersService
}

func (us UsersService) RegisterUser(user models.User) (models.User, error) {
	_, err := us.usersRepository.GetUserByEmail(user.Email)

	if err == nil {
		return models.User{}, errors.New("email already registered")
	}

	hashedPass, err := utils.GenerateHashFromPassword(user.Password)
	if err != nil {
		return models.User{}, errors.New("error during password encrypt")
	}

	userToInsert := models.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPass,
	}

	registeredUser, err := us.usersRepository.RegisterUser(userToInsert)

	return registeredUser, err
}

func (us UsersService) GetUserById(id primitive.ObjectID) (models.User, error) {
	user, err := us.usersRepository.GetUserByID(id)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (us UsersService) UpdateUser(user models.EditingUser) (models.User, error) {
	existingUser, err := us.usersRepository.GetUserByID(user.ID)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	insertedPass := user.OldPassword
	existingPass := existingUser.Password

	err = utils.ComparePasswords(insertedPass, existingPass)

	if err != nil {
		return models.User{}, errors.New("invalid credentials for user update")
	}

	newHashedPass, err := utils.GenerateHashFromPassword(user.NewPassword)

	if err != nil {
		return models.User{}, errors.New("error during password encrypt")
	}

	userToUpdate := models.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: newHashedPass,
	}

	updatedUser, err := us.usersRepository.UpdateUser(userToUpdate)

	if err != nil {
		return models.User{}, errors.New("error updating user")
	}

	return updatedUser, nil
}

func (us UsersService) LoginUser(loginObject models.UserLogin) (models.User, error) {
	existingUser, err := us.usersRepository.GetUserByEmail(loginObject.Email)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	insertedPass := loginObject.Password
	existingPass := existingUser.Password
	error := utils.ComparePasswords(insertedPass, existingPass)

	if error != nil {
		return models.User{}, errors.New("invalid credentials")
	}

	return existingUser, nil
}
