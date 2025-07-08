package request

type SpecializationRequest struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}
