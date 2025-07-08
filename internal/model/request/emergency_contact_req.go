package request

type EmergencyContactRequest struct {
	PatientID string `json:"patient_id" validate:"required,max=255"`
	Name      string `json:"name" validate:"required,min=3,max=255"`
	Phone     string `json:"phone" validate:"required,numeric,min=10,max=20"`
	Relation  string `json:"relation" validate:"required,max=255"`
}
