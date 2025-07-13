package medical_record

import (
	"github.com/jmoiron/sqlx"
	model "medical-record-api/internal/medical_record/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(tx *sqlx.Tx, e *model.Entity) (string, error) {
	insertQuery := `insert into emr_core.medical_records(patient_id, doctor_id, diagnosis,
			notes, status_id, payment_method_id,
			payment_status_id, insurance_id, anamnesis)
		values(:patient_id, :doctor_id, :diagnosis, 
		       :notes, :status_id, :payment_method_id, :payment_status_id, 
		       :insurance_id, :anamnesis)
		returning id`

	rows, err := tx.NamedQuery(insertQuery, e)
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

func (r *Repository) Delete(id string) error {
	deleteQuery := `delete from emr_core.medical_records where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
