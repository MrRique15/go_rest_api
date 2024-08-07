package main

import (
	"fmt"
	"go_rest_api/configs"
	"go_rest_api/routes"
)

var defaultAddress = "localhost:8080"

func main() {
	mainRouter := router.InitMainRouter()

	configs.ConnectDB()

	router.UsersRouter(mainRouter)

	if err := router.RunRouter(mainRouter, defaultAddress); err != nil {
        fmt.Println("Internal Server Error, Exiting API")
        return
    }
}
