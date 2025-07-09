package medicine

import (
	"github.com/jmoiron/sqlx"
	medicine "medical-record-api/internal/medicine/model"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Save(e *medicine.CategoryEntity) (*medicine.CategoryEntity, error) {
	insertQuery := `insert into emr_medicine.categories(name) values(:name) returning id, name`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result medicine.CategoryEntity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *CategoryRepository) Update(e *medicine.CategoryEntity) (*medicine.CategoryEntity, error) {
	updateQuery := `update emr_medicine.categories set name = :name where id = :id returning id, name`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result medicine.CategoryEntity
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *CategoryRepository) FindAll(search string) ([]*medicine.CategoryDTO, error) {
	var medicineCategories []*medicine.CategoryDTO

	searchPattern := "%" + search + "%"

	dataQuery := `select id, name 
		from emr_medicine.categories 
		where deleted_at isnull 
		and name ilike $1
		order by name asc`
	err := r.db.Select(&medicineCategories, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return medicineCategories, nil
}

func (r *CategoryRepository) FindAllPaginated(page, limit int, search string) ([]*medicine.CategoryDTO, int, error) {
	var medicineCategories []*medicine.CategoryDTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where deleted_at isnull and name ilike $1`

	countQuery := `select count(0) from emr_medicine.categories ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select id, name
		from emr_medicine.categories
		` + baseQuery + `
		order by name asc
		limit $2 offset $3`
	err = r.db.Select(&medicineCategories, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return medicineCategories, total, err
}

func (r *CategoryRepository) Delete(id string) error {
	deleteQuery := `update emr_medicine.categories set deleted_at = now() where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
