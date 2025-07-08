package specialization

func MapToEntity(request *Request) *Entity {
	return &Entity{
		Name:      request.Name,
		CreatedAt: nil,
		DeletedAt: nil,
	}
}

func MapToDTO(entity *Entity) *DTO {
	return &DTO{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
