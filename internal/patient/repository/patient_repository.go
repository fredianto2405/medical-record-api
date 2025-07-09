package patient

import (
	"github.com/jmoiron/sqlx"
	patient "medical-record-api/internal/patient/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e *patient.Entity) (*patient.Entity, error) {
	insertQuery := `insert into emr_patient.patients(name, 
			gender, 
			birth_date, 
			blood_type, 
			address, 
			phone, 
			email, 
			history_of_illness, 
			allergies, 
			identity_type, 
			identity_number,
			medical_record_number)
		values(:name, 
		       :gender, 
		       :birth_date, 
		       :blood_type, 
		       :address, 
		       :phone, 
		       :email, 
		       :history_of_illness, 
		       :allergies, 
		       :identity_type, 
		       :identity_number, 
		       emr_patient.generate_medical_record_number())
		returning id,
			name, 
			gender, 
			birth_date, 
			blood_type, 
			address, 
			phone, 
			email, 
			history_of_illness, 
			allergies, 
			identity_type, 
			identity_number,
			medical_record_number`

	rows, err := r.db.NamedQuery(insertQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result patient.Entity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Gender,
			&result.BirthDate,
			&result.BloodType,
			&result.Address,
			&result.Phone,
			&result.Email,
			&result.HistoryOfIllness,
			&result.Allergies,
			&result.IdentityType,
			&result.IdentityNumber,
			&result.MedicalRecordNumber)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) Update(e *patient.Entity) (*patient.Entity, error) {
	updateQuery := `update emr_patient.patients
		set name = :name,
			gender = :gender,
			birth_date = :birth_date,
			blood_type = :blood_type,
			address = :address,
			phone = :phone,
			email = :email,
			history_of_illness = :history_of_illness,
			allergies = :allergies,
			identity_type = :identity_type,
			identity_number = :identity_number
		where id = :id
		returning id,
			name, 
			gender, 
			birth_date, 
			blood_type, 
			address, 
			phone, 
			email, 
			history_of_illness, 
			allergies, 
			identity_type, 
			identity_number,
			medical_record_number`

	rows, err := r.db.NamedQuery(updateQuery, e)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result patient.Entity
	if rows.Next() {
		err = rows.Scan(&result.ID,
			&result.Name,
			&result.Gender,
			&result.BirthDate,
			&result.BloodType,
			&result.Address,
			&result.Phone,
			&result.Email,
			&result.HistoryOfIllness,
			&result.Allergies,
			&result.IdentityType,
			&result.IdentityNumber,
			&result.MedicalRecordNumber)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *Repository) FindAll(search string) ([]*patient.DTO, error) {
	var patients []*patient.DTO

	searchPattern := "%" + search + "%"

	dataQuery := `select p.id, 
			p.name, 
			p.gender, 
			to_char(p.birth_date, 'YYYY-MM-DD') birth_date, 
			p.blood_type, 
			p.address, 
			p.phone,
			p.email, 
			p.history_of_illness, 
			p.allergies,
			p.identity_type, 
			p.identity_number, 
			coalesce(ec.name, '-') emergency_contact_name,
			coalesce(ec.phone, '-') emergency_contact_phone, 
			coalesce(ec.relation, '-') emergency_contact_relation, 
			coalesce(ip.insurance_id::text, '-') insurance_id,
			coalesce(ip.insurance_number, '-') insurance_number,
			p.medical_record_number
		from emr_patient.patients p
		left join emr_patient.emergency_contact ec on ec.patient_id = p.id
		left join emr_patient.insurance_patient ip on ip.patient_id = p.id
		where p.deleted_at isnull and p.name ilike $1 
		order by p.name asc`
	err := r.db.Select(&patients, dataQuery, searchPattern)

	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (r *Repository) FindAllPaginated(page, limit int, search string) ([]*patient.DTO, int, error) {
	var patients []*patient.DTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where p.deleted_at isnull and p.name ilike $1`

	countQuery := `select count(0) from emr_patient.patients p ` + baseQuery
	err := r.db.Get(&total, countQuery, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select p.id, 
			p.name, 
			p.gender, 
			to_char(p.birth_date, 'YYYY-MM-DD') birth_date, 
			p.blood_type, 
			p.address, 
			p.phone,
			p.email, 
			p.history_of_illness, 
			p.allergies,
			p.identity_type, 
			p.identity_number, 
			coalesce(ec.name, '-') emergency_contact_name,
			coalesce(ec.phone, '-') emergency_contact_phone, 
			coalesce(ec.relation, '-') emergency_contact_relation, 
			coalesce(ip.insurance_id::text, '-') insurance_id,
			coalesce(ip.insurance_number, '-') insurance_number,
			p.medical_record_number
		from emr_patient.patients p
		left join emr_patient.emergency_contact ec on ec.patient_id = p.id
		left join emr_patient.insurance_patient ip on ip.patient_id = p.id
		` + baseQuery + `
		order by p.name asc
		limit $2 offset $3`
	err = r.db.Select(&patients, dataQuery, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return patients, total, err
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `update emr_patient.patients set deleted_at = now() where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
