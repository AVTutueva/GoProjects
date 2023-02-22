package films

import (
	"net/http"
	"time"
	db "tsi/films_website/database"
	e "tsi/films_website/error"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"strconv"
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
	// check input data
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}
	film := data.Film

	// check that listed categories are exist in category table
	for _, user_category := range film.Categories {
		check_category := db.DB.Where("category_id = ? AND name = ?", user_category.CategoryId, user_category.Name).First(&user_category)
		if check_category.Error != nil {
			render.Render(w, r, e.ErrFilmDoesNotMatchCategory(check_category.Error))
			return
		}
	}

	// set lastUpdate Field
	film.LastUpdate = time.Now()

	// create film
	result := db.DB.Create(film)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFilmResponse(film))
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	// find the film to be deleted

	id := chi.URLParam(r, "id")
	var film *Film
	result := db.DB.First(&film, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrFilmDoesNotExist(result.Error))
		return
	}

	// find film_category relations
	var film_by_category []FilmCategory
	result_fc := db.DB.Where("film_id = ?", id).First(&film_by_category)
	if result_fc.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result_fc.Error))
		println("error")
		return
	}
	// remove associations
	db.DB.Model(&film).Association("Categories").Clear()

	// remove film_category pairs
	result_del := db.DB.Delete(&film_by_category)
	if result_del.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result_fc.Error))
		return
	}

	// remove films
	db.DB.Delete(&film)
	render.Render(w, r, NewFilmResponse(film))
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// read new info
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}

	// find the film in table
	var updatedFilm *Film
	result := db.DB.Where("film_id = ?", id).First(&updatedFilm)

	if result.Error != nil {
		render.Render(w, r, e.ErrFilmDoesNotExist(result.Error))
		return
	}

	// check that listed categories are exist in category table
	for _, user_category := range data.Categories {
		check_category := db.DB.Where("category_id = ? AND name = ?", user_category.CategoryId, user_category.Name).First(&user_category)
		if check_category.Error != nil {
			render.Render(w, r, e.ErrFilmDoesNotMatchCategory(check_category.Error))
			return
		}
	}

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
	updatedFilm.Categories = data.Categories

	url_id, _ := strconv.Atoi(id)
	updatedFilm.FilmId = url_id
	updatedFilm.LastUpdate = time.Now()

	//db.DB.Save(&updatedFilm)
	db.DB.Model(&updatedFilm).Preload("Categories").Save(&updatedFilm)

	render.Render(w, r, NewFilmResponse(updatedFilm))
}

func FilmsByKeyword(w http.ResponseWriter, r *http.Request) {
	// keyword := chi.URLParam(r, "keyword")
}
