package storage

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	u "github.com/lucasscarioca/music-stash-server/internal/domain/user"
	"github.com/lucasscarioca/music-stash-server/internal/infra/db"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{
		db: db.GetDbConnection(),
	}
}

func (pr *PostgresRepository) Create(user *u.User) (*u.UserResponse, error) {
	query := `INSERT into users
	(id, name, email, password, created_at)
	values ($1, $2, $3, $4, $5)`

	row := pr.db.QueryRow(
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
	)

	var createdUser u.UserResponse
	err := row.Scan(createdUser.ID, createdUser.Name, createdUser.Email, createdUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (pr *PostgresRepository) Find(id uuid.UUID) (*u.UserResponse, error) {
	return nil, errors.New("not implemented")
}

func (pr *PostgresRepository) List() ([]*u.UserResponse, error) {
	return nil, errors.New("not implemented")
}

func (pr *PostgresRepository) Update(id uuid.UUID, user *u.UserRequest) (*u.UserResponse, error) {
	return nil, errors.New("not implemented")
}

func (pr *PostgresRepository) Delete(id uuid.UUID) error {
	return errors.New("not implemented")
}
