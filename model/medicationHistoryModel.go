package model

type MedicationHistory struct {
	Base
	MedicationID string `json:"medication_id" gorm:"size:191"`
}
