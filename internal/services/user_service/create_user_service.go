package user_service

import (
	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/request"
	"github.com/Pedro-0101/peridot/internal/models/response"
	users "github.com/Pedro-0101/peridot/internal/models/user"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) CreateUser(user request.UserRequest) (*response.UserResponse, *resterr.RestErr) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, resterr.NewInternalServerError("Error: Error on processing password")
	}

	userEntity := users.User{
		Name:  user.Name,
		Email: user.Email,
		Pass:  string(hashedPassword),
	}

	err = s.repo.CreateUser(&userEntity)
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
