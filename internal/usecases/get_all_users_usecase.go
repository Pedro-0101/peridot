package usecases

import (
	users "github.com/Pedro-0101/peridot/internal/models/user"
	"github.com/Pedro-0101/peridot/internal/repositories"
)

type GetAllUsersUseCase struct {
	Repo *repositories.UserRepository
}

func NewGetAllUserUseCase(repo *repositories.UserRepository) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{Repo: repo}
}

func (uc *GetAllUsersUseCase) Execute() ([]users.User, error) {
	return uc.Repo.GetAllUsers()
}
