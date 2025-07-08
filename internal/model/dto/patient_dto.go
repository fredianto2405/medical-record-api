package dto

type PatientDTO struct {
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
