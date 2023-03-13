package main

import (
	"github.com/Krisnaputra15/gsc-solution/config"
	"github.com/Krisnaputra15/gsc-solution/db"
	"github.com/Krisnaputra15/gsc-solution/model"
)

func init() {
	config.LoadEnv()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(
		model.User{}, model.UserProfile{}, model.DailyHealthCheck{}, model.DoctorProfile{}, model.DoctorSchedule{}, model.MedicationSchedule{},
		model.MedicationHistory{}, model.Consultation{}, model.ConsultationMeet{}, model.ConsultationChat{}, model.Payment{},
	)
}
