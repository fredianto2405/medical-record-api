package medicine

import (
	medicine "medical-record-api/internal/medicine/model"
)

func MapToCategoryEntity(request *medicine.CategoryRequest) *medicine.CategoryEntity {
	return &medicine.CategoryEntity{
		Name:      request.Name,
		CreatedAt: nil,
		DeletedAt: nil,
	}
}

func MapToCategoryDTO(entity *medicine.CategoryEntity) *medicine.CategoryDTO {
	return &medicine.CategoryDTO{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
