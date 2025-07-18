package medical_record

import (
	"github.com/jmoiron/sqlx"
	mapper "medical-record-api/internal/medical_record/mapper"
	model "medical-record-api/internal/medical_record/model"
	repository "medical-record-api/internal/medical_record/repository"
)

type RecipeService struct {
	db   *sqlx.DB
	repo *repository.RecipeRepository
}

func NewRecipeService(db *sqlx.DB, repo *repository.RecipeRepository) *RecipeService {
	return &RecipeService{
		db:   db,
		repo: repo,
	}
}

func (s *RecipeService) Delete(medicalRecordID, medicineID string) error {
	return s.repo.Delete(medicalRecordID, medicineID)
}

func (s *RecipeService) Create(medicalRecordID string, request model.RecipeRequest) error {
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

	recipeEntity := mapper.MapToRecipeEntity(&request)
	recipeEntity.MedicalRecordID = medicalRecordID
	err = s.repo.Save(tx, recipeEntity)

	if err != nil {
		return err
	}

	return nil
}
