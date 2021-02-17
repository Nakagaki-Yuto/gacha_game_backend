package main

import (

	"go_practice_mvc/database"
	"go_practice_mvc/server"
	"log"
)

func main() {

	db := database.New()
	handler := handler.New(db)
	server := server.New(handler)

	log.Fatal(server.Run())
}



