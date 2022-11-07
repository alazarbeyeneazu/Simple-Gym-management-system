package models

import (
	"time"

	"gorm.io/gorm"
)

type ReportResponse struct {
	gorm.Model
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"endDate"`
	CreatedBy  string    `json:"createdBy"`
	PymentType string    `json:"paymentType"`
	PaidBy     string    `json:"paidBy"`
	Amount     string    `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

type HttpReportResponse struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"endDate"`
	CreatedBy  string `json:"createdBy"`
	PymentType string `json:"paymentType"`
	PaidBy     string `json:"paidBy"`
	Amount     string `json:"amount"`
	CreatedAt  string `json:"created_at"`
}
