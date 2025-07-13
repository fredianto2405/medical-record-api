package medical_record

import (
	model "medical-record-api/internal/medical_record/model"
	repository "medical-record-api/internal/medical_record/repository"
)

type StatusService struct {
	repo *repository.StatusRepository
}

func NewStatusService(repo *repository.StatusRepository) *StatusService {
	return &StatusService{repo: repo}
}

func (s *StatusService) GetAll() ([]*model.StatusDTO, error) {
	return s.repo.FindAll()
}
