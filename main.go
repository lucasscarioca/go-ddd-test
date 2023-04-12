package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const PORT = ":3000"

func main() {
	s := createNewServer()
	s.mountHandlers()
	fmt.Printf("🚀 Server running on localhost%s\n", PORT)
	http.ListenAndServe(PORT, s.Router)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

func createNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) mountHandlers() {
	// Mount all Middlewares here
	s.Router.Use(middleware.Logger)

	// Mount all handlers here
	s.Router.Get("/", helloWorld)
}
