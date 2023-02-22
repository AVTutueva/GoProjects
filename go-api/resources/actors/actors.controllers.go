package actors

import (
	"net/http"
	"time"
	db "tsi/go-api/database"
	e "tsi/go-api/error"

	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListActors(w http.ResponseWriter, r *http.Request) {
	var actors []*Actor
	db.DB.Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	actor := data.Actor
	db.DB.Create(actor)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewActorResponse(actor))
}

func GetActorByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "actorID")
	var actors []*Actor
	db.DB.First(&actors, id)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func DeleteActorByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "actorID")
	var actor Actor
	result := db.DB.First(&actor, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}
	db.DB.Delete(&actor, id)
}

func UpdateActorByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "actorID")
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	var actor Actor
	result := db.DB.First(&actor, id)

	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	actor.FirstName = strings.ToUpper(data.FirstName)
	actor.LastName = strings.ToUpper(data.LastName)
	actor.LastUpdate = time.Now()

	db.DB.Save(&actor)

}
