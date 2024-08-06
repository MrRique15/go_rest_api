package Services

import (
	"errors"
	"fmt"

	Repositores "go_rest_api/repositories"
	Structs "go_rest_api/structs"
)

func LogInUser(loginUser Structs.LoginUser) (Structs.User, error) {

	existingUser, err := Repositores.GetUserByEmail(loginUser.Email)

	if err != nil{
		return existingUser, errors.New("invalid credentials")
	}

	fmt.Println(existingUser.Password + " " + loginUser.Password)
	
	if existingUser.Password == loginUser.Password {
		loggedUser := existingUser
		return loggedUser, nil
	}

	return existingUser, errors.New("invalid credentials")
}

func RegisterUser(user Structs.RegisterUser) (Structs.User, error){
	existingUser, err := Repositores.GetUserByEmail(user.Email)

	if err == nil {
		return existingUser, errors.New("email already registered")
	}

	registeredUser, err := Repositores.RegisterUser(user)

	return registeredUser, err
}