package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/lucasscarioca/music-stash-server/configs"
	"github.com/lucasscarioca/music-stash-server/internal/application"
	userStorage "github.com/lucasscarioca/music-stash-server/internal/domain/user/storage"
	"github.com/lucasscarioca/music-stash-server/internal/infra/db"
)

func main() {
	configs.Load()

	db.Connect()

	httpServer := chi.NewRouter()

	userPostgresRepository := userStorage.NewPostgresRepository()

	app := application.NewApp(httpServer, userPostgresRepository)

	app.StartServer()
}
