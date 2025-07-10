package patient

import patient "medical-record-api/internal/patient/model"

func MapToEmergencyContactEntity(request *patient.EmergencyContactRequest) *patient.EmergencyContactEntity {
	return &patient.EmergencyContactEntity{
		Name:      request.Name,
		Phone:     request.Phone,
		Relation:  request.Relation,
		PatientID: request.PatientID,
		CreatedAt: nil,
		DeletedAt: nil,
	}
}
