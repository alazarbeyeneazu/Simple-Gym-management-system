package persistant

import (
	"context"
	"errors"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func (a *dbAdapter) CreateGymGoers(ctx context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error) {
	err := validation.ValidateStruct(
		&gym_goers,
		validation.Field(&gym_goers.UserId, validation.Required),
		validation.Field(&gym_goers.CreatedByFirstName, validation.Required),
		validation.Field(&gym_goers.CreatedByLastName, validation.Required),
		validation.Field(&gym_goers.CreatedByPhoneNumber, validation.Required),
		validation.Field(&gym_goers.StartDate, validation.Required),
		validation.Field(&gym_goers.EndDate, validation.Required),
		validation.Field(&gym_goers.PaidBy, validation.Required),
	)
	if err != nil {
		return models.Gym_goers{}, err
	}
	if gym_goers.UserId == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.Gym_goers{}, errors.New("user id can not be empty")
	}
	if gym_goers.StartDate.AddDate(0, 0, 1).Before(time.Now()) {

		return models.Gym_goers{}, errors.New("start date should be equal or after today")
	}
	if gym_goers.EndDate.Before(gym_goers.StartDate) {
		return models.Gym_goers{}, errors.New("end date can not be less than start date")
	}
	gym_goers.ID = uuid.New()
	result := a.db.Create(&gym_goers)
	if result.Error != nil {
		return models.Gym_goers{}, result.Error
	}
	return gym_goers, nil

}
func (a *dbAdapter) DeleteGymGoers(ctx context.Context, gymGors models.Gym_goers) error {
	err := validation.Validate(&gymGors.ID, validation.Required)
	if err != nil {
		return err
	}
	if gymGors.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return errors.New("gym_goer's id can not be blank")
	}
	result := a.db.Where("id = ?", gymGors.ID).Delete(&gymGors)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (a *dbAdapter) GetAllGymGoers(cxt context.Context) ([]models.Gym_goers, error) {
	var gym_goers []models.Gym_goers
	result := a.db.Find(&gym_goers)
	if result.Error != nil {
		return []models.Gym_goers{}, result.Error
	}
	return gym_goers, nil
}

func (a *dbAdapter) GetGYmGorsById(cxt context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error) {
	var gymgoer models.Gym_goers
	err := validation.Validate(&gym_goers.ID, validation.Required)
	if err != nil {
		return models.Gym_goers{}, err
	}
	if gym_goers.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.Gym_goers{}, errors.New("empy gym_goer's id")
	}
	result := a.db.Where("id = ? ", gym_goers.ID).First(&gymgoer)
	if result.Error != nil {
		return models.Gym_goers{}, result.Error
	}
	return gymgoer, nil
}
func (a *dbAdapter) GetGymGoerByUserId(ctx context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error) {
	var gymgoer models.Gym_goers
	err := validation.Validate(&gym_goers.UserId, validation.Required)
	if err != nil {
		return models.Gym_goers{}, err
	}
	if gym_goers.UserId == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.Gym_goers{}, errors.New("empy user id")
	}
	result := a.db.Where("user_id = ? ", gym_goers.UserId).First(&gymgoer)
	if result.Error != nil {
		return models.Gym_goers{}, result.Error
	}
	return gymgoer, nil
}

func (a *dbAdapter) GetGymGoerByCreatedByFirstName(ctx context.Context, creator models.User) ([]models.Gym_goers, error) {
	var gymgoer []models.Gym_goers
	err := validation.Validate(&creator.FirstName, validation.Required)
	if err != nil {
		return []models.Gym_goers{}, err
	}

	result := a.db.Where("created_by_first_name = ? ", creator.FirstName).Find(&gymgoer)
	if result.Error != nil {
		return []models.Gym_goers{}, result.Error
	}
	return gymgoer, nil
}

func (a *dbAdapter) GetGymGoerByCreatedByLastName(ctx context.Context, creator models.User) ([]models.Gym_goers, error) {
	var gymgoer []models.Gym_goers
	err := validation.Validate(&creator.LastName, validation.Required)
	if err != nil {
		return []models.Gym_goers{}, err
	}

	result := a.db.Where("created_by_last_name = ? ", creator.LastName).Find(&gymgoer)
	if result.Error != nil {
		return []models.Gym_goers{}, result.Error
	}
	return gymgoer, nil
}

func (a *dbAdapter) GetGymGoerByCreatedByPhoneNumber(ctx context.Context, creator models.User) ([]models.Gym_goers, error) {
	var gymgoer []models.Gym_goers
	err := validation.Validate(&creator.PhoneNumber, validation.Required)
	if err != nil {
		return []models.Gym_goers{}, err
	}

	result := a.db.Where("created_by_phone_number = ? ", creator.PhoneNumber).Find(&gymgoer)
	if result.Error != nil {
		return []models.Gym_goers{}, result.Error
	}
	return gymgoer, nil
}

func (a *dbAdapter) GetGymGoerByPaidBy(ctx context.Context, gym_goers models.Gym_goers) ([]models.Gym_goers, error) {
	var gymgoer []models.Gym_goers
	err := validation.Validate(&gym_goers.PaidBy, validation.Required)
	if err != nil {
		return []models.Gym_goers{}, err
	}

	result := a.db.Where("paid_by = ? ", gym_goers.PaidBy).Find(&gymgoer)
	if result.Error != nil {
		return []models.Gym_goers{}, result.Error
	}
	return gymgoer, nil
}
