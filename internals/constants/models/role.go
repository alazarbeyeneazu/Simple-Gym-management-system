package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID     uuid.UUID `json:"id"`
	Action string    `json:"action"`
	Method string    `json:"method"`
	Path   string    `json:"path"`
	Status bool      `json:"status"`
}
type Role struct {
	gorm.Model
	ID           uuid.UUID `json:"id"`
	RoleName     string    `json:"role_name"`
	PermissionID uuid.UUID `json:"permision_id"`
}
type CreateRoleRequest struct {
	RoleName      string      `json:"role_name"`
	PermissionIDs []uuid.UUID `json:"permissions"`
}
type UserRole struct {
	gorm.Model
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	RoleName string    `json:"role_name"`
}

var Permissions []Permission = []Permission{

	{ID: uuid.New(), Action: "list Users", Method: "GET", Path: "/view/users", Status: true},
	{ID: uuid.New(), Action: "Edit User", Method: "GET", Path: "/view/users/:userid", Status: true},
	{ID: uuid.New(), Action: "Delete User", Method: "GET", Path: "/v1/api/admins/:adminid", Status: true},

	{ID: uuid.New(), Action: "Delete Gymgoer", Method: "GET", Path: "/v1/api/admins/:adminid", Status: true},
	{ID: uuid.New(), Action: "List GymGoer", Method: "GET", Path: "/view/gym-goers", Status: true},
	{ID: uuid.New(), Action: "CheckIn gymgoer", Method: "GET", Path: "/view/gym-goers/:userid", Status: true},
	{ID: uuid.New(), Action: "View Detail of Gym-Goer", Method: "GET", Path: "/view/gym-goers-detail/:userid", Status: true},

	{ID: uuid.New(), Action: "List Payment", Method: "GET", Path: "/view/payment", Status: true},
	{ID: uuid.New(), Action: "Create Payment", Method: "POST", Path: "/v1/api/pyments", Status: true},
	{ID: uuid.New(), Action: "Edit Payment", Method: "GET", Path: "/view/payment/:id", Status: true},

	{ID: uuid.New(), Action: "View Roles", Method: "GET", Path: "/view/roles", Status: true},
	{ID: uuid.New(), Action: "Create New Role", Method: "POST", Path: "/v1/api/roles", Status: true},
	{ID: uuid.New(), Action: "Delete Role", Method: "DELETE", Path: "/v1/api/roles", Status: true},

	{ID: uuid.New(), Action: "view report", Method: "DELETE", Path: "/view/reportbydays", Status: true},
}
