package main

import (
	"tsi/go-api/database"
	"tsi/go-api/resources/actors"
	"tsi/go-api/server"
)

func main() {
	database.Init()
	database.DB.AutoMigrate(&actors.Actor{})

	server.Init()

}
