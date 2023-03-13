package model

type ConsultationChat struct {
	Base
	ConsultationID string `json:"consultation_id" gorm:"size:191"`
	Value string `json:"value"`
	IsMedia bool `json:"is_media"`
}