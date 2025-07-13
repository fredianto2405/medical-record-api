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
