package auth

import (
	"fmt"
	"medical-record-api/internal/constant"
	"medical-record-api/pkg/password"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Login(request *LoginRequest) (*UserDTO, error) {
	user, err := s.repo.FindUserByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf(constant.MsgUserNotFound)
	}

	isPasswordMatch := password.CheckPasswordHash(request.Password, user.Password)
	if !isPasswordMatch {
		return nil, fmt.Errorf(constant.MsgInvalidPassword)
	}

	return user, nil
}
