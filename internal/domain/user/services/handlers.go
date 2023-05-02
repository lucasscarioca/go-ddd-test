package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	app "github.com/lucasscarioca/music-stash-server/internal/application/utils"
	u "github.com/lucasscarioca/music-stash-server/internal/domain/user"
)

func UserRoutes(r chi.Router) {
	r.Get("/{id}", app.NewHandler(find))

	r.Get("/", app.NewHandler(list))

	r.Post("/", app.NewHandler(create))

	r.Put("/{id}", app.NewHandler(update))

	r.Delete("/{id}", app.NewHandler(delete))
}

func create(w http.ResponseWriter, r *http.Request) error {
	req := new(u.UserCreateRequest)

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

func find(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	foundUser, err := u.StorageRepo.Find(id)
	if err != nil {
		return err
	}

	return app.Send(w, 200, foundUser)
}

func list(w http.ResponseWriter, r *http.Request) error {
	users, err := u.StorageRepo.List()
	if err != nil {
		return err
	}

	return app.Send(w, 200, users)
}

func update(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	foundUser, err := u.StorageRepo.Find(id)
	if err != nil {
		return err
	}

	userData := new(u.UserUpdateRequest)
	if err = json.NewDecoder(r.Body).Decode(userData); err != nil {
		return err
	}

	userData.Update(foundUser)

	updatedUser, err := u.StorageRepo.Update(id, userData)
	if err != nil {
		return err
	}

	return app.Send(w, 200, updatedUser)
}

func delete(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	err = u.StorageRepo.Delete(id)
	if err != nil {
		return err
	}

	return app.Send(w, 204, nil)
}
