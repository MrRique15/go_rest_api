package main

import (
	"fmt"

	"github.com/MrRique15/go_rest_api/main_api/repositories"
	router "github.com/MrRique15/go_rest_api/main_api/routes"
)

var defaultAddress = "localhost:8080"

func main() {
	repositories.StartKafkaProducer()

	mainRouter := router.InitMainRouter()

	router.UsersRouter(mainRouter)
	router.OrdersRouter(mainRouter)
	router.ProductsRouter(mainRouter)

	if err := router.RunRouter(mainRouter, defaultAddress); err != nil {
		fmt.Println("Internal Server Error, Exiting API")
		return
	}
}
