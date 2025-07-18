package medical_record

import (
	"github.com/jmoiron/sqlx"
	model "medical-record-api/internal/medical_record/model"
)

type NurseAssignmentRepository struct {
	db *sqlx.DB
}

func NewNurseAssignmentRepository(db *sqlx.DB) *NurseAssignmentRepository {
	return &NurseAssignmentRepository{db: db}
}

func (r *NurseAssignmentRepository) Save(tx *sqlx.Tx, e *model.NurseAssignmentEntity) error {
	insertQuery := `insert into emr_core.record_nurse_assignments(medical_record_id, nurse_id) 
		values(:medical_record_id, :nurse_id)`

	_, err := tx.NamedExec(insertQuery, e)
	return err
}

func (r *NurseAssignmentRepository) Delete(medicalRecordID, nurseID string) error {
	deleteQuery := `delete from emr_core.record_nurse_assignments
		where medical_record_id = $1
		and nurse_id = $2`
	_, err := r.db.Exec(deleteQuery, medicalRecordID, nurseID)
	return err
}

func (r *NurseAssignmentRepository) DeleteAll(medicalRecordID string) error {
	deleteQuery := `delete from emr_core.record_nurse_assignments
		where medical_record_id = $1`
	_, err := r.db.Exec(deleteQuery, medicalRecordID)
	return err
}

func (r *NurseAssignmentRepository) FindByMedicalRecordID(medicalRecordID string) ([]*model.NurseDTO, error) {
	var nurses []*model.NurseDTO

	dataQuery := `select rna.nurse_id, n."name" as nurse_name
		from emr_core.record_nurse_assignments rna 
		join emr_nurse.nurses n on n.id = rna.nurse_id 
		where rna.medical_record_id = $1`

	err := r.db.Select(&nurses, dataQuery, medicalRecordID)
	if err != nil {
		return nil, err
	}

	return nurses, nil
}
