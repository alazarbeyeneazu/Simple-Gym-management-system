package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PymentType struct {
	gorm.Model
	ID                 uuid.UUID `json:"id"`
	PymentType         string    `json:"pyment_type"`
	CreatedByFirstName string    `json:"created_by_firstname"`
	CreatedByLastName  string    `json:"created_by_lastname"`
	Payment            string    `json:"pyment"`
	NumberOfDays       int64     `json:"number_of_days"`
}
