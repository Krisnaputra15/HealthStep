package model

import "database/sql"

type DailyHealthCheck struct {
	Base
	UserID    string         `json:"user_id" gorm:"size:191"`
	IsHealthy bool           `json:"is_healthy"`
	Complaint sql.NullString `json:"complaint"`
}
