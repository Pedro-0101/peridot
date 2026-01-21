package user_service

import (
	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/response"
	"github.com/google/uuid"
)

func (s *userService) GetUserById(id string) (*response.UserResponse, *resterr.RestErr) {

	if id == "" {
		return nil, resterr.NewBadRequestError("Error: User id must be informed")
	}

	if _, err := uuid.Parse(id); err != nil {
		return nil, resterr.NewBadRequestError("Error: Invalid ID format")
	}

	user, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, resterr.NewInternalServerError("Error: Error searching for user")
	}

	if user.ID.String() == "" {
		return nil, resterr.NewNotFoundError("Error: User not found")
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
