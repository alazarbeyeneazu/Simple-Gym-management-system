package persistant

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/google/uuid"
)

func (a *dbAdapter) CheckInUser(ctx context.Context, user models.Checkins) (models.Checkins, error) {

	user.ID = uuid.New()

	result := a.db.Create(&user)
	if result.Error != nil {
		return models.Checkins{}, result.Error
	}
	return user, nil
}
func (a *dbAdapter) GetCheckedInByUserId(ctx context.Context, user models.Checkins) ([]models.Checkins, error) {
	var users []models.Checkins
	result := a.db.Where("user_id = ?", user.UserId).Find(&users)
	if result.Error != nil {
		return []models.Checkins{}, result.Error
	}

	return users, nil
}

func (a *dbAdapter) GetAllCheckIns(ctx context.Context) ([]models.Checkins, error) {
	var checkins []models.Checkins
	result := a.db.Find(&checkins)
	if result.Error != nil {
		return []models.Checkins{}, result.Error
	}
	return checkins, nil
}
