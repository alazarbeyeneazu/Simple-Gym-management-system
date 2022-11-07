package admin

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func (a *adminService) GetAllAdmins(ctx context.Context) ([]models.AdminUsers, error) {
	return a.db.GetAllAdmins()
}

func (a *adminService) GetAdminByUserId(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error) {
	err := validation.Validate(&admin.UserId, validation.Required, validation.NotIn(uuid.Nil))
	if err != nil {
		return models.AdminUsers{}, err
	}
	return a.db.GetAdminByUserId(ctx, admin)
}
func (a *adminService) GetAdminById(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error) {
	err := validation.Validate(&admin.ID, validation.Required, validation.NotIn(uuid.Nil))
	if err != nil {
		return models.AdminUsers{}, err
	}
	return a.db.GetAdminById(ctx, admin)
}
