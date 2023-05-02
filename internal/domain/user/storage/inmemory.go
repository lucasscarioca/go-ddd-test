package storage

import (
	"sync"

	"github.com/google/uuid"
	u "github.com/lucasscarioca/music-stash-server/internal/domain/user"
)

type MemoryRepository struct {
	users map[uuid.UUID]u.User
	sync.Mutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: make(map[uuid.UUID]u.User),
	}
}

func (mr *MemoryRepository) Create(user *u.User) (*u.UserResponse, error) {
	if mr.users == nil {
		mr.Lock()
		mr.users = make(map[uuid.UUID]u.User)
		mr.Unlock()
	}
	if _, ok := mr.users[user.ID]; ok {
		return nil, u.ErrCreateUser
	}
	mr.Lock()
	mr.users[user.ID] = *user
	mr.Unlock()
	return user.ToUserResponse(), nil
}

func (mr *MemoryRepository) Find(id uuid.UUID) (*u.UserResponse, error) {
	if user, ok := mr.users[id]; ok {
		return user.ToUserResponse(), nil
	}

	return nil, u.ErrUserNotFound
}

func (mr *MemoryRepository) List() ([]*u.UserResponse, error) {
	var users []*u.UserResponse
	for _, user := range mr.users {
		users = append(users, user.ToUserResponse())
	}
	return users, nil
}

func (mr *MemoryRepository) Update(id uuid.UUID, user *u.UserUpdateRequest) (*u.UserResponse, error) {
	foundUser, ok := mr.users[id]
	if !ok {
		return nil, u.ErrUpdateUser
	}

	foundUser.Update(user)

	mr.Lock()
	mr.users[id] = foundUser
	mr.Unlock()

	return foundUser.ToUserResponse(), nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.users[id]; !ok {
		return u.ErrUserNotFound
	}
	delete(mr.users, id)
	return nil
}
