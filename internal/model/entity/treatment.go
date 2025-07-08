package entity

import "time"

type Treatment struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Price       int       `db:"price"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}
