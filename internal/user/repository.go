package user

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e *Entity) error {
	insertQuery := `insert into emr_auth.users(email, password, is_active, role_id)
		values(:email, :password, :is_active, :role_id)`
	_, err := r.db.NamedExec(insertQuery, e)
	return err
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `update emr_auth.users set deleted_at = now() where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}

func (r *Repository) EmailExists(email string) (bool, error) {
	var exists bool
	dataQuery := `select exists (select 1 from emr_auth.users where email = $1 and deleted_at isnull)`

	err := r.db.Get(&exists, dataQuery, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}
