package Routes

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	Controllers "go_rest_api/controllers"
)

type RouterType *gin.Engine

func InitMainRouter() (newRouter *gin.Engine) {
	newRouter = gin.Default()
	return
}

func AlbumsRouter(router *gin.Engine) {
	router.GET("/albums", Controllers.GetAlbums)
	router.GET("/albums/:id", Controllers.GetAlbumByID)
	router.POST("/albums", Controllers.PostAlbums)
}

func UsersRouter(router *gin.Engine) {
	router.POST("/users/login", Controllers.LogInUser)
}

func RunRouter(router *gin.Engine, addr string) (status string) {
	var addressArray = strings.Split(addr, ":")

	if len(addressArray) != 2 {
		fmt.Println("Invalid Host Address Inserted: " + addr)
		return "fail"
	}

	if addressArray[1] == "" || addressArray[1] == " " {
		fmt.Println("Invalid Host Address Inserted: " + addr)
		return "fail"
	}

	router.Run(addr)
	return "sucess"
}
