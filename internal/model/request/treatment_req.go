package request

type TreatmentRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required,max=255"`
}
