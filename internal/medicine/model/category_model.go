package medicine

import "time"

type CategoryEntity struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type CategoryRequest struct {
	Name string `json:"name" validate:"required,max=255"`
}

type CategoryDTO struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
