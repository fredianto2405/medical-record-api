package doctor

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
	insertQuery := `insert into emr_doctor.doctors(email, name, gender, 
			specialization_id, phone, address, 
			registration_number, sharing_fee)
		values(:email, :name, :gender, 
		       :specialization_id, :phone, :address, 
		       :registration_number, :sharing_fee)
		returning id, email, name, 
			gender, specialization_id, phone, 
			address, registration_number, sharing_fee`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Email, &result.Name,
			&result.Gender, &result.SpecializationID, &result.Phone,
			&result.Address, &result.RegistrationNumber, &result.SharingFee)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) Update(e *Entity) (*Entity, error) {
	updateQuery := `update emr_doctor.doctors
		set email = :email, name = :name, gender = :gender,
			specialization_id = :specialization_id, phone = :phone, address = :address,
			registration_number = :registration_number, sharing_fee = :sharing_fee
		where id = :id 
		returning id, email, name, 
			gender, specialization_id, phone, 
			address, registration_number, sharing_fee`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Email, &result.Name,
			&result.Gender, &result.SpecializationID, &result.Phone,
			&result.Address, &result.RegistrationNumber, &result.SharingFee)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) FindAll(search string) ([]*DTO, error) {
	var doctors []*DTO

	searchPattern := "%" + search + "%"

	dataQuery := `select d.id as id,
			d.email,
			d.name,
			d.gender,
			d.specialization_id,
			s.name as specialization_name,
			d.phone,
			d.address,
			d.registration_number,
			d.sharing_fee
		from emr_doctor.doctors d 
		join emr_doctor.specializations s on s.id = d.specialization_id
		where d.deleted_at isnull
		and d.name ilike $1
		order by d.name asc`
	err := r.db.Select(&doctors, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return doctors, nil
}

func (r *Repository) FindAllPaginated(page, limit int, search string) ([]*DTO, int, error) {
	var doctors []*DTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where d.deleted_at isnull and d.name ilike $1`

	countQuery := `select count(0) from emr_doctor.doctors d ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select d.id as id,
			d.email,
			d.name,
			d.gender,
			d.specialization_id,
			s.name as specialization_name,
			d.phone,
			d.address,
			d.registration_number,
			d.sharing_fee
		from emr_doctor.doctors d 
		join emr_doctor.specializations s on s.id = d.specialization_id
		` + baseQuery + `
		order by d.name asc
		limit $2 offset $3`
	err = r.db.Select(&doctors, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return doctors, total, err
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `update emr_doctor.doctors set deleted_at = now() where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
