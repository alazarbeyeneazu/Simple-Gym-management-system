package admin

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func (a *adminService) DeleteAdmin(ctx context.Context, admin models.AdminUsers) error {
	err := validation.Validate(&admin.ID, validation.Required, validation.NotIn(uuid.Nil))
	if err != nil {
		return err
	}
	return a.db.DeleteAdmin(ctx, admin)
}
