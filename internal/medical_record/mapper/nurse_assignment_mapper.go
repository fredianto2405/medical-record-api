package medical_record

import (
	model "medical-record-api/internal/medical_record/model"
)

func MapToNurseAssignmentEntity(medicalRecordID, nurseID string) *model.NurseAssignmentEntity {
	return &model.NurseAssignmentEntity{
		MedicalRecordID: medicalRecordID,
		NurseID:         nurseID,
		CreatedAt:       nil,
	}
}
