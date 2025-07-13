package payment

import "time"

type MethodEntity struct {
	ID          string     `db:"id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	CreatedAt   *time.Time `db:"created_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

type MethodRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=255"`
}

type MethodDTO struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type StatusDTO struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	SortNumber  int    `json:"sort_number" db:"sort_number"`
}
