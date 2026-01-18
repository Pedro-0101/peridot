package usecases

import (
	"errors"

	users "github.com/Pedro-0101/peridot/internal/models/user"
	"github.com/Pedro-0101/peridot/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	Repo *repositories.UserRepository
}

func NewCreateUserUseCase(repo *repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{Repo: repo}
}

func (uc *CreateUserUseCase) Execute(u *users.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Pass), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Erro ao processar senha")
	}
	u.Pass = string(hashedPassword)

	return uc.Repo.CreateUser(u)

}
