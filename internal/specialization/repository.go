package specialization

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e *Entity) (*Entity, error) {
	insertQuery := `insert into emr_doctor.specializations(name) values(:name) returning id, name`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) Update(e *Entity) (*Entity, error) {
	updateQuery := `update emr_doctor.specializations set name = :name where id = :id returning id, name`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) FindAll(search string) ([]*DTO, error) {
	var specializations []*DTO

	searchPattern := "%" + search + "%"

	dataQuery := `select id, 
    		name 
		from emr_doctor.specializations 
		where deleted_at isnull 
		and name ilike $1 order by name asc`
	err := r.db.Select(&specializations, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return specializations, nil
}

func (r *Repository) FindAllPaginated(page, limit int, search string) ([]*DTO, int, error) {
	var specializations []*DTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where deleted_at isnull and name ilike $1`

	countQuery := `select count(0) from emr_doctor.specializations ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select id, 
    		name 
		from emr_doctor.specializations 
		` + baseQuery + `
		order by name asc 
		limit $2 offset $3`
	err = r.db.Select(&specializations, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return specializations, total, err
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `update emr_doctor.specializations set deleted_at = now() where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
