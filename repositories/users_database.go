package Repositores

import (
	"errors"
	Structs "go_rest_api/structs"
)

var Users = []Structs.User{
	{ID: "1", Name: "Admin", Email: "henrique-favaro@hotmail.com", Password: "1234"},
}

func GetUserByEmail(email string) (Structs.User, error) {
	var foundUser Structs.User

	for _, user := range Users {
		if user.Email == email {
			foundUser = user
			return foundUser, nil
		}
	}

	return foundUser, errors.New("user not found")
}

func RegisterUser(user Structs.RegisterUser) (Structs.User, error){
	var userToRegister Structs.User

	userToRegister.ID = user.ID
	userToRegister.Name = user.Name
	userToRegister.Email = user.Email
	userToRegister.Password = user.Password

	Users = append(Users, userToRegister)

	registeredUser := userToRegister

	if(registeredUser.ID == ""){
		return registeredUser, errors.New("error during user registration")
	}
	
	return registeredUser, nil
}
