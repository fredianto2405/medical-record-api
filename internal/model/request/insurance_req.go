package request

type InsuranceRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
	Code     string `json:"code" validate:"required,max=255"`
	Contact  string `json:"contact" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	IsActive bool   `json:"is_active" validate:"required"`
}
