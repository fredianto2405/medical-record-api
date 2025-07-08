package doctor

func MapToEntity(request *Request) *Entity {
	return &Entity{
		Email:              request.Email,
		Name:               request.Name,
		Gender:             request.Gender,
		SpecializationID:   request.SpecializationID,
		Phone:              request.Phone,
		Address:            request.Address,
		RegistrationNumber: request.RegistrationNumber,
		CreatedAt:          nil,
		DeletedAt:          nil,
	}
}

func MapToDTO(entity *Entity) *DTO {
	return &DTO{
		ID:                 entity.ID,
		Email:              entity.Email,
		Name:               entity.Name,
		Gender:             entity.Gender,
		SpecializationID:   entity.SpecializationID,
		Phone:              entity.Phone,
		Address:            entity.Address,
		RegistrationNumber: entity.RegistrationNumber,
		SharingFee:         entity.SharingFee,
	}
}
