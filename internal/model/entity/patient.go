package entity

import "time"

type Patient struct {
	ID                  string    `db:"id"`
	MedicalRecordNumber string    `db:"medical_record_number"`
	Name                string    `db:"name"`
	Gender              string    `db:"gender"`
	BirthDate           time.Time `db:"birth_date"`
	BloodType           string    `db:"blood_type"`
	Address             string    `db:"address"`
	Phone               string    `db:"phone"`
	Email               string    `db:"email"`
	HistoryOfIllness    string    `db:"history_of_illness"`
	Allergies           string    `db:"allergies"`
	IdentityType        string    `db:"identity_type"`
	IdentityNumber      string    `db:"identity_number"`
	CreatedAt           time.Time `db:"created_at"`
	DeletedAt           time.Time `db:"deleted_at"`
}
