package user

import (
	"github.com/jmoiron/sqlx"
)

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

func (r *Repository) FindAllPaginated(page, limit int, search string) ([]*DTO, int, error) {
	var users []*DTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where u.deleted_at isnull and u.email ilike $1`

	countQuery := `select count(0) 
		from emr_auth.users u
		join emr_auth.roles r on r.id = u.role_id ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select u.id, 
			u.email,
			r."name" as role,
			u.is_active
		from emr_auth.users u
		join emr_auth.roles r on r.id = u.role_id
		` + baseQuery + `
		order by u.email asc
		limit $2 offset $3`
	err = r.db.Select(&users, dataQuery, searchPattern, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
