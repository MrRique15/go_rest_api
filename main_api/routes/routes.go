package router

import (
	"errors"
	"strings"

	"main_api/controllers"

	"github.com/gin-gonic/gin"
)

type RouterType *gin.Engine

func InitMainRouter() *gin.Engine {
	newRouter := gin.Default()
	return newRouter
}

func UsersRouter(router *gin.Engine) {
	router.PUT("/users/edit/:id", controllers.EditUser)
	router.GET("/users/:id", controllers.GetUserByID)
	router.POST("/users/register", controllers.RegisterUser)
	router.POST("/users/login", controllers.LoginUser)
}

func OrdersRouter(router *gin.Engine) {
	router.POST("/orders/new", controllers.NewOrder)
	router.PUT("/orders/update", controllers.UpdateOrder)
}

func RunRouter(router *gin.Engine, addr string) error {
	addressArray := strings.Split(addr, ":")

	if len(addressArray) != 2 {
		return errors.New("Invalid Host Address Inserted: " + addr)
	}

	if addressArray[1] == "" || addressArray[1] == " " {
		return errors.New("Invalid Host Address Inserted: " + addr)
	}

	router.Run(addr)
	return nil
}
