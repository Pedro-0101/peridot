package repositories

import (
	"database/sql"

	users "github.com/Pedro-0101/peridot/internal/models/user"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		connection: db,
	}
}

func (r *UserRepository) CreateUser(u *users.User) error {
	query := `INSERT INTO users (username, user_email, password_hash) 
              VALUES ($1, $2, $3) 
              RETURNING id, created_at, updated_at`

	return r.connection.QueryRow(
		query,
		u.Name,
		u.Email,
		u.Pass,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}
