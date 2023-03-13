package model

import (
	"time"

	"HealthStep/db"
	"HealthStep/entity"
)

type User struct {
	Base
	Username       string    `json:"username" gorm:"uniqueIndex:username;size:60"`
	Email          string    `json:"email" gorm:"email"`
	IsVerified     bool      `json:"is_verified" gorm:"default:false"`
	Birthdate      time.Time `json:"birthdate"`
	Gender         string    `json:"gender" gorm:"uniqueIndex;size:60"`
	ProfilePicture string    `json:"profile_picture"`
	Level          int       `json:"level"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`

	UserProfile         UserProfile          `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DoctorProfile       DoctorProfile        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DailyHealthChecks   []DailyHealthCheck   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserConsult         []Consultation       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DoctorConsult       []Consultation       `gorm:"foreignKey:DoctorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MedicationSchedules []MedicationSchedule `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type APIUser struct {
	ID             string    `json:"id"`
	Username       string    `json:"username" gorm:"uniqueIndex:username"`
	Email          string    `json:"email" gorm:"uniqueIndex:email"`
	IsVerified     bool      `json:"is_verified" gorm:"default:false"`
	Birthdate      time.Time `json:"birthdate"`
	Gender         string    `json:"gender" gorm:"uniqueIndex"`
	ProfilePicture string    `json:"profile_picture"`
	Level          int       `json:"level"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
}

func UserCreate(userDetail entity.User) (APIUser, error) {
	var userModel User
	var user APIUser
	var err error

	createUser := User{
		Username:       userDetail.Name,
		Email:          userDetail.Email,
		IsVerified:     true,
		ProfilePicture: userDetail.Picture,
		FirstName:      userDetail.GivenName,
		LastName:       userDetail.FamilyName,
		Birthdate:      time.Now(),
	}

	db.DB.Create(&createUser)
	db.DB.Model(&userModel).First(&user, "username = ?", createUser.Username)
	if err != nil {
		return user, err
	}

	return user, nil
}
