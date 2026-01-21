package user_service

import (
	resterr "github.com/Pedro-0101/peridot/configuration/rest_err"
)

func (s *userService) DeleteUser(id string) *resterr.RestErr {

	if id == "" {
		return resterr.NewBadRequestError("Error: User id must be informed")
	}

	err := s.repo.DeleteUser(id)
	if err != nil {
		return resterr.NewInternalServerError("Error: Error deleting user")
	}

	return nil
}
