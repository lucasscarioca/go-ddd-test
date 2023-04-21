package routes

import (
	"github.com/go-chi/chi/v5"
	userService "github.com/lucasscarioca/music-stash-server/internal/domain/user/services"
)

func Mount(r chi.Router) {
	// Users
	r.Route("/users", userService.UserRoutes)

	// Songs

	// Playlists

	// Groups
}
