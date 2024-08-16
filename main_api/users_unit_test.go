package main

import (
	"testing"

	"main_api/models"
	"main_api/repositories"
	"main_api/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var usersRepository = repositories.NewUsersRepository(&repositories.MemoryDBHandlerUsers{})
var usersService = services.NewUsersService(usersRepository)
var originalUser = models.User{
	Name:     "test user",
	ID:       primitive.NewObjectID(),
	Email:    "teste@teste.com",
	Password: "123456",
}

func TestRegisterUser(t *testing.T) {
	registeredUser, err := usersService.RegisterUser(originalUser)
	if err != nil {
		t.Errorf("Registering User failed with reason: `%s`", err)
	}

	if registeredUser.Email != originalUser.Email {
		t.Errorf("Registering User failed with reason: Registered User does not Match")
	}
}

func TestGetUserById(t *testing.T) {
	gottenUser, err := usersService.GetUserById(originalUser.ID)
	if err != nil {
		t.Errorf("Getting User by ID failed with reason: `%s`", err)
	}

	if gottenUser.Email != originalUser.Email {
		t.Errorf("Getting User by ID failed with reason: Gotten User does not Match")
	}
}

func TestLoginUser(t *testing.T) {
	loginObject := models.UserLogin{
		Email:    originalUser.Email,
		Password: originalUser.Password,
	}

	loggedUser, err := usersService.LoginUser(loginObject)
	if err != nil {
		t.Errorf("Logging User failed with reason: `%s`", err)
	}

	if loggedUser.Email != originalUser.Email {
		t.Errorf("Logging User failed with reason: Logged User does not Match")
	}
}

func TestUpdateUser(t *testing.T) {
	updatingUser := models.EditingUser{
		ID:          originalUser.ID,
		Name:        "teste atualizado",
		Email:       originalUser.Email,
		OldPassword: originalUser.Password,
		NewPassword: originalUser.Password,
	}

	updatedUser, err := usersService.UpdateUser(updatingUser)
	if err != nil {
		t.Errorf("Updating User failed with reason: `%s`", err)
	}

	if updatedUser.Name != updatingUser.Name {
		t.Errorf("Updating User failed with reason: Updated User does not Match")
	}
}
