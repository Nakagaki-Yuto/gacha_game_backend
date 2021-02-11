package main

import (

	"go_practice_mvc/server"
	"go_practice_mvc/database"
)

func main() {
	router := server.NewRouter()
	database.ConnectDB()

    router.Logger.Fatal(router.Start(":8080"))
}



