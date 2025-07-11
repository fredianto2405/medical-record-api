package treatment

func MapToEntity(request *Request) *Entity {
	return &Entity{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		CreatedAt:   nil,
		DeletedAt:   nil,
	}
}

func MapToDTO(entity *Entity) *DTO {
	return &DTO{
		ID:          entity.ID,
		Name:        entity.Name,
		Price:       entity.Price,
		Description: entity.Description,
	}
}
