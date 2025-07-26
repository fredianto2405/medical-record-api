package auth

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindUserByEmail(email string) (*UserDTO, error) {
	var user UserDTO

	dataQuery := `select u.id, 
		   u.email, 
		   u.password, 
		   r.name role,
		   u.failed_login_attempts,
		   u.is_locked
		from emr_auth.users u 
		join emr_auth.roles r on r.id = u.role_id 
		where email = $1 
		and deleted_at isnull`
	err := r.db.Get(&user, dataQuery, email)

	return &user, err
}

func (r *Repository) UpdateFailedLoginByEmail(email string, isLocked bool) error {
	updateQuery := `update emr_auth.users 
		set failed_login_attempts = failed_login_attempts + 1,
			last_failed_login = now(), 
			is_locked = $1
		where email = $2`
	_, err := r.db.Exec(updateQuery, isLocked, email)
	return err
}

func (r *Repository) ResetFailedLoginByEmail(email string) error {
	updateQuery := `update emr_auth.users 
		set failed_login_attempts = 0, 
			last_failed_login = null,
			is_locked = false
		where email = $1 
		and failed_login_attempts > 0`
	_, err := r.db.Exec(updateQuery, email)
	return err
}

func (r *Repository) UpdatePassword(email string, password string) error {
	updateQuery := `update emr_auth.users set password = $1 where email = $2`
	_, err := r.db.Exec(updateQuery, password, email)
	return err
}

func (r *Repository) SavePasswordReset(userID string) (string, error) {
	insertQuery := `insert into emr_auth.password_resets(user_id, token, expired_at, created_at)
		values($1, uuid_generate_v4(), now() + interval '30 minutes', now())
		returning token`

	var token string
	err := r.db.QueryRow(insertQuery, userID).Scan(&token)
	return token, err
}

func (r *Repository) FindPasswordResetByToken(token string) (string, error) {
	dataQuery := `select u.email 
		from emr_auth.password_resets pr 
		join emr_auth.users u on u.id = pr.user_id 
		where pr.token = $1 
		and pr.used = false 
		and pr.expired_at > now()`

	var email string
	err := r.db.Get(&email, dataQuery, token)

	return email, err
}

func (r *Repository) UpdatePasswordResetUsed(token string) error {
	updateQuery := `update emr_auth.password_resets 
		set used = true 
		where token = $1`
	_, err := r.db.Exec(updateQuery, token)
	return err
}
