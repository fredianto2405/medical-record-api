package medical_record

import "time"

type RecipeEntity struct {
	ID              string     `db:"id"`
	MedicalRecordID string     `db:"medical_record_id"`
	MedicineID      string     `db:"medicine_id"`
	Price           int        `db:"price"`
	Quantity        int        `db:"quantity"`
	Dosage          string     `db:"dosage"`
	Instruction     string     `db:"instruction"`
	CreatedAt       *time.Time `db:"created_at"`
}

type RecipeRequest struct {
	MedicineID  string `json:"medicine_id" validate:"required,max=255"`
	Price       int    `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	Dosage      string `json:"dosage" validate:"required,max=255"`
	Instruction string `json:"instruction" validate:"required,max=255"`
}

type RecipeDTO struct {
	MedicineID   string `json:"medicine_id" db:"medicine_id"`
	MedicineName string `json:"medicine_name" db:"medicine_name"`
	CategoryName string `json:"category_name" db:"category_name"`
	UnitName     string `json:"unit_name" db:"unit_name"`
	Price        int    `json:"price" db:"price"`
	Quantity     int    `json:"quantity" db:"quantity"`
	Dosage       string `json:"dosage" db:"dosage"`
	Instruction  string `json:"instruction" db:"instruction"`
}
