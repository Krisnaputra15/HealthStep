package migration

import (
	"HealthStep/db"
	"HealthStep/model"
)

func Migrate() error {
	return db.DB.AutoMigrate(
		model.User{}, model.UserProfile{}, model.DailyHealthCheck{}, model.DoctorProfile{}, model.DoctorSchedule{}, model.MedicationSchedule{},
		model.MedicationHistory{}, model.Consultation{}, model.ConsultationMeet{}, model.ConsultationChat{}, model.Payment{},
	)
}
