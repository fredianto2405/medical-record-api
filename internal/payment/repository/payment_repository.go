package payment

import (
	"github.com/jmoiron/sqlx"
	payment "medical-record-api/internal/payment/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e *payment.MethodEntity) (*payment.MethodEntity, error) {
	insertQuery := `insert into emr_payment.payment_methods(name, description)
		values(:name, :description)
		returning id, name, description`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result payment.MethodEntity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Description)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) Update(e *payment.MethodEntity) (*payment.MethodEntity, error) {
	updateQuery := `update emr_payment.payment_methods
		set name = :name,
			description = :description
		where id = :id
		returning id, name, description`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result payment.MethodEntity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Description)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) FindAll(search string) ([]*payment.MethodDTO, error) {
	var methods []*payment.MethodDTO

	searchPattern := "%" + search + "%"

	dataQuery := `select id, name, description
		from emr_payment.payment_methods
		where deleted_at isnull
		and name ilike $1
		order by name asc`
	err := r.db.Select(&methods, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return methods, nil
}

func (r *Repository) FindAllPaginated(page, limit int, search string) ([]*payment.MethodDTO, int, error) {
	var methods []*payment.MethodDTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where deleted_at isnull and name ilike $1`

	countQuery := `select count(0) from emr_payment.payment_methods ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select id, 
			name,
			description
		from emr_payment.payment_methods
		` + baseQuery + `
		order by name asc
		limit $2 offset $3`
	err = r.db.Select(&methods, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return methods, total, err
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `update emr_payment.payment_methods
		set deleted_at = now()
		where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}

func (r *Repository) FindAllStatus() ([]*payment.StatusDTO, error) {
	var statuses []*payment.StatusDTO

	dataQuery := `select id, name, description, sort_number
		from emr_payment.payment_statuses
		where deleted_at isnull
		order by sort_number asc`
	err := r.db.Select(&statuses, dataQuery)

	return statuses, err
}
