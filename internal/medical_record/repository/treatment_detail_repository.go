package medical_record

import (
	"github.com/jmoiron/sqlx"
	model "medical-record-api/internal/medical_record/model"
)

type TreatmentDetailRepository struct {
	db *sqlx.DB
}

func NewTreatmentDetailRepository(db *sqlx.DB) *TreatmentDetailRepository {
	return &TreatmentDetailRepository{db: db}
}

func (r *TreatmentDetailRepository) Save(tx *sqlx.Tx, e *model.TreatmentDetailEntity) error {
	insertQuery := `insert into emr_core.record_treatment_details(medical_record_id, 
			treatment_id, 
			price) 
		values(:medical_record_id, :treatment_id, :price)`

	_, err := tx.NamedExec(insertQuery, e)
	return err
}

func (r *TreatmentDetailRepository) Delete(medicalRecordID, treatmentID string) error {
	deleteQuery := `delete from emr_core.record_treatment_details
		where medical_record_id = $1
		and treatment_id = $2`
	_, err := r.db.Exec(deleteQuery, medicalRecordID, treatmentID)
	return err
}

func (r *TreatmentDetailRepository) DeleteAll(medicalRecordID string) error {
	deleteQuery := `delete from emr_core.record_treatment_details
		where medical_record_id = $1`
	_, err := r.db.Exec(deleteQuery, medicalRecordID)
	return err
}

func (r *TreatmentDetailRepository) FindByMedicalRecordID(medicalRecordID string) ([]*model.TreatmentDTO, error) {
	var treatments []*model.TreatmentDTO

	dataQuery := `select rtd.treatment_id,
			t."name" as treatment_name,
			rtd.price
		from emr_core.record_treatment_details rtd 
		join emr_treatment.treatments t on t.id = rtd.treatment_id 
		where rtd.medical_record_id = $1`

	err := r.db.Select(&treatments, dataQuery, medicalRecordID)
	if err != nil {
		return nil, err
	}

	return treatments, nil
}
