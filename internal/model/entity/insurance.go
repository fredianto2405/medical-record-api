package entity

import "time"

type Insurance struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Code      string    `db:"code"`
	Contact   string    `db:"contact"`
	Email     string    `db:"email"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
