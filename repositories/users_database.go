package Repositores

import (
	Structs "go_rest_api/structs"
)

var Users = []Structs.User{
	{ID: "1", Name: "Admin", Email: "henrique-favaro@hotmail.com", Password: "1234"},
}

func GetUserByEmail(email string) (foundUser Structs.User) {
	for _, user := range Users {
		if user.Email == email {
			foundUser = user
			return
		}
	}
	return
}
