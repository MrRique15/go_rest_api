package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Services "go_rest_api/services"
	Structs "go_rest_api/structs"
)

func LogInUser(user *gin.Context) {
	var logginUser Structs.LoginUser

	if err := user.BindJSON(&logginUser); err != nil {
		user.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid credentialsd"})
		return
	}

	loggedUser, err := Services.LogInUser(logginUser)

	if err != nil{
		user.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid credentials"})
		return
	}

	user.IndentedJSON(http.StatusOK, loggedUser)
}

func RegisterUser(user *gin.Context){
	var registerUser Structs.RegisterUser

	if err := user.BindJSON(&registerUser); err != nil {
		user.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Insert a valid user to register"})
		return
	}

	if(registerUser.Password != registerUser.ConfirmPassword){
		user.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Passwords does not match"})
		return
	}

	registeredUser, err := Services.RegisterUser(registerUser)

	if err != nil {
		user.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Inserted user can not be registered"})
		return
	}

	user.IndentedJSON(http.StatusCreated, registeredUser)

}