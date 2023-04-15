package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/lucasscarioca/music-stash-server/controller"
)

func UserRoutes(r chi.Router) {
	r.Get("/", newHandlerFunc(controller.FindUser))

	r.Post("/", newHandlerFunc(controller.CreateUser))
}
