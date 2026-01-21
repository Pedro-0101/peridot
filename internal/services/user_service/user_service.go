package user_service

import (
	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/request"
	"github.com/Pedro-0101/peridot/internal/models/response"
	"github.com/Pedro-0101/peridot/internal/repositories"
)

type UserDomainService interface {
	CreateUser(user request.UserRequest) (*response.UserResponse, *resterr.RestErr)
	GetUserById(id string) (*response.UserResponse, *resterr.RestErr)
	GetUserByEmail(email string) (*response.UserResponse, *resterr.RestErr)
	GetAllUsers() ([]*response.UserResponse, *resterr.RestErr)
	UpdateUser(id string, user request.UserRequest) (*response.UserResponse, *resterr.RestErr)
	DeleteUser(id string) *resterr.RestErr
}

type userService struct {
	repo *repositories.UserRepository
}

var _ UserDomainService = (*userService)(nil)

func NewUserService(repo *repositories.UserRepository) UserDomainService {
	return &userService{repo: repo}
}
