package model

import (
	"github.com/lucasscarioca/music-stash-server/db"
	"github.com/lucasscarioca/music-stash-server/dto"
)

func CreateUser(user *dto.User) (*dto.UserResponse, error) {
	query := `INSERT into users
	(id, name, email, password, created_at)
	values ($1, $2, $3, $4, $5)`

	row := db.Conn.QueryRow(
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
	)

	var createdUser dto.UserResponse
	err := row.Scan(createdUser.ID, createdUser.Name, createdUser.Email, createdUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
