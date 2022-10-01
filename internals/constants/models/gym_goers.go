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
}
