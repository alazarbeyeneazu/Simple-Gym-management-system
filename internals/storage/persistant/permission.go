package persistant

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/google/uuid"
)

func (p *dbAdapter) CreatePermissions(ctx context.Context, permission models.Permission) (models.Permission, error) {
	permission.ID = uuid.New()
	result := p.db.Create(&permission)
	if result.Error != nil {
		return models.Permission{}, result.Error
	}
	return permission, nil
}
func (p *dbAdapter) GetAllPermissions(ctx context.Context) ([]models.Permission, error) {
	var permissions []models.Permission
	result := p.db.Find(&permissions)
	if result.Error != nil {
		return []models.Permission{}, result.Error
	}
	return permissions, nil
}

func (p *dbAdapter) GetPermissionById(ctx context.Context, permission models.Permission) (models.Permission, error) {
	var perm models.Permission
	result := p.db.Where("id = ? ", permission.ID).First(&perm)
	if result.Error != nil {
		return models.Permission{}, result.Error
	}
	return perm, nil

}
