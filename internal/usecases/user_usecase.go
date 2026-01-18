package usecases

import (
	"errors"

	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/request"
	"github.com/Pedro-0101/peridot/internal/models/response"
	users "github.com/Pedro-0101/peridot/internal/models/user"
	"github.com/Pedro-0101/peridot/internal/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Repo *repositories.UserRepository
}

func NewUserUseCase(repo *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{Repo: repo}
}

func (uc *UserUseCase) CreateUser(req *request.UserRequest) (*response.UserResponse, *resterr.RestErr) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, resterr.NewInternalServerError("Error: Error on processing password")
	}

	userEntity := users.User{
		Name:  req.Name,
		Email: req.Email,
		Pass:  string(hashedPassword),
	}

	err = uc.Repo.CreateUser(&userEntity)
	if err != nil {
		return nil, resterr.NewInternalServerError("Error saving user")
	}

	res := &response.UserResponse{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}

	return res, nil
}

func (uc *UserUseCase) GetAllUsers() ([]users.User, error) {
	return uc.Repo.GetAllUsers()
}

func (uc *UserUseCase) GetUserById(id string) (users.User, error) {

	if id == "" {
		err := errors.New("Error: User id must be informed")
		return users.User{}, err
	}

	user, err := uc.Repo.GetUserById(id)

	if err != nil {
		err := errors.New("Error: Error searching for user")
		return users.User{}, err
	}

	if user.ID == uuid.Nil {
		return users.User{}, errors.New("Error: User not found")
	}

	return user, nil
}
