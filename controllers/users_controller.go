package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Services "example/web-service-gin/services"
	Structs "example/web-service-gin/structs"
)

func LogInUser(user *gin.Context){
	var logginUser Structs.LoginUser

	if err := user.BindJSON(&logginUser); err != nil {
		return
	}

	var loggedUser = Services.LogInUser(logginUser)

	if(loggedUser.ID == ""){
		user.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Email or Password Inserted"})
		return
	}

	user.IndentedJSON(http.StatusOK, loggedUser)
}