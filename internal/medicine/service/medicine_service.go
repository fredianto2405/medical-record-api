package medicine

import (
	mapper "medical-record-api/internal/medicine/mapper"
	model "medical-record-api/internal/medicine/model"
	repository "medical-record-api/internal/medicine/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll(search string) ([]*model.DTO, error) {
	return s.repo.FindAll(search)
}

func (s *Service) GetAllPaginated(page, limit int, search string) ([]*model.DTO, int, error) {
	return s.repo.FindAllPaginated(page, limit, search)
}

func (s *Service) Create(request *model.Request) (*model.DTO, error) {
	entity := mapper.MapToEntity(request)

	savedEntity, err := s.repo.Save(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToDTO(savedEntity), nil
}

func (s *Service) Update(id string, request *model.Request) (*model.DTO, error) {
	entity := mapper.MapToEntity(request)
	entity.ID = id

	updatedEntity, err := s.repo.Update(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToDTO(updatedEntity), nil
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
