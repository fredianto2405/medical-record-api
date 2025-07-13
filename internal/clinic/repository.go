package clinic

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Update(e *Entity) (*Entity, error) {
	updateQuery := `update emr_clinic.clinics
		set name = :name,
		    address = :address,
		    logo = :logo,
		    sharing_fee_type = :sharing_fee_type,
		    patient_medication_cost = :patient_medication_cost,
		    nurse_sharing_fee = :nurse_sharing_fee
		where id = :id
		returning id, name, address,
		logo, sharing_fee_type, patient_medication_cost, 
		nurse_sharing_fee`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Entity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Address,
			&result.Logo,
			&result.SharingFeeType,
			&result.PatientMedicationCost,
			&result.NurseSharingFee)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) FindOne() (*DTO, error) {
	var clinic DTO

	dataQuery := `select id, name, address,
			logo, sharing_fee_type, patient_medication_cost,
			nurse_sharing_fee 
		from emr_clinic.clinics 
		limit 1`

	err := r.db.Get(&clinic, dataQuery)
	if err != nil {
		return nil, err
	}

	return &clinic, nil
}
