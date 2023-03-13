package model

import (
	"github.com/Krisnaputra15/gsc-solution/db"
	"github.com/Krisnaputra15/gsc-solution/entity"
)

type UserProfile struct {
	Base
	UserID      string  `json:"id" gorm:"size:191"`
	HealthPoint int64   `json:"health_point"`
	IsVolunteer bool    `json:"is_volunteer"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type APIUserProfile struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id" gorm:"size:191"`
	HealthPoint int64   `json:"health_point"`
	IsVolunteer bool    `json:"is_volunteer"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

func UserProfileCreate(userDetail entity.User) (APIUserProfile, error) {
	var userProfile APIUserProfile

	createUserProfile := UserProfile{
		UserID: userDetail.ID,
		HealthPoint: 0,
		IsVolunteer: false,
		Latitude: 0,
		Longitude: 0,
	}

	resultCreate := db.DB.Create(&createUserProfile)
	if resultCreate.Error != nil {
		return userProfile, resultCreate.Error
	}
	db.DB.Model(&userProfile).First(&userProfile, "user_id = ?", createUserProfile.UserID)

	return userProfile, nil
}
