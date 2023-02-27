package model

import (
	"time"
)

type Profile struct {
	Base
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Birthdate time.Time `json:"birthdate "`
	Gender    string    `json:"gender" gorm:"uniqueIndex"`
}
