package medicine

import medicine "medical-record-api/internal/medicine/model"

func MapToUnitEntity(request *medicine.UnitRequest) *medicine.UnitEntity {
	return &medicine.UnitEntity{
		Name:      request.Name,
		CreatedAt: nil,
		DeletedAt: nil,
	}
}

func MapToUnitDTO(entity *medicine.UnitEntity) *medicine.UnitDTO {
	return &medicine.UnitDTO{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
