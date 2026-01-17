package repositories

import "github.com/Pedro-0101/peridot/internal/models/users"

type Repositories struct {
	User interface {
		GetAll() ([]users.User, error)
		Add(newUser users.User) (bool, error)
	}
}

func New() *Repositories {
	return &Repositories{}
}
