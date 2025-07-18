package medical_record

import "time"

type TreatmentDetailEntity struct {
	ID              string     `db:"id"`
	MedicalRecordID string     `db:"medical_record_id"`
	TreatmentID     string     `db:"treatment_id"`
	Price           int        `db:"price"`
	CreatedAt       *time.Time `db:"created_at"`
}

type TreatmentRequest struct {
	TreatmentID string `json:"treatment_id" validate:"required,max=255"`
	Price       int    `json:"price" validate:"required"`
}

type TreatmentDTO struct {
	TreatmentID   string `json:"treatment_id" db:"treatment_id"`
	TreatmentName string `json:"treatment_name" db:"treatment_name"`
	Price         int    `json:"price" db:"price"`
}
