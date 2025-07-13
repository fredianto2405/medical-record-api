package nurse

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
	insertQuery := `insert into emr_nurse.nurses(name, 
			gender, 
			address, 
			phone, 
			registration_number, 
			sharing_fee)
		values(:name, :gender, :address, :phone, :registration_number, :sharing_fee)
		returning id, 
			name, 
			gender, 
			address, 
			phone, 
			registration_number, 
			sharing_fee`
	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Gender,
			&result.Address,
			&result.Phone,
			&result.RegistrationNumber,
			&result.SharingFee)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) Update(e *Entity) (*Entity, error) {
	updateQuery := `update emr_nurse.nurses
		set name = :name,
			gender = :gender,
			address = :address,
			phone = :phone,
			registration_number = :registration_number,
			sharing_fee = :sharing_fee
		where id = :id
		returning id, 
			name, 
			gender, 
			address, 
			phone, 
			registration_number, 
			sharing_fee`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Gender,
			&result.Address,
			&result.Phone,
			&result.RegistrationNumber,
			&result.SharingFee)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) FindAll(search string) ([]*DTO, error) {
	var nurses []*DTO

	searchPattern := "%" + search + "%"

	dataQuery := `select id, 
			name,
			gender,
			address,
			phone,
			registration_number,
			sharing_fee
		from emr_nurse.nurses
		where deleted_at isnull
		order by name asc`
	err := r.db.Select(&nurses, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return nurses, nil
}

func (r *Repository) FindAllPaginated(page, limit int, search string) ([]*DTO, int, error) {
	var nurses []*DTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where deleted_at isnull and name ilike $1`

	countQuery := `select count(0) from emr_nurse.nurses ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select id, 
			name,
			gender,
			address,
			phone,
			registration_number,
			sharing_fee
		from emr_nurse.nurses
		` + baseQuery + `
		order by name asc
		limit $2 offset $3`
	err = r.db.Select(&nurses, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return nurses, total, err
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `update emr_nurse.nurses set deleted_at = now() where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
