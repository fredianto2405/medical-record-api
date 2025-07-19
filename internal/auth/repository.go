package auth

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindUserByEmail(email string) (*UserDTO, error) {
	var user UserDTO

	dataQuery := `select u.id, 
		   u.email, 
		   u.password, 
		   r.name role 
		from emr_auth.users u 
		join emr_auth.roles r on r.id = u.role_id 
		where email = $1 
		and deleted_at isnull`
	err := r.db.Get(&user, dataQuery, email)

	return &user, err
}
