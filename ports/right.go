package ports

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/google/uuid"
)

//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package mockdb -destination ../mocks/db/mockdb.go github.com/alazarbeyeneazu/Simple-Gym-management-system/ports DBPort
type DBPort interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, newUser, user models.User) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserByFirstName(ctx context.Context, firstname string) ([]models.User, error)
	GetUserByLastName(ctx context.Context, lastname string) ([]models.User, error)
	GetUseByPhoneNumber(ctx context.Context, phonenumber string) (models.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (models.User, error)
	CreatePymentType(ctx context.Context, pyment models.PymentType) (models.PymentType, error)
	DeletePyment(ctx context.Context, pyment models.PymentType) error
	GetAllPyments(ctx context.Context) ([]models.PymentType, error)
	GetPymentById(ctx context.Context, pyment models.PymentType) (models.PymentType, error)
	UpdatePyment(ctx context.Context, pyment models.PymentType) (models.PymentType, error)
	CreateGymGoers(ctx context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error)
	DeleteGymGoers(ctx context.Context, gymGors models.Gym_goers) error
	GetAllGymGoers(cxt context.Context) ([]models.Gym_goers, error)
	GetGYmGorsById(cxt context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error)
	GetGymGoerByUserId(ctx context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error)
	GetGymGoerByCreatedByFirstName(ctx context.Context, gym_goers models.User) ([]models.Gym_goers, error)
	GetGymGoerByCreatedByLastName(ctx context.Context, gym_goers models.User) ([]models.Gym_goers, error)
	GetGymGoerByCreatedByPhoneNumber(ctx context.Context, gym_goers models.User) ([]models.Gym_goers, error)
	GetGymGoerByPaidBy(ctx context.Context, gym_goers models.Gym_goers) ([]models.Gym_goers, error)
	UpdateGymGoer(ctx context.Context, newGymGoer models.Gym_goers) (models.Gym_goers, error)
	CreateAdmin(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error)
	GetAllAdmins() ([]models.AdminUsers, error)
	GetAdminByUserId(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error)
	DeleteAdmin(ctx context.Context, admin models.AdminUsers) error
	GetAdminById(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error)
	CreatePermissions(ctx context.Context, permission models.Permission) (models.Permission, error)
	CreateRole(ctx context.Context, role models.Role) (models.Role, error)
	GetAllPermissions(ctx context.Context) ([]models.Permission, error)
	GetAllRoles(ctx context.Context) ([]models.Role, error)
	GetRolesByName(ctx context.Context, role models.Role) ([]models.Role, error)
	GetPermissionById(ctx context.Context, permission models.Permission) (models.Permission, error)
	DeleteRole(ctx context.Context, role models.Role) error
	AssignRole(ctx context.Context, role models.UserRole) (models.UserRole, error)
	GetAssignRoleByUserId(ctx context.Context, role models.UserRole) (models.UserRole, error)
	CheckInUser(ctx context.Context, user models.Checkins) (models.Checkins, error)
	GetCheckedInByUserId(ctx context.Context, user models.Checkins) ([]models.Checkins, error)
	GetAllCheckIns(ctx context.Context) ([]models.Checkins, error)
	CreateReport(ctx context.Context, report models.ReportResponse) error
	GetAllReports(ctx context.Context) ([]models.ReportResponse, error)
}
