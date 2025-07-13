package medical_record

import repository "medical-record-api/internal/medical_record/repository"

type RecipeService struct {
	repo *repository.RecipeRepository
}

func NewRecipeService(repo *repository.RecipeRepository) *RecipeService {
	return &RecipeService{repo: repo}
}

func (s *RecipeService) Delete(medicalRecordID, medicineID string) error {
	return s.repo.Delete(medicalRecordID, medicineID)
}
