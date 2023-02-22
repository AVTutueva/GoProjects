package films

import (
	"net/http"
	"time"
	db "tsi/films_website/database"
	e "tsi/films_website/error"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	// var films []*Film
	// result := db.DB.Find(&films)
	// if result.Error != nil {
	// 	render.Render(w, r, e.ErrEmptyTable(result.Error))
	// 	return
	// }
	// render.RenderList(w, r, NewFilmListResponse(films))

	var films []*Film

	result := db.DB.Model(&Film{}).Preload("Categories").Find(&films)
	// result := db.DB.Model(&Film{}).Preload(clause.Associations).Find(&films)

	if result.Error != nil {
		render.Render(w, r, e.ErrEmptyTable(result.Error))
		return
	}
	render.RenderList(w, r, NewFilmListResponse(films))
}

func FilmByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var film *Film

	// result := db.DB.First(&film, id)
	result := db.DB.Model(&Film{}).Preload("Categories").First(&film, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrFilmDoesNotExist(result.Error))
		return
	}
	render.Render(w, r, NewFilmResponse(film))
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	film := data.Film

	result := db.DB.Create(film)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	// set correct last update time
	film.LastUpdate = time.Now()

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFilmResponse(film))
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var film *Film
	result := db.DB.First(&film, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrFilmDoesNotExist(result.Error))
		return
	}

	db.DB.Delete(&film, id)
	render.Render(w, r, NewFilmResponse(film))
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// read new info
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	// find the film in table
	var updatedFilm *Film
	result := db.DB.First(&updatedFilm, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrFilmDoesNotExist(result.Error))
		return
	}

	updatedFilm.FilmId = data.FilmId
	updatedFilm.Title = data.Title
	updatedFilm.Description = data.Description
	updatedFilm.ReleaseYear = data.ReleaseYear
	updatedFilm.LanguageId = data.LanguageId
	updatedFilm.OriginalLanguageId = data.LanguageId
	updatedFilm.RentalDuration = data.RentalDuration
	updatedFilm.RentalRate = data.RentalRate
	updatedFilm.Length = data.Length
	updatedFilm.ReplacementCost = data.ReplacementCost
	updatedFilm.Rating = data.Rating
	updatedFilm.SpecialFeatures = data.SpecialFeatures
	updatedFilm.LastUpdate = time.Now()
	updatedFilm.Categories = data.Categories

	// for i, _ := range data.Categories {
	// 	updatedFilm.Categories[i].LastUpdate = time.Now()
	// }

	updatedFilm.Categories = data.Categories

	db.DB.Save(&updatedFilm)
	render.Render(w, r, NewFilmResponse(updatedFilm))
}

func FilmsByKeyword(w http.ResponseWriter, r *http.Request) {
	// keyword := chi.URLParam(r, "keyword")
}
