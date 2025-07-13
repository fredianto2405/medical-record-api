package medical_record

import model "medical-record-api/internal/medical_record/model"

func MapToEntity(request *model.Request) *model.Entity {
	return &model.Entity{
		PatientID:       request.PatientID,
		DoctorID:        request.DoctorID,
		Diagnosis:       request.Diagnosis,
		Notes:           request.Notes,
		PaymentMethodID: request.PaymentMethodID,
		PaymentStatusID: request.PaymentStatusID,
		InsuranceID:     request.InsuranceID,
		Anamnesis:       request.Anamnesis,
		CreatedAt:       nil,
	}
}
