package user_service

import (
	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/response"
)

func (s *userService) GetUserByEmail(email string) (*response.UserResponse, *resterr.RestErr) {

	if email == "" {
		return nil, resterr.NewBadRequestError("Error: Email must be informed")
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, resterr.NewInternalServerError("Error: Error searching for user")
	}

	responseUser := &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return responseUser, nil

}
