package medicine

import (
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
