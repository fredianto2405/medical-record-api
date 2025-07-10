package patient

import (
	mapper "medical-record-api/internal/patient/mapper"
	model "medical-record-api/internal/patient/model"
	repository "medical-record-api/internal/patient/repository"
)

type InsurancePatientService struct {
	repo *repository.InsurancePatientRepository
}

func NewInsurancePatientService(repo *repository.InsurancePatientRepository) *InsurancePatientService {
	return &InsurancePatientService{repo}
}

func (s *InsurancePatientService) Create(request *model.InsurancePatientRequest) error {
	entity := mapper.MapToInsurancePatientEntity(request)
	return s.repo.Save(entity)
}

func (s *InsurancePatientService) Delete(patientID string) error {
	return s.repo.Delete(patientID)
}
