package main

import (
	"log"

	"go_practice_mvc/database"
	"go_practice_mvc/handler"
	"go_practice_mvc/server"
)

func main() {

	db := database.New()
	handler := handler.New(db)
	server := server.New(handler)

	log.Fatal(server.Run())
}
