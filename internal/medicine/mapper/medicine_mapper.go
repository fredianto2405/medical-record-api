package medicine

import medicine "medical-record-api/internal/medicine/model"

func MapToEntity(request *medicine.Request) *medicine.Entity {
	return &medicine.Entity{
		Code:       request.Code,
		Name:       request.Name,
		CategoryID: request.CategoryID,
		UnitID:     request.UnitID,
		Price:      request.Price,
		Stock:      request.Stock,
		ExpiryDate: request.ExpiryDate,
		Dosage:     request.Dosage,
		CreatedAt:  nil,
		DeletedAt:  nil,
	}
}

func MapToDTO(entity *medicine.Entity) *medicine.DTO {
	return &medicine.DTO{
		ID:         entity.ID,
		Code:       entity.Code,
		Name:       entity.Name,
		CategoryID: entity.CategoryID,
		UnitID:     entity.UnitID,
		Price:      entity.Price,
		Stock:      entity.Stock,
		ExpiryDate: entity.ExpiryDate,
		Dosage:     entity.Dosage,
	}
}
