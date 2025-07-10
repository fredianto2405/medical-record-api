package patient

import patient "medical-record-api/internal/patient/model"

func MapToInsurancePatientEntity(request *patient.InsurancePatientRequest) *patient.InsurancePatientEntity {
	return &patient.InsurancePatientEntity{
		PatientID:       request.PatientID,
		InsuranceID:     request.InsuranceID,
		InsuranceNumber: request.InsuranceNumber,
		CreatedAt:       nil,
		DeletedAt:       nil,
	}
}
