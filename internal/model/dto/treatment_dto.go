package dto

type TreatmentDTO struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}
