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

func (r *RecipeRepository) FindByMedicalRecordID(medicalRecordID string) ([]*model.RecipeDTO, error) {
	var recipes []*model.RecipeDTO

	dataQuery := `select rr.medicine_id,
			m."name" as medicine_name,
			c."name" as category_name,
			u."name" as unit_name,
			rr.price,
			rr.quantity,
			rr.dosage,
			rr.instruction 
		from emr_core.record_recipes rr 
		join emr_medicine.medicines m on m.id = rr.medicine_id 
		join emr_medicine.categories c on c.id = m.category_id 
		join emr_medicine.units u on u.id = m.unit_id 
		where rr.medical_record_id = $1`

	err := r.db.Select(&recipes, dataQuery, medicalRecordID)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
