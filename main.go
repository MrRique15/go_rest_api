package main

import (
	"fmt"
	Routes "go_rest_api/routes"
)

var defaultAddress = "localhost:8080"

func main() {
	var mainRouter Routes.RouterType = Routes.InitMainRouter()

	Routes.AlbumsRouter(mainRouter)
	Routes.UsersRouter(mainRouter)

	var runStatus = Routes.RunRouter(mainRouter, defaultAddress)

	if runStatus == "fail" {
		fmt.Println("Internal Server Error, Exiting API")
		return
	}
}
