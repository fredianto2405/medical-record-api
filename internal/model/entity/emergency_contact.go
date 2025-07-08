package entity

import "time"

type EmergencyContact struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Relation  string    `db:"relation"`
	PatientID string    `db:"patient_id"`
	CreatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
