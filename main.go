package main

import (

	"go_practice_mvc/server"
	"go_practice_mvc/utilities"
)

func main() {
	router := server.NewRouter()
	utilities.ConnectDB()

    router.Logger.Fatal(router.Start(":8080"))
}



