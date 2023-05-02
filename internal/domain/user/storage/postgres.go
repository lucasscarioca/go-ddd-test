package storage

import (
	"database/sql"

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
	query := `INSERT INTO users
	(id, name, email, password, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, name, email, created_at;`

	row := pr.db.QueryRow(
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
	)

	var createdUser u.UserResponse
	err := row.Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email, &createdUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (pr *PostgresRepository) Find(id uuid.UUID) (*u.UserResponse, error) {
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1;`

	row := pr.db.QueryRow(query, id)

	var foundUser u.UserResponse
	err := row.Scan(&foundUser.ID, &foundUser.Name, &foundUser.Email, &foundUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func (pr *PostgresRepository) List() ([]*u.UserResponse, error) {
	query := `SELECT id, name, email, created_at FROM users;` // TODO: apply limit/pagination

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []*u.UserResponse

	for rows.Next() {
		var user u.UserResponse

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			continue // Ideally we should log these errors
		}

		users = append(users, &user)
	}

	return users, nil
}

func (pr *PostgresRepository) Update(id uuid.UUID, user *u.UserUpdateRequest) (*u.UserResponse, error) {
	query := `UPDATE users
	SET (name, email) = ($2, $3)
	WHERE id = $1
	RETURNING id, name, email, created_at;`

	row := pr.db.QueryRow(query, id, user.Name, user.Email)

	var updatedUser u.UserResponse
	err := row.Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email, &updatedUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (pr *PostgresRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1;`

	res, err := pr.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return u.ErrUserNotFound
	}

	return nil
}
