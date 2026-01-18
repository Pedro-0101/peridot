package usecases

import (
	"errors"

	users "github.com/Pedro-0101/peridot/internal/models/user"
	"github.com/Pedro-0101/peridot/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Repo *repositories.UserRepository
}

func NewUserUseCase(repo *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{Repo: repo}
}

func (uuc *UserUseCase) CreateUser(u *users.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Pass), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Erro ao processar senha")
	}
	u.Pass = string(hashedPassword)

	return uuc.Repo.CreateUser(u)
}

func (uc *UserUseCase) GetAllUsers() ([]users.User, error) {
	return uc.Repo.GetAllUsers()
}
