package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasscarioca/music-stash-server/configs"
	"github.com/lucasscarioca/music-stash-server/routes"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}

	app := chi.NewRouter()

	// Middlewares
	app.Use(middleware.Logger)

	// Routes
	app.Route("/api", routes.Mount)

	fmt.Printf("🚀 Server running on localhost:%s\n", configs.GetServerEnv().PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerEnv().PORT), app)
}
