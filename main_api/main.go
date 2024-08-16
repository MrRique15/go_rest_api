package main

import (
	"fmt"
	"main_api/configs"
	router "main_api/routes"
)

var defaultAddress = "localhost:8080"

func main() {
	mainRouter := router.InitMainRouter()

	configs.ConnectMongoDB()
	configs.ConnectPrismaDB()

	router.UsersRouter(mainRouter)
	router.OrdersRouter(mainRouter)

	if err := router.RunRouter(mainRouter, defaultAddress); err != nil {
		fmt.Println("Internal Server Error, Exiting API")
		return
	}
}
