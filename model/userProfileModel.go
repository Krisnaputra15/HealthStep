package model

import (
	"HealthStep/db"
	"HealthStep/entity"
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

func UserProfileCreate(userDetail entity.User, id string) (APIUserProfile, error) {
	var userProfile APIUserProfile

	// Check if the user exists
	var user entity.User
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return userProfile, err
	}

	// Create user profile
	createUserProfile := UserProfile{
		UserID:      id,
		HealthPoint: 0,
		IsVolunteer: false,
		Latitude:    0,
		Longitude:   0,
	}

	if err := db.DB.Create(&createUserProfile).Error; err != nil {
		return userProfile, err
	}

	return userProfile, nil
}
