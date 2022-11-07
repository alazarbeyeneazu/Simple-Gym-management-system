package authz

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func (az *authService) CreatePermission(ctx context.Context, permission models.Permission) (models.Permission, error) {
	err := validation.ValidateStruct(&permission,
		validation.Field(&permission.Path, validation.Required),
		validation.Field(&permission.Action, validation.Required),
	)
	if err != nil {
		return models.Permission{}, err
	}
	return az.db.CreatePermissions(ctx, permission)
}

func (az *authService) GetPermissionById(ctx context.Context, permission models.Permission) (models.Permission, error) {

	if err := validation.Validate(&permission.ID, validation.NotIn(uuid.Nil).Error("permission id required")); err != nil {
		return models.Permission{}, err
	}
	return az.db.GetPermissionById(ctx, permission)
}
