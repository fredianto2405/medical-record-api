package medical_record

import "time"

type Entity struct {
	ID              string     `db:"id"`
	PatientID       string     `db:"patient_id"`
	DoctorID        string     `db:"doctor_id"`
	Diagnosis       string     `db:"diagnosis"`
	Notes           string     `db:"notes"`
	StatusID        int        `db:"status_id"`
	PaymentMethodID string     `db:"payment_method_id"`
	PaymentStatusID int        `db:"payment_status_id"`
	InsuranceID     string     `db:"insurance_id"`
	Anamnesis       string     `db:"anamnesis"`
	CreatedAt       *time.Time `db:"created_at"`
}

type Request struct {
	PatientID       string             `json:"patient_id" validate:"required,max=255"`
	DoctorID        string             `json:"doctor_id" validate:"required,max=255"`
	Diagnosis       string             `json:"diagnosis" validate:"required,max=500"`
	Notes           string             `json:"notes" validate:"required,max=500"`
	PaymentMethodID string             `json:"payment_method_id" validate:"required,max=255"`
	PaymentStatusID int                `json:"payment_status_id" validate:"required"`
	InsuranceID     string             `json:"insurance_id" validate:"required,max=255"`
	Anamnesis       string             `json:"anamnesis" validate:"required,max=500"`
	NurseIDs        []string           `json:"nurse_ids" validate:"required,min=1,dive"`
	Treatments      []TreatmentRequest `json:"treatments" validate:"required,min=1,dive"`
	Recipes         []RecipeRequest    `json:"recipes" validate:"required,min=1,dive"`
}

type DTO struct {
	MedicalRecordID string `json:"medical_record_id"`
}
