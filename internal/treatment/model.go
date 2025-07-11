package treatment

import "time"

type Entity struct {
	ID          string     `db:"id"`
	Name        string     `db:"name"`
	Price       int        `db:"price"`
	Description string     `db:"description"`
	CreatedAt   *time.Time `db:"created_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

type Request struct {
	Name        string `json:"name" validate:"required,max=255"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required,max=255"`
}

type DTO struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}
