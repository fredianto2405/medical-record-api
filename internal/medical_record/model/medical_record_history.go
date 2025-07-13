package medical_record

import "time"

type HistoryEntity struct {
	ID              string     `db:"id"`
	MedicalRecordID string     `db:"medical_record_id"`
	StatusID        string     `db:"status_id"`
	CreatedAt       *time.Time `db:"created_at"`
}
