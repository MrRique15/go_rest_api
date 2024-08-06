package Services

import (
	Repositores "example/web-service-gin/repositories"
	Structs "example/web-service-gin/structs"
)

func LogInUser(loginUser Structs.LoginUser) (loggedUser Structs.User){

	var existingUser = Repositores.GetUserByEmail(loginUser.Email)

	if(existingUser.Password == loginUser.Password){
		loggedUser = existingUser
		return
	}

	return
}