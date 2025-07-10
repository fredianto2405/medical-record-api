package patient

import (
	"github.com/jmoiron/sqlx"
	patient "medical-record-api/internal/patient/model"
)

type InsurancePatientRepository struct {
	db *sqlx.DB
}

func NewInsurancePatientRepository(db *sqlx.DB) *InsurancePatientRepository {
	return &InsurancePatientRepository{db: db}
}

func (r *InsurancePatientRepository) Save(e *patient.InsurancePatientEntity) error {
	insertQuery := `insert into emr_patient.insurance_patient(patient_id, 
			insurance_id, 
			insurance_number)
		values(:patient_id, :insurance_id, :insurance_number)`
	_, err := r.db.NamedExec(insertQuery, e)
	return err
}

func (r *InsurancePatientRepository) Delete(patientID string) error {
	deleteQuery := `delete from emr_patient.insurance_patient where patient_id = $1`
	_, err := r.db.Exec(deleteQuery, patientID)
	return err
}
