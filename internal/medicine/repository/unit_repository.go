package medicine

import (
	"github.com/jmoiron/sqlx"
	medicine "medical-record-api/internal/medicine/model"
)

type UnitRepository struct {
	db *sqlx.DB
}

func NewUnitRepository(db *sqlx.DB) *UnitRepository {
	return &UnitRepository{db: db}
}

func (r *UnitRepository) Save(e *medicine.UnitEntity) (*medicine.UnitEntity, error) {
	insertQuery := `insert into emr_medicine.units(name) values(:name) returning id, name`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result medicine.UnitEntity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *UnitRepository) Update(e *medicine.UnitEntity) (*medicine.UnitEntity, error) {
	updateQuery := `update emr_medicine.units set name = :name where id = :id returning id, name`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result medicine.UnitEntity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *UnitRepository) FindAll(search string) ([]*medicine.UnitDTO, error) {
	var medicineUnits []*medicine.UnitDTO

	searchPattern := "%" + search + "%"

	dataQuery := `select id, name
		from emr_medicine.units
		where deleted_at isnull 
		and name ilike $1
		order by name asc`
	err := r.db.Select(&medicineUnits, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return medicineUnits, nil
}

func (r *UnitRepository) FindAllPaginated(page, limit int, search string) ([]*medicine.UnitDTO, int, error) {
	var medicineUnits []*medicine.UnitDTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where deleted_at isnull and name ilike $1`

	countQuery := `select count(0) from emr_medicine.units ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select id, name
		from emr_medicine.units
		` + baseQuery + `
		order by name asc
		limit $2 offset $3`
	err = r.db.Select(&medicineUnits, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return medicineUnits, total, err
}

func (r *UnitRepository) Delete(id string) error {
	deleteQuery := `update emr_medicine.units set deleted_at = now() where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
