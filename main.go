package main

import (
	"fmt"
	Router "go_rest_api/routes"
)

var defaultAddress = "localhost:8080"

func main() {
	mainRouter := Router.InitMainRouter()

	Router.AlbumsRouter(mainRouter)
	Router.UsersRouter(mainRouter)

	if err := Router.RunRouter(mainRouter, defaultAddress); err != nil {
        fmt.Println("Internal Server Error, Exiting API")
        return
    }
}
