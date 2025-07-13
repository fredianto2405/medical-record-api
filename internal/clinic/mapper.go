package clinic

func MapToEntity(request *Request) *Entity {
	return &Entity{
		Name:                  request.Name,
		Address:               request.Address,
		Logo:                  request.Logo,
		SharingFeeType:        request.SharingFeeType,
		PatientMedicationCost: request.PatientMedicationCost,
		NurseSharingFee:       request.NurseSharingFee,
	}
}

func MapToDTO(entity *Entity) *DTO {
	return &DTO{
		ID:                    entity.ID,
		Name:                  entity.Name,
		Address:               entity.Address,
		Logo:                  entity.Logo,
		SharingFeeType:        entity.SharingFeeType,
		PatientMedicationCost: entity.PatientMedicationCost,
		NurseSharingFee:       entity.NurseSharingFee,
	}
}
