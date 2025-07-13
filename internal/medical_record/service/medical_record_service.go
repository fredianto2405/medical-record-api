package medical_record

import (
	"github.com/jmoiron/sqlx"
	"medical-record-api/internal/constant"
	mapper "medical-record-api/internal/medical_record/mapper"
	model "medical-record-api/internal/medical_record/model"
	repository "medical-record-api/internal/medical_record/repository"
)

type Service struct {
	db                  *sqlx.DB
	repo                *repository.Repository
	nurseAssignmentRepo *repository.NurseAssignmentRepository
	treatmentDetailRepo *repository.TreatmentDetailRepository
	recipeRepo          *repository.RecipeRepository
	historyRepo         *repository.HistoryRepository
}

func NewService(db *sqlx.DB,
	repo *repository.Repository,
	nurseAssignmentRepository *repository.NurseAssignmentRepository,
	treatmentDetailRepo *repository.TreatmentDetailRepository,
	recipeRepository *repository.RecipeRepository,
	historyRepo *repository.HistoryRepository) *Service {
	return &Service{
		db:                  db,
		repo:                repo,
		nurseAssignmentRepo: nurseAssignmentRepository,
		treatmentDetailRepo: treatmentDetailRepo,
		recipeRepo:          recipeRepository,
		historyRepo:         historyRepo,
	}
}

func (s *Service) Create(request *model.Request) (string, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return "", err
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

	// save medical record
	statusID := constant.MedicalRecordStatusAppointment
	entity := mapper.MapToEntity(request)
	entity.StatusID = statusID
	medicalRecordID, err := s.repo.Save(tx, entity)
	if err != nil {
		return "", err
	}

	// save nurse assignment
	for _, nurseID := range request.NurseIDs {
		nurseAssignmentEntity := mapper.MapToNurseAssignmentEntity(medicalRecordID, nurseID)
		err = s.nurseAssignmentRepo.Save(tx, nurseAssignmentEntity)
		if err != nil {
			return "", err
		}
	}

	// save treatment detail
	for _, treatment := range request.Treatments {
		treatmentEntity := mapper.MapToTreatmentDetailEntity(&treatment)
		treatmentEntity.MedicalRecordID = medicalRecordID

		err = s.treatmentDetailRepo.Save(tx, treatmentEntity)
		if err != nil {
			return "", err
		}
	}

	// save recipe
	for _, recipe := range request.Recipes {
		recipeEntity := mapper.MapToRecipeEntity(&recipe)
		recipeEntity.MedicalRecordID = medicalRecordID

		err = s.recipeRepo.Save(tx, recipeEntity)
		if err != nil {
			return "", err
		}
	}

	// save history
	historyEntity := mapper.MapToHistoryEntity(medicalRecordID, statusID)
	err = s.historyRepo.Save(tx, historyEntity)
	if err != nil {
		return "", err
	}

	return medicalRecordID, nil
}

func (s *Service) Delete(id string) error {
	// delete history
	if err := s.historyRepo.Delete(id); err != nil {
		return err
	}

	// delete recipe
	if err := s.recipeRepo.DeleteAll(id); err != nil {
		return err
	}

	// delete treatment
	if err := s.treatmentDetailRepo.DeleteAll(id); err != nil {
		return err
	}

	// delete nurse assignment
	if err := s.nurseAssignmentRepo.DeleteAll(id); err != nil {
		return err
	}

	// delete medical record
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
