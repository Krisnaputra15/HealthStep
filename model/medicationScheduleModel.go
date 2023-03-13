package model

import "time"

type MedicationSchedule struct {
	Base
	UserID             string    `json:"user_id" gorm:"size:191"`
	MedicineName       string    `json:"medicine_name"`
	PerDayConsumptions int       `json:"per_day_consumptions"`
	ConsumeUntil       time.Time `json:"consume_until"`

	MedicationHistories []MedicationHistory `gorm:"foreignKey:MedicationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
