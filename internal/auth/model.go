package auth

type UserDTO struct {
	ID                  string `db:"id"`
	Email               string `db:"email"`
	Password            string `db:"password"`
	Role                string `db:"role"`
	FailedLoginAttempts int    `db:"failed_login_attempts"`
	IsLocked            bool   `db:"is_locked"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginDTO struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
}

type ChangePasswordRequest struct {
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	NewPassword        string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}
