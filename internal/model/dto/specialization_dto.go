package dto

type SpecializationDTO struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
