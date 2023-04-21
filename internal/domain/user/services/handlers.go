package services

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	app "github.com/lucasscarioca/music-stash-server/internal/application/utils"
	u "github.com/lucasscarioca/music-stash-server/internal/domain/user"
)

func UserRoutes(r chi.Router) {
	r.Get("/", app.NewHandler(findUser))

	r.Post("/", app.NewHandler(createUser))
}

func createUser(w http.ResponseWriter, r *http.Request) error {
	req := new(u.UserRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	userData, err := u.NewUser(req.Name, req.Email, req.Password)
	if err != nil {
		return err
	}

	newUser, err := u.StorageRepo.Create(userData)
	if err != nil {
		return err
	}

	return app.Send(w, 201, newUser)
}

func findUser(w http.ResponseWriter, r *http.Request) error {
	return app.Send(w, 200, &u.UserResponse{ID: uuid.New(), Name: "Test", Email: "test@gmail.com", CreatedAt: time.Now()})
}
