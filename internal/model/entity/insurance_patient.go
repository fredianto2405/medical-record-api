package entity

import "time"

type InsurancePatient struct {
	ID              string    `db:"id"`
	InsuranceID     string    `db:"insurance_id"`
	InsuranceNumber string    `db:"insurance_number"`
	CreatedAt       time.Time `db:"created_at"`
	DeletedAt       time.Time `db:"deleted_at"`
}
