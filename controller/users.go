package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lucasscarioca/music-stash-server/api"
	"github.com/lucasscarioca/music-stash-server/dto"
	"github.com/lucasscarioca/music-stash-server/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	req := new(dto.UserRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return api.NewResponse(500, "Internal server error"), err
	}

	userData, err := dto.NewUser(req.Name, req.Email, req.Password)
	if err != nil {
		return api.NewResponse(400, "Invalid User type"), err
	}

	newUser, err := model.CreateUser(userData)
	if err != nil {
		return api.NewResponse(500, "Internal server error"), err
	}

	return api.NewResponse(201, newUser), nil
}

func FindUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return api.NewResponse(200, &dto.UserResponse{ID: uuid.New(), Name: "Test", Email: "test@gmail.com", CreatedAt: time.Now()}), nil
}
