package payment

import (
	mapper "medical-record-api/internal/payment/mapper"
	model "medical-record-api/internal/payment/model"
	repository "medical-record-api/internal/payment/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll(search string) ([]*model.MethodDTO, error) {
	return s.repo.FindAll(search)
}

func (s *Service) GetAllPaginated(page, limit int, search string) ([]*model.MethodDTO, int, error) {
	return s.repo.FindAllPaginated(page, limit, search)
}

func (s *Service) Create(request *model.MethodRequest) (*model.MethodDTO, error) {
	entity := mapper.MapToMethodEntity(request)

	savedEntity, err := s.repo.Save(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToMethodDTO(savedEntity), nil
}

func (s *Service) Update(id string, request *model.MethodRequest) (*model.MethodDTO, error) {
	entity := mapper.MapToMethodEntity(request)
	entity.ID = id

	updatedEntity, err := s.repo.Update(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToMethodDTO(updatedEntity), nil
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *Service) GetAllStatus() ([]*model.StatusDTO, error) {
	return s.repo.FindAllStatus()
}
