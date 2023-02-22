package server

import (
	"tsi/films_website/resources/films"

	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Mount("/films", films.Routes())

	return router
}
