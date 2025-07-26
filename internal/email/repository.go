package email

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) QueueEmail(recipient, subject, body string, isHTML bool) error {
	insertQuery := `insert into emr_clinic.email_queue(recipient, subject, body, is_html)
		values($1, $2, $3, $4)`
	_, err := r.db.Exec(insertQuery, recipient, subject, body, isHTML)
	return err
}

func (r *Repository) FindAllPendingEmail() ([]DTO, error) {
	dataQuery := `select id, 
			recipient, 
			subject,
			body,
			is_html
		from emr_clinic.email_queue
		where status = 'pending'`

	var emails []DTO
	err := r.db.Select(&emails, dataQuery)
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (r *Repository) UpdateQueueFailed(id int, errorMessage string) error {
	updateQuery := `update emr_clinic.email_queue 
		set status = 'failed', error_message = $1
		where id = $2`
	_, err := r.db.Exec(updateQuery, errorMessage, id)
	return err
}

func (r *Repository) UpdateQueueSuccess(id int) error {
	updateQuery := `update emr_clinic.email_queue
		set status = 'sent', sent_at = now()
		where id = $1`
	_, err := r.db.Exec(updateQuery, id)
	return err
}
