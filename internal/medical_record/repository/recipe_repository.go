package medical_record

import (
	"github.com/jmoiron/sqlx"
	model "medical-record-api/internal/medical_record/model"
)

type RecipeRepository struct {
	db *sqlx.DB
}

func NewRecipeRepository(db *sqlx.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

func (r *RecipeRepository) Save(tx *sqlx.Tx, e *model.RecipeEntity) error {
	insertQuery := `insert into emr_core.record_recipes(medical_record_id, 
		medicine_id, 
		price, 
		quantity, 
		dosage, 
		instruction)
	values(:medical_record_id, :medicine_id, :price, 
	       :quantity, :dosage, :instruction)`

	_, err := tx.NamedExec(insertQuery, e)
	return err
}

func (r *RecipeRepository) Delete(medicalRecordID, medicineID string) error {
	deleteQuery := `delete from emr_core.record_recipes
		where medical_record_id = $1
		and medicine_id = $2`
	_, err := r.db.Exec(deleteQuery, medicalRecordID, medicineID)
	return err
}

func (r *RecipeRepository) DeleteAll(medicalRecordID string) error {
	deleteQuery := `delete from emr_core.record_recipes
		where medical_record_id = $1`
	_, err := r.db.Exec(deleteQuery, medicalRecordID)
	return err
}
