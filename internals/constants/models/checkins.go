package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Checkins struct {
	gorm.Model
	ID            uuid.UUID `json:"id"`
	UserId        uuid.UUID `json:"user_id"`
	CheckedInDate time.Time `json:"checking_date"`
}

type CheckinResponse struct {
	UserFirstName       string    `json:"user_first_name"`
	UserNumberOfDayLeft string    `json:"left_days"`
	UserLastName        string    `json:"user_last_name"`
	IsChackedIn         string    `json:"isChackedIn"`
	CheckedInDate       time.Time `json:"checking_date"`
}

type DateResponse struct {
	Month        string `json:"month"`
	DayMonthYear string `json:"DayMonthYear"`
	Hour         string `json:"hour"`
}
