package persistant

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func (a *dbAdapter) CreateAdmin(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error) {
	err := validation.ValidateStruct(
		&admin,
		validation.Field(&admin.UserId, validation.Required, validation.NotIn(uuid.Nil)),
	)
	if err != nil {
		return models.AdminUsers{}, err
	}
	admin.ID = uuid.New()

	result := a.db.Create(&admin)
	if result.Error != nil {
		return models.AdminUsers{}, result.Error
	}
	return admin, nil
}
func (a *dbAdapter) GetAllAdmins() ([]models.AdminUsers, error) {
	var admins []models.AdminUsers
	result := a.db.Find(&admins)
	if result.Error != nil {
		return []models.AdminUsers{}, result.Error
	}
	return admins, nil
}
func (a *dbAdapter) GetAdminByUserId(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error) {
	var adminUser models.AdminUsers
	result := a.db.Where("user_id = ? ", admin.UserId).First(&adminUser)
	if result.Error != nil {
		return models.AdminUsers{}, result.Error
	}
	return adminUser, nil

}
func (a *dbAdapter) DeleteAdmin(ctx context.Context, admin models.AdminUsers) error {
	result := a.db.Where("id = ?", admin.ID).Delete(&admin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *dbAdapter) GetAdminById(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error) {
	var adminUser models.AdminUsers
	result := a.db.Where("id = ? ", admin.ID).First(&adminUser)
	if result.Error != nil {
		return models.AdminUsers{}, result.Error
	}
	return adminUser, nil

}
