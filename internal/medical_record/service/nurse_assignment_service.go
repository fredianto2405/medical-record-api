package medical_record

import repository "medical-record-api/internal/medical_record/repository"

type NurseAssignmentService struct {
	repo *repository.NurseAssignmentRepository
}

func NewNurseAssignmentService(repo *repository.NurseAssignmentRepository) *NurseAssignmentService {
	return &NurseAssignmentService{repo: repo}
}

func (s *NurseAssignmentService) Delete(medicalRecordID, nurseID string) error {
	return s.repo.Delete(medicalRecordID, nurseID)
}
