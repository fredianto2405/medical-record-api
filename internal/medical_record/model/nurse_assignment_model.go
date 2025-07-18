package medical_record

import "time"

type NurseAssignmentEntity struct {
	ID              string     `db:"id"`
	MedicalRecordID string     `db:"medical_record_id"`
	NurseID         string     `db:"nurse_id"`
	CreatedAt       *time.Time `db:"created_at"`
}

type NurseAssignmentRequest struct {
	NurseID string `json:"nurse_id" validate:"required,max=255"`
}

type NurseDTO struct {
	NurseID   string `json:"nurse_id" db:"nurse_id"`
	NurseName string `json:"nurse_name" db:"nurse_name"`
}
