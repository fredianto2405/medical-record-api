package medical_record

import model "medical-record-api/internal/medical_record/model"

func MapToRecipeEntity(request *model.RecipeRequest) *model.RecipeEntity {
	return &model.RecipeEntity{
		MedicineID:  request.MedicineID,
		Price:       request.Price,
		Quantity:    request.Quantity,
		Dosage:      request.Dosage,
		Instruction: request.Instruction,
	}
}
