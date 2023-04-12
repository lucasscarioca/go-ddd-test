package routes

import "github.com/go-chi/chi/v5"

func Mount(r chi.Router) {
	// Users
	r.Route("/users", UserRoutes)

	// Songs

	// Playlists

	// Groups
}
