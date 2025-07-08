package request

type InsurancePatientRequest struct {
	PatientID       string `json:"patient_id" validate:"required,max=255"`
	InsuranceID     string `json:"insurance_id" validate:"required,max=255"`
	InsuranceNumber string `json:"insurance_number" validate:"required,max=255"`
}
