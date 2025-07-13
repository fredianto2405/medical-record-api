package payment

import (
	"github.com/jmoiron/sqlx"
	payment "medical-record-api/internal/payment/model"
)

type InsuranceRepository struct {
	db *sqlx.DB
}

func NewInsuranceRepository(db *sqlx.DB) *InsuranceRepository {
	return &InsuranceRepository{db: db}
}

func (r *InsuranceRepository) Save(e *payment.InsuranceEntity) (*payment.InsuranceEntity, error) {
	insertQuery := `insert into emr_payment.insurances(name, code, contact, email, is_active)
		values(:name, :code, :contact, :email, :is_active)
		returning id, name, code, contact, email, is_active`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result payment.InsuranceEntity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Code,
			&result.Contact,
			&result.Email,
			&result.IsActive)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *InsuranceRepository) Update(e *payment.InsuranceEntity) (*payment.InsuranceEntity, error) {
	updateQuery := `update emr_payment.insurances
		set name = :name,
			code = :code,
			contact = :contact,
			email = :email,
			is_active = :is_active
		where id = :id
		returning id, name, code, contact, email, is_active`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result payment.InsuranceEntity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Code,
			&result.Contact,
			&result.Email,
			&result.IsActive)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *InsuranceRepository) FindAll(search string) ([]*payment.InsuranceDTO, error) {
	var insurances []*payment.InsuranceDTO

	searchPattern := "%" + search + "%"

	dataQuery := `select id, name, code, contact, email, is_active
		from emr_payment.insurances
		where deleted_at isnull
		and name ilike $1
		order by name asc`
	err := r.db.Select(&insurances, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return insurances, nil
}

func (r *InsuranceRepository) FindAllPaginated(page, limit int, search string) ([]*payment.InsuranceDTO, int, error) {
	var insurances []*payment.InsuranceDTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where deleted_at isnull and name ilike $1`

	countQuery := `select count(0) from emr_payment.insurances ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select id, 
			name, 
			code, 
			contact, 
			email, 
			is_active
		from emr_payment.insurances
		` + baseQuery + `
		order by name asc
		limit $2 offset $3`
	err = r.db.Select(&insurances, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return insurances, total, err
}

func (r *InsuranceRepository) Delete(id string) error {
	deleteQuery := `update emr_payment.insurances
		set deleted_at = now()
		where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
