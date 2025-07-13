package medical_record

import repository "medical-record-api/internal/medical_record/repository"

type TreatmentDetailService struct {
	repo *repository.TreatmentDetailRepository
}

func NewTreatmentDetailService(repo *repository.TreatmentDetailRepository) *TreatmentDetailService {
	return &TreatmentDetailService{repo: repo}
}

func (s *TreatmentDetailService) Delete(medicalRecordID, treatmentID string) error {
	return s.repo.Delete(medicalRecordID, treatmentID)
}
