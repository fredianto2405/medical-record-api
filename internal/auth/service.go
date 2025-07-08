package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"medical-record-api/internal/constant"
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

	isPasswordMatch := CheckPasswordHash(request.Password, user.Password)
	if !isPasswordMatch {
		return nil, fmt.Errorf(constant.MsgInvalidPassword)
	}

	return user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
