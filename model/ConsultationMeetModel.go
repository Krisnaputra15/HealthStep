package model

type ConsultationMeet struct {
	Base
	ConsultationID string `json:"consultation_id" gorm:"size:191"`
	Link string
}