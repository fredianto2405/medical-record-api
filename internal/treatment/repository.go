package treatment

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e *Entity) (*Entity, error) {
	insertQuery := `insert into emr_treatment.treatments(name, price, description)
		values(:name, :price, :description)
		returning id, name, price, description`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.Price, &result.Description)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) Update(e *Entity) (*Entity, error) {
	updateQuery := `update emr_treatment.treatments
		set name = :name,
			price = :price,
			description = :description
		where id = :id
		returning id, name, price, description`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.Price, &result.Description)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) FindAll(search string) ([]*DTO, error) {
	var treatments []*DTO

	searchPattern := "%" + search + "%"

	dataQuery := `select id, name, price, description
		from emr_treatment.treatments
		where deleted_at isnull
		and name ilike $1
		order by name asc`
	err := r.db.Select(&treatments, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return treatments, nil
}

func (r *Repository) FindAllPaginated(page, limit int, search string) ([]*DTO, int, error) {
	var treatments []*DTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where deleted_at isnull and name ilike $1`

	countQuery := `select count(0) from emr_treatment.treatments ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select id, name, price, description
		from emr_treatment.treatments
		` + baseQuery + `
		order by name asc
		limit $2 offset $3`
	err = r.db.Select(&treatments, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return treatments, total, err
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `update emr_treatment.treatments
		set deleted_at = now()
		where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
