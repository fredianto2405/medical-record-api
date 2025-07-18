package medical_record

import (
	"github.com/jmoiron/sqlx"
	mapper "medical-record-api/internal/medical_record/mapper"
	model "medical-record-api/internal/medical_record/model"
	repository "medical-record-api/internal/medical_record/repository"
)

type TreatmentDetailService struct {
	db   *sqlx.DB
	repo *repository.TreatmentDetailRepository
}

func NewTreatmentDetailService(db *sqlx.DB, repo *repository.TreatmentDetailRepository) *TreatmentDetailService {
	return &TreatmentDetailService{
		db:   db,
		repo: repo,
	}
}

func (s *TreatmentDetailService) Delete(medicalRecordID, treatmentID string) error {
	return s.repo.Delete(medicalRecordID, treatmentID)
}

func (s *TreatmentDetailService) Create(medicalRecordID string, request model.TreatmentRequest) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			commitErr := tx.Commit()
			if commitErr != nil {
				err = commitErr
			}
		}
	}()

	treatmentEntity := mapper.MapToTreatmentDetailEntity(&request)
	treatmentEntity.MedicalRecordID = medicalRecordID
	err = s.repo.Save(tx, treatmentEntity)

	if err != nil {
		return err
	}

	return nil
}
