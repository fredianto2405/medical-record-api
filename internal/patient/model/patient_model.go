package patient

import "time"

type Entity struct {
	ID                  string     `db:"id"`
	MedicalRecordNumber string     `db:"medical_record_number"`
	Name                string     `db:"name"`
	Gender              string     `db:"gender"`
	BirthDate           string     `db:"birth_date"`
	BloodType           string     `db:"blood_type"`
	Address             string     `db:"address"`
	Phone               string     `db:"phone"`
	Email               string     `db:"email"`
	HistoryOfIllness    string     `db:"history_of_illness"`
	Allergies           string     `db:"allergies"`
	IdentityType        string     `db:"identity_type"`
	IdentityNumber      string     `db:"identity_number"`
	CreatedAt           *time.Time `db:"created_at"`
	DeletedAt           *time.Time `db:"deleted_at"`
}

type Request struct {
	Name             string `json:"name" validate:"required,min=3,max=255"`
	Gender           string `json:"gender" validate:"required,max=255"`
	BirthDate        string `json:"birth_date" validate:"required"`
	BloodType        string `json:"blood_type" validate:"required"`
	Address          string `json:"address" validate:"required,max=255"`
	Phone            string `json:"phone" validate:"required,numeric,min=10,max=20"`
	Email            string `json:"email" validate:"required,email,max=255"`
	HistoryOfIllness string `json:"history_of_illness"`
	Allergies        string `json:"allergies"`
	IdentityType     string `json:"identity_type" validate:"required,max=255"`
	IdentityNumber   string `json:"identity_number" validate:"required,max=255"`
}

type DTO struct {
	ID                       string `json:"id" db:"id"`
	Name                     string `json:"name" db:"name"`
	Gender                   string `json:"gender" db:"gender"`
	BirthDate                string `json:"birth_date" db:"birth_date"`
	BloodType                string `json:"blood_type" db:"blood_type"`
	Address                  string `json:"address" db:"address"`
	Phone                    string `json:"phone" db:"phone"`
	Email                    string `json:"email" db:"email"`
	HistoryOfIllness         string `json:"history_of_illness" db:"history_of_illness"`
	Allergies                string `json:"allergies" db:"allergies"`
	IdentityType             string `json:"identity_type" db:"identity_type"`
	IdentityNumber           string `json:"identity_number" db:"identity_number"`
	EmergencyContactName     string `json:"emergency_contact_name,omitempty" db:"emergency_contact_name"`
	EmergencyContactPhone    string `json:"emergency_contact_phone,omitempty" db:"emergency_contact_phone"`
	EmergencyContactRelation string `json:"emergency_contact_relation,omitempty" db:"emergency_contact_relation"`
	InsuranceID              string `json:"insurance_id,omitempty" db:"insurance_id"`
	InsuranceNumber          string `json:"insurance_number,omitempty" db:"insurance_number"`
	MedicalRecordNumber      string `json:"medical_record_number" db:"medical_record_number"`
}
