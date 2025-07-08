package nurse

import "time"

type Entity struct {
	ID                 string     `db:"id"`
	Name               string     `db:"name"`
	Gender             string     `db:"gender"`
	Address            string     `db:"address"`
	Phone              string     `db:"phone"`
	RegistrationNumber string     `db:"registration_number"`
	SharingFee         float32    `db:"sharing_fee"`
	CreatedAt          *time.Time `db:"created_at"`
	DeletedAt          *time.Time `db:"deleted_at"`
}

type Request struct {
	Name               string  `json:"name" validate:"required,min=3,max=255"`
	Gender             string  `json:"gender" validate:"required,max=255"`
	Phone              string  `json:"phone" validate:"required,numeric,min=10,max=20"`
	Address            string  `json:"address" validate:"required,max=255"`
	RegistrationNumber string  `json:"registration_number" validate:"required,max=255"`
	SharingFee         float32 `json:"sharing_fee" validate:"required"`
}

type DTO struct {
	ID                 string  `json:"id" db:"id"`
	Name               string  `json:"name" db:"name"`
	Gender             string  `json:"gender" db:"gender"`
	Phone              string  `json:"phone" db:"phone"`
	Address            string  `json:"address" db:"address"`
	RegistrationNumber string  `json:"registration_number" db:"registration_number"`
	SharingFee         float32 `json:"sharing_fee" db:"sharing_fee"`
}
