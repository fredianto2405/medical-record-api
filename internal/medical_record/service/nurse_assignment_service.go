package medical_record

import (
	"github.com/jmoiron/sqlx"
	mapper "medical-record-api/internal/medical_record/mapper"
	repository "medical-record-api/internal/medical_record/repository"
)

type NurseAssignmentService struct {
	db   *sqlx.DB
	repo *repository.NurseAssignmentRepository
}

func NewNurseAssignmentService(db *sqlx.DB, repo *repository.NurseAssignmentRepository) *NurseAssignmentService {
	return &NurseAssignmentService{
		db:   db,
		repo: repo,
	}
}

func (s *NurseAssignmentService) Delete(medicalRecordID, nurseID string) error {
	return s.repo.Delete(medicalRecordID, nurseID)
}

func (s *NurseAssignmentService) Create(medicalRecordID, nurseID string) error {
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

	nurseAssignmentEntity := mapper.MapToNurseAssignmentEntity(medicalRecordID, nurseID)
	err = s.repo.Save(tx, nurseAssignmentEntity)
	if err != nil {
		return err
	}

	return nil
}
