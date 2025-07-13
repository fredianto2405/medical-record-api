package medical_record

import (
	"github.com/jmoiron/sqlx"
	model "medical-record-api/internal/medical_record/model"
)

type StatusRepository struct {
	db *sqlx.DB
}

func NewStatusRepository(db *sqlx.DB) *StatusRepository {
	return &StatusRepository{db: db}
}

func (r *StatusRepository) FindAll() ([]*model.StatusDTO, error) {
	var statuses []*model.StatusDTO

	dataQuery := `select id, 
			name, 
			description
		from emr_core.medical_record_statuses
		where deleted_at isnull 
		order by sort_number asc`

	err := r.db.Select(&statuses, dataQuery)
	return statuses, err
}
