package medical_record

import (
	"github.com/jmoiron/sqlx"
	model "medical-record-api/internal/medical_record/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(tx *sqlx.Tx, e *model.Entity) (string, error) {
	insertQuery := `insert into emr_core.medical_records(patient_id, doctor_id, diagnosis,
			notes, status_id, payment_method_id,
			payment_status_id, insurance_id, anamnesis)
		values(:patient_id, :doctor_id, :diagnosis, 
		       :notes, :status_id, :payment_method_id, :payment_status_id, 
		       :insurance_id, :anamnesis)
		returning id`

	rows, err := tx.NamedQuery(insertQuery, e)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var id string
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return "", err
		}
	}

	return id, nil
}

func (r *Repository) Delete(id string) error {
	deleteQuery := `delete from emr_core.medical_records where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}

func (r *Repository) Update(e *model.Entity) error {
	updateQuery := `update emr_core.medical_records 
		set diagnosis = :diagnosis, 
		    notes = :notes,
		    payment_method_id = :payment_method_id,
		    payment_status_id = :payment_status_id,
		    insurance_id = :insurance_id,
		    anamnesis = :anamnesis
		where id = :id`
	_, err := r.db.NamedExec(updateQuery, e)
	return err
}

func (r *Repository) UpdateStatus(tx *sqlx.Tx, id string, statusID int) error {
	updateQuery := `update emr_core.medical_records set status_id = $1 where id = $2`
	_, err := tx.Exec(updateQuery, statusID, id)
	return err
}

func (r *Repository) FindAllPaginated(startDate string, endDate string, page, limit int, search string) ([]*model.DTO, int, error) {
	var records []*model.DTO
	var total int

	searchPattern := "%" + search + "%"

	baseQuery := `where 1 = 1 
		and ($1 = '' or mr.created_at::date >= $1::date)
		and ($2 = '' or mr.created_at::date <= $2::date)
		and ($3 = '' or p.name ilike $3)
		and ($3 = '' or p.medical_record_number ilike $3)`

	countQuery := `select count(0)
		from emr_core.medical_records mr
		join emr_patient.patients p on p.id = mr.patient_id 
		left join emr_patient.insurance_patient ip on ip.patient_id = p.id
		join emr_doctor.doctors d on d.id = mr.doctor_id 
		join emr_doctor.specializations s on s.id = d.specialization_id 
		join emr_payment.payment_methods pm on pm.id = mr.payment_method_id 
		join emr_payment.payment_statuses ps on ps.id = mr.payment_status_id 
		join emr_payment.insurances i on i.id = mr.insurance_id
		join emr_core.medical_record_statuses mrs on mrs.id = mr.status_id ` + baseQuery

	err := r.db.Get(&total, countQuery, startDate, endDate, searchPattern)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `select mr.id,
			to_char(mr.created_at, 'YYYY-MM-DD HH24:MI') as created_at,
			mr.patient_id,
			p."name" as patient_name,
			p.gender as patient_gender,
			p.address as patient_address,
			to_char(p.birth_date, 'YYYY-MM-DD') as patient_birth_date,
			p.medical_record_number,
			mr.doctor_id,
			d."name" as doctor_name,
			s."name" as doctor_specialization,
			mr.diagnosis,
			mr.notes,
			mr.status_id,
			mrs."name" as status_name,
			mr.payment_method_id,
			pm."name" as payment_method_name,
			mr.payment_status_id,
			ps."name" as payment_status_name,
			mr.insurance_id,
			i."name" as insurance_name,
			coalesce(ip.insurance_number, '-') as insurance_number,
			mr.anamnesis 
		from emr_core.medical_records mr
		join emr_patient.patients p on p.id = mr.patient_id 
		left join emr_patient.insurance_patient ip on ip.patient_id = p.id
		join emr_doctor.doctors d on d.id = mr.doctor_id 
		join emr_doctor.specializations s on s.id = d.specialization_id 
		join emr_payment.payment_methods pm on pm.id = mr.payment_method_id 
		join emr_payment.payment_statuses ps on ps.id = mr.payment_status_id 
		join emr_payment.insurances i on i.id = mr.insurance_id
		join emr_core.medical_record_statuses mrs on mrs.id = mr.status_id 
		` + baseQuery + `
		order by mr.created_at desc
		limit $4 offset $5`

	err = r.db.Select(&records, dataQuery, startDate, endDate, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}
