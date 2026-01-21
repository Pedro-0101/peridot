package user_service

import (
	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/request"
	"github.com/Pedro-0101/peridot/internal/models/response"
	users "github.com/Pedro-0101/peridot/internal/models/user"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) UpdateUser(id string, newUser request.UserRequest) (*response.UserResponse, *resterr.RestErr) {

	_, err := s.GetUserById(id)
	if err != nil {
		return nil, resterr.NewInternalServerError("Error: Error searching for user")
	}

	hashedPassword, bcryptErr := bcrypt.GenerateFromPassword([]byte(newUser.Pass), bcrypt.DefaultCost)
	if bcryptErr != nil {
		return nil, resterr.NewInternalServerError("Error: Error on processing password")
	}

	updateUser := users.User{
		Name:  newUser.Name,
		Email: newUser.Email,
		Pass:  string(hashedPassword),
	}

	updateErr := s.repo.UpdateUser(id, updateUser)
	if updateErr != nil {
		return nil, resterr.NewInternalServerError("Error: Error updating user")
	}

	return s.GetUserById(id)
}
