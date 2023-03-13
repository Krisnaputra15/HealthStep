package model

import "time"

type Consultation struct {
	Base
	UserID            string    `json:"user_id" gorm:"size:191"`
	DoctorID          string    `json:"doctor_id" gorm:"size:191"`
	ConsultationDate  time.Time `json:"consultation_date"`
	ConsultationEnd   time.Time `json:"consultation_end"`
	HealthTitle       string    `json:"health_title"`
	HealthDescription string    `json:"health_description"`
	Star              int       `json:"star"`
	IsEnded           bool      `json:"is_end"`
	Price             int32     `json:"price"`

	ConsultationMeets []ConsultationMeet `gorm:"foreignKey:ConsultationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ConsultationChats []ConsultationChat `gorm:"foreignKey:ConsultationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Payments          []Payment          `gorm:"foreignKey:ConsultationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
