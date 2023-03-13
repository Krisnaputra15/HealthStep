package model

import "time"

type DoctorSchedule struct {
	Base
	DoctorID string    `json:"doctor_id" gorm:"size:191"`
	Date     time.Time `json:"date"`
	IsBooked bool      `json:"is_booked" gorm:"default:true"`
	IsActive bool      `json:"is_active"`
}
