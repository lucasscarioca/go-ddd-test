package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasscarioca/music-stash-server/api"
)

func Mount(r chi.Router) {
	// Users
	r.Route("/users", UserRoutes)

	// Songs

	// Playlists

	// Groups
}

type controllerFunc func(http.ResponseWriter, *http.Request) (*api.Response, error)

func newHandlerFunc(f controllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := f(w, r)
		if err != nil {
			fmt.Println(err.Error())
		}
		response.WriteJSON(w)
	}
}
