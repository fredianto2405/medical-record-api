package user

import "medical-record-api/pkg/password"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(request *Request) error {
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return err
	}

	entity := MapToEntity(request)
	entity.Password = hashedPassword

	return s.repo.Save(entity)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *Service) EmailExists(email string) (bool, error) {
	return s.repo.EmailExists(email)
}
