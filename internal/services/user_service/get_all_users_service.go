package user_service

import (
	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
	"github.com/Pedro-0101/peridot/internal/models/response"
)

func (s *userService) GetAllUsers() ([]*response.UserResponse, *resterr.RestErr) {

	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userResponses := []*response.UserResponse{}

	for _, user := range *users {
		userResponse := &response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}
