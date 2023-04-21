package application

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasscarioca/music-stash-server/configs"
	"github.com/lucasscarioca/music-stash-server/internal/application/routes"
	"github.com/lucasscarioca/music-stash-server/internal/domain/user"
)

type server struct {
	r *chi.Mux
}

var Server *server

func NewApp(r *chi.Mux, userRepo user.Repository) *server {
	user.WithUserStorageRepository(userRepo)
	Server = &server{r}
	return Server
}

func (s *server) StartServer() {
	// Middlewares
	s.r.Use(middleware.Logger)

	// Routes
	s.r.Route("/api", routes.Mount)

	fmt.Printf("🚀 Server running on localhost:%s\n", configs.GetPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetPort()), s.r)
}
