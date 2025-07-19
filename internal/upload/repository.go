package upload

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e *Entity) (string, error) {
	insertQuery := `insert into emr_clinic.uploaded_files(filename, filepath, size, mime_type)
		values(:filename, :filepath, :size, :mime_type)
		returning id`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var id string
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return "", err
		}
	}

	return id, nil
}
