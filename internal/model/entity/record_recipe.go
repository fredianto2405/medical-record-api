package entity

import "time"

type RecordRecipe struct {
	ID              string    `db:"id"`
	MedicalRecordID string    `db:"medical_record_id"`
	MedicineID      string    `db:"medicine_id"`
	Price           int       `db:"price"`
	Quantity        int       `db:"quantity"`
	Dosage          string    `db:"dosage"`
	Instruction     string    `db:"instruction"`
	CreatedAt       time.Time `db:"created_at"`
}
