package Services

import (
	Repositores "go_rest_api/repositories"
	Structs "go_rest_api/structs"
)

func LogInUser(loginUser Structs.LoginUser) (loggedUser Structs.User) {

	var existingUser = Repositores.GetUserByEmail(loginUser.Email)

	if existingUser.Password == loginUser.Password {
		loggedUser = existingUser
		return
	}

	return
}
