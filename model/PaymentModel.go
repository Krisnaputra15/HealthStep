package model

type Payment struct {
	Base
	ConsultationID string `json:"consultation_id" gorm:"size:191"`
	PaymentMethod  string `json:"payment_method"`
	Amount         int32  `json:"amount"`
}
