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

type UpdateRequest struct {
	DoctorID        string `json:"doctor_id" validate:"required,max=255"`
	Diagnosis       string `json:"diagnosis" validate:"required,max=500"`
	Notes           string `json:"notes" validate:"required,max=500"`
	PaymentMethodID string `json:"payment_method_id" validate:"required,max=255"`
	PaymentStatusID int    `json:"payment_status_id" validate:"required"`
	InsuranceID     string `json:"insurance_id" validate:"required,max=255"`
	Anamnesis       string `json:"anamnesis" validate:"required,max=500"`
}

type UpdateStatusRequest struct {
	StatusID int `json:"status_id" validate:"required"`
}

type SavedDTO struct {
	MedicalRecordID string `json:"medical_record_id"`
}

type DTO struct {
	ID                   string          `json:"id" db:"id"`
	CreatedAt            string          `json:"created_at" db:"created_at"`
	PatientID            string          `json:"patient_id" db:"patient_id"`
	PatientName          string          `json:"patient_name" db:"patient_name"`
	PatientGender        string          `json:"patient_gender" db:"patient_gender"`
	PatientAddress       string          `json:"patient_address" db:"patient_address"`
	PatientBirthDate     string          `json:"patient_birth_date" db:"patient_birth_date"`
	MedicalRecordNumber  string          `json:"medical_record_number" db:"medical_record_number"`
	DoctorID             string          `json:"doctor_id" db:"doctor_id"`
	DoctorName           string          `json:"doctor_name" db:"doctor_name"`
	DoctorSpecialization string          `json:"doctor_specialization" db:"doctor_specialization"`
	Diagnosis            string          `json:"diagnosis" db:"diagnosis"`
	Notes                string          `json:"notes" db:"notes"`
	StatusID             int             `json:"status_id" db:"status_id"`
	StatusName           string          `json:"status_name" db:"status_name"`
	PaymentMethodID      string          `json:"payment_method_id" db:"payment_method_id"`
	PaymentMethodName    string          `json:"payment_method_name" db:"payment_method_name"`
	PaymentStatusID      int             `json:"payment_status_id" db:"payment_status_id"`
	PaymentStatusName    string          `json:"payment_status_name" db:"payment_status_name"`
	InsuranceID          string          `json:"insurance_id" db:"insurance_id"`
	InsuranceName        string          `json:"insurance_name" db:"insurance_name"`
	InsuranceNumber      string          `json:"insurance_number" db:"insurance_number"`
	Anamnesis            string          `json:"anamnesis" db:"anamnesis"`
	Histories            []*HistoryDTO   `json:"histories"`
	Nurses               []*NurseDTO     `json:"nurses"`
	Treatments           []*TreatmentDTO `json:"treatments"`
	Recipes              []*RecipeDTO    `json:"recipes"`
}
