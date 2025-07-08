package medicine

import "time"

type Entity struct {
	ID         string     `db:"id"`
	Code       string     `db:"code"`
	Name       string     `db:"name"`
	CategoryID string     `db:"category_id"`
	UnitID     string     `db:"unit_id"`
	Price      int        `db:"price"`
	Stock      int        `db:"stock"`
	ExpiryDate string     `db:"expiry_date"`
	Dosage     string     `db:"dosage"`
	CreatedAt  *time.Time `db:"created_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}

type Request struct {
	Code       string `json:"code" validate:"required,max=255"`
	Name       string `json:"name" validate:"required,max=255"`
	CategoryID string `json:"category_id" validate:"required,max=255"`
	UnitID     string `json:"unit_id" validate:"required,max=255"`
	Price      int    `json:"price" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	ExpiryDate string `json:"expiry_date" validate:"required,max=255"`
	Dosage     string `json:"dosage" validate:"required,max=255"`
}

type DTO struct {
	ID           string `json:"id" db:"id"`
	Code         string `json:"code" db:"code"`
	Name         string `json:"name" db:"name"`
	CategoryID   string `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name,omitempty" db:"category_name"`
	UnitID       string `json:"unit_id" db:"unit_id"`
	UnitName     string `json:"unit_name,omitempty" db:"unit_name"`
	Price        int    `json:"price" db:"price"`
	Stock        int    `json:"stock" db:"stock"`
	ExpiryDate   string `json:"expiry_date" db:"expiry_date"`
	Dosage       string `json:"dosage" db:"dosage"`
}
