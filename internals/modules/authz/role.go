package authz

import (
	"context"
	"log"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

func (az *authService) CreateRole(ctx context.Context, role models.Role) (models.Role, error) {
	err := validation.ValidateStruct(&role,
		validation.Field(&role.RoleName, validation.Required.Error("role name required")),
		validation.Field(&role.PermissionID, is.UUID, validation.NotIn(uuid.Nil).Error("invalide permission id")),
	)
	if err != nil {
		return models.Role{}, err
	}
	return az.db.CreateRole(ctx, role)
}
func (az *authService) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	return az.db.GetAllRoles(ctx)
}

func (az *authService) GetRolesByName(ctx context.Context, role models.Role) ([]models.Role, error) {
	err := validation.Validate(&role.RoleName)
	if err != nil {
		log.Println(err)
		return []models.Role{}, err
	}
	return az.db.GetRolesByName(ctx, role)
}

func (az *authService) DeleteRole(ctx context.Context, role models.Role) error {

	if err := validation.Validate(&role.RoleName, validation.Required.Error("role name required")); err != nil {
		log.Print(err)
		return err
	}
	return az.db.DeleteRole(ctx, role)

}

func (az *authService) GetAssignRoleByUserId(ctx context.Context, role models.UserRole) (models.UserRole, error) {
	err := validation.ValidateStruct(&role,
		validation.Field(&role.UserId, validation.NotIn(uuid.Nil)),
	)
	if err != nil {
		log.Print(err)
		return models.UserRole{}, err
	}
	return az.db.GetAssignRoleByUserId(ctx, role)

}
