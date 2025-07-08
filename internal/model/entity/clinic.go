package entity

type Clinic struct {
	ID                    string `db:"id"`
	Name                  string `db:"name"`
	Address               string `db:"address"`
	Logo                  string `db:"logo"`
	SharingFeeType        string `db:"sharing_fee_type"`
	PatientMedicationCost bool   `db:"patient_medication_cost"`
	NurseSharingFee       bool   `db:"nurse_sharing_fee"`
}
