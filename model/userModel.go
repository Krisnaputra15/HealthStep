package model

type User struct {
	Base
	Email string `json:"email"`
	IsVerified bool `json:"is_verified" gorm:"default:false"`
}