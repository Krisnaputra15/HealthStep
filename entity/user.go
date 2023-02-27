package entity

type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	IsEmailVerified bool `json:"is_email_verified"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	Picture string `json:"picture"`
	Language string `json:"language"`
}