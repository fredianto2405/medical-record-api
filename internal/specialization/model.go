package specialization

import "time"

type Entity struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type Request struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type DTO struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
