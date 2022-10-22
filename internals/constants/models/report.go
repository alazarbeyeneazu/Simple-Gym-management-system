package models

import "time"

type ReportResponse struct {
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	StartDate  string    `json:"start_date"`
	EndDate    time.Time `json:"endDate"`
	CreatedBy  string    `json:"createdBy"`
	PymentType string    `json:"paymentType"`
	PaidBy     string    `json:"paidBy"`
	Amount     string    `json:"amount"`
}
