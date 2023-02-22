package main

import (
	"tsi/films_website/database"
	"tsi/films_website/resources/films"
	"tsi/films_website/server"
)

func main() {
	database.Init()

	database.DB.AutoMigrate(&films.Film{})
	database.DB.AutoMigrate(&films.Category{})

	server.Init()

}
