package patient

import (
	"github.com/jmoiron/sqlx"
	patient "medical-record-api/internal/patient/model"
)

type EmergencyContactRepository struct {
	db *sqlx.DB
}

func NewEmergencyContactRepository(db *sqlx.DB) *EmergencyContactRepository {
	return &EmergencyContactRepository{db: db}
}

func (r *EmergencyContactRepository) Save(e *patient.EmergencyContactEntity) error {
	insertQuery := `insert into emr_patient.emergency_contact(patient_id, 
			name, 
			phone, 
			relation)
		values(:patient_id, :name, :phone, :relation)`

	_, err := r.db.NamedExec(insertQuery, e)
	return err
}

func (r *EmergencyContactRepository) Delete(patientID string) error {
	deleteQuery := `delete from emr_patient.emergency_contact where patient_id = $1`

	_, err := r.db.Exec(deleteQuery, patientID)
	return err
}
