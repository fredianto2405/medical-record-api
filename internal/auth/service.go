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

	if user.FailedLoginAttempts >= 3 {
		return nil, fmt.Errorf(constant.MsgAccountLocked)
	}

	isPasswordMatch := password.CheckPasswordHash(request.Password, user.Password)
	if !isPasswordMatch {
		isLocked := (user.FailedLoginAttempts + 1) == 3
		err = s.repo.UpdateFailedLoginByEmail(request.Email, isLocked)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(constant.MsgInvalidPassword)
	}

	err = s.repo.ResetFailedLoginByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
