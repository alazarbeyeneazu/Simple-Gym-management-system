package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
type AdminUsers struct {
	gorm.Model
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
}
type CreateAdminRequest struct {
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	FirstName   string `json:"first_name"`
	Password    string `json:"password"`
	LastName    string `json:"last_name"`
}

type PermissionMap struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}
