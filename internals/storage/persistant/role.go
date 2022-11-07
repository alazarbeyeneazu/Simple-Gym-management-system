package persistant

import (
	"context"
	"fmt"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/google/uuid"
)

func (p *dbAdapter) CreateRole(ctx context.Context, role models.Role) (models.Role, error) {
	role.ID = uuid.New()
	result := p.db.Create(&role)
	if result.Error != nil {
		return models.Role{}, result.Error
	}
	return role, nil
}

func (p *dbAdapter) GetAllRoles(ctx context.Context) ([]models.Role, error) {

	var roles []models.Role
	rolesMap := make(map[string]models.Role)
	finalRole := []models.Role{}
	result := p.db.Find(&roles)
	if result.Error != nil {
		return []models.Role{}, result.Error
	}
	for _, role := range roles {
		rolesMap[role.RoleName] = role
	}
	for _, rolem := range rolesMap {
		finalRole = append(finalRole, rolem)
	}
	return finalRole, nil
}

func (p *dbAdapter) GetRolesByName(ctx context.Context, role models.Role) ([]models.Role, error) {
	var roles []models.Role
	result := p.db.Where("role_name = ?", role.RoleName).Find(&roles)
	if result.Error != nil {
		return []models.Role{}, result.Error
	}

	return roles, nil

}

func (p *dbAdapter) DeleteRole(ctx context.Context, role models.Role) error {

	result := p.db.Where("role_name = ?", role.RoleName).Delete(&role)
	if result.Error != nil {
		return result.Error

	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (p *dbAdapter) AssignRole(ctx context.Context, role models.UserRole) (models.UserRole, error) {
	role.ID = uuid.New()
	result := p.db.Create(&role)
	if result.Error != nil {
		return models.UserRole{}, result.Error
	}
	return role, nil
}

func (p *dbAdapter) GetAssignRoleByUserId(ctx context.Context, role models.UserRole) (models.UserRole, error) {
	var roles models.UserRole
	result := p.db.Where("user_id = ?", role.UserId).Find(&roles)
	if result.Error != nil {
		return models.UserRole{}, result.Error
	}

	return roles, nil
}
