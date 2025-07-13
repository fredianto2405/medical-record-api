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
