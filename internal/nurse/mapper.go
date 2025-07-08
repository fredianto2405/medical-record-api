package nurse

func MapToEntity(request *Request) *Entity {
	return &Entity{
		Name:               request.Name,
		Gender:             request.Gender,
		Address:            request.Address,
		Phone:              request.Phone,
		RegistrationNumber: request.RegistrationNumber,
		SharingFee:         request.SharingFee,
		CreatedAt:          nil,
		DeletedAt:          nil,
	}
}

func MapToDTO(entity *Entity) *DTO {
	return &DTO{
		ID:                 entity.ID,
		Name:               entity.Name,
		Gender:             entity.Gender,
		Phone:              entity.Phone,
		Address:            entity.Address,
		RegistrationNumber: entity.RegistrationNumber,
		SharingFee:         entity.SharingFee,
	}
}
