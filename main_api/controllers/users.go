package controllers

import (
	"net/http"

	"main_api/models"
	"main_api/repositories"
	"main_api/responses"
	"main_api/services"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

var usersRepository = repositories.NewUsersRepository(&repositories.MongoDBHandlerUsers{})
var usersService = services.NewUsersService(usersRepository)

func RegisterUser(c *gin.Context) {
	var user models.RegisterUser

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": validationErr.Error()}})
		return
	}

	if user.Password != user.ConfirmPassword {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": "passwords does not match"}})
		return
	}

	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	result, err := usersService.RegisterUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &gin.H{"data": result}})
}

func GetUserByID(c *gin.Context) {
	userId := c.Param("id")
	var user models.User

	objId, _ := primitive.ObjectIDFromHex(userId)

	user, err := usersService.GetUserById(objId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	returnedUser := models.ReturnedUser{
		ID:    primitive.NewObjectID(),
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &gin.H{"data": returnedUser}})
}

func EditUser(c *gin.Context) {
	userId := c.Param("id")
	var user models.EditingUser

	objId, _ := primitive.ObjectIDFromHex(userId)

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": validationErr.Error()}})
		return
	}

	gottenUser, err := usersService.GetUserById(objId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	_, error := usersService.UpdateUser(user)
	if error != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	updatedUser := models.ReturnedUser{
		ID:    gottenUser.ID,
		Name:  gottenUser.Name,
		Email: gottenUser.Email,
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &gin.H{"data": updatedUser}})
}

func LoginUser(c *gin.Context) {
	var loginObject models.UserLogin

	if err := c.BindJSON(&loginObject); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&loginObject); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": validationErr.Error()}})
		return
	}

	user, err := usersService.LoginUser(loginObject)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &gin.H{"data": err.Error()}})
		return
	}

	returnedUser := models.ReturnedUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "sucess", Data: &gin.H{"data": returnedUser}})
}
