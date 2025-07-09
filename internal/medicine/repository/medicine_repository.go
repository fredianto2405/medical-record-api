package medicine

import (
	"github.com/jmoiron/sqlx"
	medicine "medical-record-api/internal/medicine/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(entity *medicine.Entity) (*medicine.Entity, error) {
	insertQuery := `insert into emr_medicine.medicines(code, 
			name, 
			category_id,
			unit_id,
			price,
			stock,
			expiry_date,
			dosage)
		values(:code, 
			:name, 
			:category_id,
			:unit_id,
			:price,
			:stock,
			:expiry_date,
			:dosage)
		returning id, 
			code, 
			name, 
			category_id,
			unit_id,
			price,
			stock,
			expiry_date,
			dosage`

	rows, err := r.db.NamedQuery(insertQuery, entity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result medicine.Entity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Code,
			&result.Name,
			&result.CategoryID,
			&result.UnitID,
			&result.Price,
			&result.Stock,
			&result.ExpiryDate,
			&result.Dosage)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) Update(entity *medicine.Entity) (*medicine.Entity, error) {
	updateQuery := `update emr_medicine.medicines
		set code = :code,
			name = :name,
			category_id = :category_id,
			unit_id = :unit_id,
			price = :price,
			stock = :stock,
			expiry_date = :expiry_date,
			dosage = :dosage
		where id = :id
		returning id, 
			code, 
			name, 
			category_id,
			unit_id,
			price,
			stock,
			expiry_date,
			dosage`

	rows, err := r.db.NamedQuery(updateQuery, entity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result medicine.Entity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Code,
			&result.Name,
			&result.CategoryID,
			&result.UnitID,
			&result.Price,
			&result.Stock,
			&result.ExpiryDate,
			&result.Dosage)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) FindAll(search string) ([]*medicine.DTO, error) {
	var medicines []*medicine.DTO

	searchPattern := "%" + search + "%"

	dataQuery := `select m.id, 
			m.code,
			m.name,
			c.id category_id,
			c.name category_name,
			u.id unit_id,
			u.name unit_name,
			m.price,
			m.stock,
			to_char(m.expiry_date, 'YYYY-MM-DD') expiry_date,
			m.dosage 
		from emr_medicine.medicines m 
		inner join emr_medicine.categories c on c.id = m.category_id
		inner join emr_medicine.units u on u.id = m.unit_id 
		where m.deleted_at isnull and m.name ilike $1`

	err := r.db.Select(&medicines, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return medicines, nil
}

func (r *Repository) FindAllPaginated(page, limit int, search string) ([]*medicine.DTO, int, error) {
	var medicines []*medicine.DTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where m.deleted_at isnull and m.name ilike $1`

	countQuery := `select count(0) from emr_medicine.medicines m ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select m.id, 
			m.code,
			m.name,
			c.id category_id,
			c.name category_name,
			u.id unit_id,
			u.name unit_name,
			m.price,
			m.stock,
			to_char(m.expiry_date, 'YYYY-MM-DD') expiry_date,
			m.dosage 
		from emr_medicine.medicines m 
		inner join emr_medicine.categories c on c.id = m.category_id
		inner join emr_medicine.units u on u.id = m.unit_id 
		` + baseQuery + `
		order by m.name asc 
		limit $2 offset $3`
	err = r.db.Select(&medicines, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return medicines, total, err
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `update emr_medicine.medicines set deleted_at = now() where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
