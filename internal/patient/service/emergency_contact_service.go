package patient

import (
	mapper "medical-record-api/internal/patient/mapper"
	model "medical-record-api/internal/patient/model"
	repository "medical-record-api/internal/patient/repository"
)

type EmergencyContactService struct {
	repo *repository.EmergencyContactRepository
}

func NewEmergencyContactService(repo *repository.EmergencyContactRepository) *EmergencyContactService {
	return &EmergencyContactService{repo}
}

func (s *EmergencyContactService) Create(request *model.EmergencyContactRequest) error {
	entity := mapper.MapToEmergencyContactEntity(request)
	return s.repo.Save(entity)
}

func (s *EmergencyContactService) Delete(patientID string) error {
	return s.repo.Delete(patientID)
}
