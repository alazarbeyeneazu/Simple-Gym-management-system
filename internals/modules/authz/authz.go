package authz

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
)

type AuthService interface {
	CreatePermission(ctx context.Context, permission models.Permission) (models.Permission, error)
	CreateRole(ctx context.Context, role models.Role) (models.Role, error)
	GetAllPermission(ctx context.Context) ([]models.Permission, error)
	GetAllRoles(ctx context.Context) ([]models.Role, error)
	GetRolesByName(ctx context.Context, role models.Role) ([]models.Role, error)
	GetPermissionById(ctx context.Context, permission models.Permission) (models.Permission, error)
	DeleteRole(ctx context.Context, role models.Role) error
	InitatePermission()
	GetAssignRoleByUserId(ctx context.Context, role models.UserRole) (models.UserRole, error)
}
type authService struct {
	db ports.DBPort
}

func InitService(db ports.DBPort) AuthService {
	return &authService{
		db: db,
	}
}

func (az *authService) InitatePermission() {
	for _, permissions := range models.Permissions {
		az.db.CreatePermissions(context.Background(), permissions)
	}
}

func (az *authService) GetAllPermission(ctx context.Context) ([]models.Permission, error) {
	return az.db.GetAllPermissions(ctx)
}
