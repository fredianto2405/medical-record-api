package user

import "time"

type Entity struct {
	ID        string     `db:"id"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	IsActive  bool       `db:"is_active"`
	RoleID    int        `db:"role_id"`
	CreatedAt *time.Time `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}
