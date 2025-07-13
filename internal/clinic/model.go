package clinic

type Entity struct {
	ID                    string `db:"id"`
	Name                  string `db:"name"`
	Address               string `db:"address"`
	Logo                  string `db:"logo"`
	SharingFeeType        string `db:"sharing_fee_type"`
	PatientMedicationCost bool   `db:"patient_medication_cost"`
	NurseSharingFee       bool   `db:"nurse_sharing_fee"`
}

type Request struct {
	Name                  string `json:"name" validate:"required,min=3,max=255"`
	Address               string `json:"address" validate:"required,max=255"`
	Logo                  string `json:"logo" validate:"required,max=255"`
	SharingFeeType        string `json:"sharing_fee_type" validate:"required,max=255"`
	PatientMedicationCost bool   `json:"patient_medication_cost"`
	NurseSharingFee       bool   `json:"nurse_sharing_fee"`
}

type DTO struct {
	ID                    string `json:"id" db:"id"`
	Name                  string `json:"name" db:"name"`
	Address               string `json:"address" db:"address"`
	Logo                  string `json:"logo" db:"logo"`
	SharingFeeType        string `json:"sharing_fee_type" db:"sharing_fee_type"`
	PatientMedicationCost bool   `json:"patient_medication_cost" db:"patient_medication_cost"`
	NurseSharingFee       bool   `json:"nurse_sharing_fee" db:"nurse_sharing_fee"`
}
