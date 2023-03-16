package model

import (
	"time"

	"github.com/Krisnaputra15/gsc-solution/db"
	"github.com/Krisnaputra15/gsc-solution/entity"
)

type User struct {
	Base
	Username       string    `json:"username" gorm:"uniqueIndex:username"`
	Email          string    `json:"email" gorm:"uniqueIndex:email"`
	Password       string    `json:"password"`
	IsVerified     bool      `json:"is_verified" gorm:"default:false"`
	Birthdate      time.Time `json:"birthdate"`
	Gender         string    `json:"gender" gorm:"uniqueIndex"`
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
		Email:          userDetail.Email,
		IsVerified:     true,
		ProfilePicture: userDetail.Picture,
		FirstName:      userDetail.GivenName,
		LastName:       userDetail.FamilyName,
		Level:          1,
		Birthdate:      time.Now(),
		Gender:         "Male",
	}

	db.DB.Create(&createUser)
	db.DB.Model(&userModel).First(&user, "id = ?", createUser.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}
