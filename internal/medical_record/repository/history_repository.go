package medical_record

import (
	"github.com/jmoiron/sqlx"
	model "medical-record-api/internal/medical_record/model"
)

type HistoryRepository struct {
	db *sqlx.DB
}

func NewHistoryRepository(db *sqlx.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (r *HistoryRepository) Save(tx *sqlx.Tx, e *model.HistoryEntity) error {
	insertQuery := `insert into emr_core.medical_record_histories(medical_record_id, status_id) 
		values(:medical_record_id, :status_id)`

	_, err := tx.NamedExec(insertQuery, e)
	return err
}

func (r *HistoryRepository) Delete(medicalRecordID string) error {
	deleteQuery := `delete from emr_core.medical_record_histories 
       where medical_record_id = $1`
	_, err := r.db.Exec(deleteQuery, medicalRecordID)
	return err
}
