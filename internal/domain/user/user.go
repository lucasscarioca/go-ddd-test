package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdatePassword struct {
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(name, email, password string) (*User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// TODO: validate email

	return &User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  string(hashedPass),
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (u *User) ToUserResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}

func (u *UserUpdateRequest) Update(userData *UserResponse) {
	if u.Name == "" {
		u.Name = userData.Name
	}
	if u.Email == "" {
		u.Email = userData.Email
	}
}

func (u *User) Update(userData *UserUpdateRequest) {
	if userData.Name != "" {
		u.Name = userData.Name
	}
	if userData.Email != "" {
		u.Email = userData.Email
	}
}
