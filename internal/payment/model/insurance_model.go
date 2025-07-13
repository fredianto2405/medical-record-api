package payment

import "time"

type InsuranceEntity struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	Code      string     `db:"code"`
	Contact   string     `db:"contact"`
	Email     string     `db:"email"`
	IsActive  bool       `db:"is_active"`
	CreatedAt *time.Time `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type InsuranceRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
	Code     string `json:"code" validate:"required,max=255"`
	Contact  string `json:"contact" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	IsActive bool   `json:"is_active" validate:"required"`
}

type InsuranceDTO struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Code     string `json:"code" db:"code"`
	Contact  string `json:"contact" db:"contact"`
	Email    string `json:"email" db:"email"`
	IsActive bool   `json:"is_active" db:"is_active"`
}
