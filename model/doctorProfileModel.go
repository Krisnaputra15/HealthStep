package model

type DoctorProfile struct {
	Base
	UserID         string `json:"user_id" gorm:"size:191"`
	Specialization string `json:"specialization"`
}
