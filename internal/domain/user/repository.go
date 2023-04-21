package user

import (
	"errors"

	"github.com/google/uuid"
)

type Repository interface {
	Create(*User) (*UserResponse, error)
	Find(uuid.UUID) (*UserResponse, error)
	List() ([]*UserResponse, error)
	Update(uuid.UUID, *UserRequest) (*UserResponse, error)
	Delete(uuid.UUID) error
}

var StorageRepo Repository

func WithUserStorageRepository(repository Repository) {
	StorageRepo = repository
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrCreateUser   = errors.New("failed to create user")
	ErrUpdateUser   = errors.New("failed to update user")
)
