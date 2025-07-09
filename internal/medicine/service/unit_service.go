package medicine

import (
	mapper "medical-record-api/internal/medicine/mapper"
	model "medical-record-api/internal/medicine/model"
	repository "medical-record-api/internal/medicine/repository"
)

type UnitService struct {
	repo *repository.UnitRepository
}

func NewUnitService(repo *repository.UnitRepository) *UnitService {
	return &UnitService{repo}
}

func (s *UnitService) GetAll(search string) ([]*model.UnitDTO, error) {
	return s.repo.FindAll(search)
}

func (s *UnitService) GetAllPaginated(page, limit int, search string) ([]*model.UnitDTO, int, error) {
	return s.repo.FindAllPaginated(page, limit, search)
}

func (s *UnitService) Create(request *model.UnitRequest) (*model.UnitDTO, error) {
	entity := mapper.MapToUnitEntity(request)

	savedEntity, err := s.repo.Save(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToUnitDTO(savedEntity), nil
}

func (s *UnitService) Update(id string, request *model.UnitRequest) (*model.UnitDTO, error) {
	entity := mapper.MapToUnitEntity(request)
	entity.ID = id

	updatedEntity, err := s.repo.Update(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToUnitDTO(updatedEntity), nil
}

func (s *UnitService) Delete(id string) error {
	return s.repo.Delete(id)
}
