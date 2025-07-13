package medical_record

import model "medical-record-api/internal/medical_record/model"

func MapToHistoryEntity(medicalRecordID string, statusID int) *model.HistoryEntity {
	return &model.HistoryEntity{
		MedicalRecordID: medicalRecordID,
		StatusID:        statusID,
		CreatedAt:       nil,
	}
}
