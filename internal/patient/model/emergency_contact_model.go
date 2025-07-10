package patient

import "time"

type EmergencyContactEntity struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	Phone     string     `db:"phone"`
	Relation  string     `db:"relation"`
	PatientID string     `db:"patient_id"`
	CreatedAt *time.Time `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type EmergencyContactRequest struct {
	PatientID string `json:"patient_id" validate:"required,max=255"`
	Name      string `json:"name" validate:"required,min=3,max=255"`
	Phone     string `json:"phone" validate:"required,numeric,min=10,max=20"`
	Relation  string `json:"relation" validate:"required,max=255"`
}
