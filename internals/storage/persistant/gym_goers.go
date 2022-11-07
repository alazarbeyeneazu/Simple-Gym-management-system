package persistant

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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

func (a *dbAdapter) UpdateGymGoer(ctx context.Context, newGymGoer models.Gym_goers) (models.Gym_goers, error) {
	err := validation.Validate(&newGymGoer.ID, validation.Required, is.UUID)
	var UpdateGymUser models.Gym_goers

	result := a.db.Where("user_id = ?", newGymGoer.UserId).First(&UpdateGymUser)

	if result.RowsAffected == 0 {
		return models.Gym_goers{}, fmt.Errorf("can not find the user with is id %v", newGymGoer.ID)
	}
	if err != nil {
		return newGymGoer, fmt.Errorf("user id %s", err.Error())
	}
	if len(newGymGoer.CreatedByFirstName) > 0 {
		err := validation.Validate(&newGymGoer.CreatedByFirstName, validation.Length(2, 100))
		if err != nil {
			return newGymGoer, fmt.Errorf("creator first_name %s", err.Error())
		}
		a.db.Model(&models.Gym_goers{}).Where("user_id = ? ", newGymGoer.UserId).Update("created_by_first_name", newGymGoer.CreatedByFirstName)

	}
	if len(newGymGoer.CreatedByLastName) > 0 {
		err := validation.Validate(&newGymGoer.CreatedByFirstName, validation.Length(2, 100))
		if err != nil {
			return newGymGoer, fmt.Errorf("creator last_name %s", err.Error())
		}
		a.db.Model(&models.Gym_goers{}).Where("user_id = ? ", newGymGoer.UserId).Update("created_by_last_name", newGymGoer.CreatedByLastName)

	}
	if len(newGymGoer.CreatedByPhoneNumber) > 0 {
		err := validation.Validate(&newGymGoer.CreatedByFirstName, validation.Length(2, 100))
		if err != nil {
			return newGymGoer, fmt.Errorf("creator phone_number %s", err.Error())
		}
		a.db.Model(&models.Gym_goers{}).Where("user_id = ? ", newGymGoer.UserId).Update("created_by_phone_number", newGymGoer.CreatedByPhoneNumber)

	}

	a.db.Model(&models.Gym_goers{}).Where("user_id = ? ", newGymGoer.UserId).Update("start_date", newGymGoer.StartDate)
	a.db.Model(&models.Gym_goers{}).Where("user_id = ? ", newGymGoer.UserId).Update("end_date", newGymGoer.EndDate)
	a.db.Model(&models.Gym_goers{}).Where("user_id = ? ", newGymGoer.UserId).Update("paid_by", newGymGoer.PaidBy)

	return newGymGoer, nil
}
