package medical_record

import "time"

type TreatmentDetailEntity struct {
	ID              string     `db:"id"`
	MedicalRecordID string     `db:"medical_record_id"`
	TreatmentID     string     `db:"treatment_id"`
	Price           int        `db:"price"`
	CreatedAt       *time.Time `db:"created_at"`
}
