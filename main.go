package main

import (
	"go_practice_mvc/database"
	"go_practice_mvc/server"
	"log"
)

func main() {
	db := database.New()
	hander := hander.New(db)
	server := server.New(handler)

	// database.ConnectDB()
	// database.ConnectDB()

	user := UserRepo.Find()

	log.Fatal(server.Run())
}
