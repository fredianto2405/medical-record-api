package medical_record

import model "medical-record-api/internal/medical_record/model"

func MapToTreatmentDetailEntity(request *model.TreatmentRequest) *model.TreatmentDetailEntity {
	return &model.TreatmentDetailEntity{
		TreatmentID: request.TreatmentID,
		Price:       request.Price,
		CreatedAt:   nil,
	}
}
