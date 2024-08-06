package Router

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	Controllers "go_rest_api/controllers"
)

type RouterType *gin.Engine

func InitMainRouter() *gin.Engine {
	newRouter := gin.Default()
	return newRouter
}

func AlbumsRouter(router *gin.Engine) {
	router.GET("/albums", Controllers.GetAlbums)
	router.GET("/albums/:id", Controllers.GetAlbumByID)
	router.POST("/albums/add", Controllers.PostAlbums)
}

func UsersRouter(router *gin.Engine) {
	router.POST("/users/login", Controllers.LogInUser)
	router.POST("/users/register", Controllers.RegisterUser)
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
