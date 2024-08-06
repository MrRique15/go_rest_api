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
		return
	}

	var loggedUser = Services.LogInUser(logginUser)

	if loggedUser.ID == "" {
		user.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Email or Password Inserted"})
		return
	}

	user.IndentedJSON(http.StatusOK, loggedUser)
}
