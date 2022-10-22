package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gym_goers struct {
	gorm.Model
	ID                   uuid.UUID `json:"id"`
	UserId               uuid.UUID `json:"user_id"`
	CreatedByFirstName   string    `json:"created_by_firstname"`
	CreatedByLastName    string    `json:"created_by_lastname"`
	CreatedByPhoneNumber string    `json:"created_by_phonenumber"`
	StartDate            time.Time `json:"start_date"`
	EndDate              time.Time `json:"end_date"`
	PaidBy               string    `json:"paid_by"`
	PaidAmount           string    `json:"paid_amount"`
}

type Gym_goerRequest struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	PaymentType uuid.UUID `json:"payment_type"`
	Start_date  string    `json:"start_date"`
	PaidBy      string    `json:"paid_by"`
}
