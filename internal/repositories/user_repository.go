package repositories

import (
	"database/sql"

	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
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

func (r *UserRepository) GetAllUsers() (*[]users.User, *resterr.RestErr) {
	query := `SELECT id, username, user_email, created_at, updated_at FROM users`

	rows, err := r.connection.Query(query)
	if err != nil {
		return nil, resterr.NewInternalServerError("Error searching for users")
	}
	defer rows.Close()

	var allUsers []users.User

	for rows.Next() {
		var u users.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, resterr.NewInternalServerError("Error searching for users")
		}
		allUsers = append(allUsers, u)
	}

	if err = rows.Err(); err != nil {
		return nil, resterr.NewInternalServerError("Error searching for users")
	}

	return &allUsers, nil
}

func (r *UserRepository) GetUserById(id string) (users.User, error) {
	query := `SELECT id, username, user_email, created_at, updated_at FROM users WHERE id = $1`

	var u users.User

	err := r.connection.QueryRow(query, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return users.User{}, nil
		}
		return users.User{}, err
	}

	return u, nil
}

func (r *UserRepository) GetUserByEmail(email string) (users.User, error) {
	query := `SELECT id, username, user_email, created_at, updated_at FROM users WHERE user_email = $1`

	var u users.User

	err := r.connection.QueryRow(query, email).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return users.User{}, err
	}

	return u, nil
}

func (r *UserRepository) DeleteUser(id string) error {

	query := `DELETE FROM users WHERE id = $1`

	_, err := r.connection.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(id string, user users.User) error {
	query := `UPDATE users SET username = $1, user_email = $2, password_hash = $3 WHERE id = $4`

	_, err := r.connection.Exec(query, user.Name, user.Email, user.Pass, id)
	if err != nil {
		return err
	}

	return nil
}
