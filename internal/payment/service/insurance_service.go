package payment

import (
	mapper "medical-record-api/internal/payment/mapper"
	model "medical-record-api/internal/payment/model"
	repository "medical-record-api/internal/payment/repository"
)

type InsuranceService struct {
	repo *repository.InsuranceRepository
}

func NewInsuranceService(repo *repository.InsuranceRepository) *InsuranceService {
	return &InsuranceService{repo: repo}
}

func (s *InsuranceService) GetAll(search string) ([]*model.InsuranceDTO, error) {
	return s.repo.FindAll(search)
}

func (s *InsuranceService) GetAllPaginated(page, limit int, search string) ([]*model.InsuranceDTO, int, error) {
	return s.repo.FindAllPaginated(page, limit, search)
}

func (s *InsuranceService) Create(request *model.InsuranceRequest) (*model.InsuranceDTO, error) {
	entity := mapper.MapToInsuranceEntity(request)

	savedEntity, err := s.repo.Save(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToInsuranceDTO(savedEntity), nil
}

func (s *InsuranceService) Update(id string, request *model.InsuranceRequest) (*model.InsuranceDTO, error) {
	entity := mapper.MapToInsuranceEntity(request)
	entity.ID = id

	updatedEntity, err := s.repo.Update(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToInsuranceDTO(updatedEntity), nil
}

func (s *InsuranceService) Delete(id string) error {
	return s.repo.Delete(id)
}
