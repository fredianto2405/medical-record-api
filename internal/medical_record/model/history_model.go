package medical_record

import "time"

type HistoryEntity struct {
	ID              string     `db:"id"`
	MedicalRecordID string     `db:"medical_record_id"`
	StatusID        int        `db:"status_id"`
	CreatedAt       *time.Time `db:"created_at"`
}

type HistoryDTO struct {
	StatusID   int    `json:"status_id" db:"status_id"`
	StatusName string `json:"status_name" db:"status_name"`
	Timestamp  string `json:"timestamp" db:"timestamp"`
}
