package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (base *Base) beforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()

	if err != nil {
		err = errors.New("invalid data")
	}

	return
}
