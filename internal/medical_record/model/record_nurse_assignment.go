package medical_record

import "time"

type NurseAssignmentEntity struct {
	ID              string     `db:"id"`
	MedicalRecordID string     `db:"medical_record_id"`
	NurseID         string     `db:"nurse_id"`
	CreatedAt       *time.Time `db:"created_at"`
}
