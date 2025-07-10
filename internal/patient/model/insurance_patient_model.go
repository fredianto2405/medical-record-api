package patient

import "time"

type InsurancePatientEntity struct {
	ID              string     `db:"id"`
	PatientID       string     `db:"patient_id"`
	InsuranceID     string     `db:"insurance_id"`
	InsuranceNumber string     `db:"insurance_number"`
	CreatedAt       *time.Time `db:"created_at"`
	DeletedAt       *time.Time `db:"deleted_at"`
}

type InsurancePatientRequest struct {
	PatientID       string `json:"patient_id" validate:"required,max=255"`
	InsuranceID     string `json:"insurance_id" validate:"required,max=255"`
	InsuranceNumber string `json:"insurance_number" validate:"required,max=255"`
}
