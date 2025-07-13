package payment

import payment "medical-record-api/internal/payment/model"

func MapToInsuranceEntity(request *payment.InsuranceRequest) *payment.InsuranceEntity {
	return &payment.InsuranceEntity{
		Name:      request.Name,
		Code:      request.Code,
		Contact:   request.Contact,
		Email:     request.Email,
		IsActive:  request.IsActive,
		CreatedAt: nil,
		DeletedAt: nil,
	}
}

func MapToInsuranceDTO(entity *payment.InsuranceEntity) *payment.InsuranceDTO {
	return &payment.InsuranceDTO{
		ID:       entity.ID,
		Name:     entity.Name,
		Code:     entity.Code,
		Contact:  entity.Contact,
		Email:    entity.Email,
		IsActive: entity.IsActive,
	}
}
