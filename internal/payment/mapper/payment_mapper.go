package payment

import payment "medical-record-api/internal/payment/model"

func MapToMethodEntity(request *payment.MethodRequest) *payment.MethodEntity {
	return &payment.MethodEntity{
		Name:        request.Name,
		Description: request.Description,
		CreatedAt:   nil,
		DeletedAt:   nil,
	}
}

func MapToMethodDTO(entity *payment.MethodEntity) *payment.MethodDTO {
	return &payment.MethodDTO{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
	}
}
