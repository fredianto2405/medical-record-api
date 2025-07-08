package request

type PatientRequest struct {
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
