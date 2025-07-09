package patient

import patient "medical-record-api/internal/patient/model"

func MapToEntity(request *patient.Request) *patient.Entity {
	return &patient.Entity{
		Name:             request.Name,
		Gender:           request.Gender,
		BirthDate:        request.BirthDate,
		BloodType:        request.BloodType,
		Address:          request.Address,
		Phone:            request.Phone,
		Email:            request.Email,
		HistoryOfIllness: request.HistoryOfIllness,
		Allergies:        request.Allergies,
		IdentityType:     request.IdentityType,
		IdentityNumber:   request.IdentityNumber,
		CreatedAt:        nil,
		DeletedAt:        nil,
	}
}

func MapToDTO(entity *patient.Entity) *patient.DTO {
	return &patient.DTO{
		ID:                  entity.ID,
		Name:                entity.Name,
		Gender:              entity.Gender,
		BirthDate:           entity.BirthDate,
		BloodType:           entity.BloodType,
		Address:             entity.Address,
		Phone:               entity.Phone,
		Email:               entity.Email,
		HistoryOfIllness:    entity.HistoryOfIllness,
		Allergies:           entity.Allergies,
		IdentityType:        entity.IdentityType,
		IdentityNumber:      entity.IdentityNumber,
		MedicalRecordNumber: entity.MedicalRecordNumber,
	}
}
