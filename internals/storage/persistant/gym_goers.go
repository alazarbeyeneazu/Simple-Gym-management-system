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
	)
	if err != nil {
		return models.Gym_goers{}, err
	}
	if gym_goers.UserId == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.Gym_goers{}, errors.New("user id can not be empty")
	}
	if gym_goers.StartDate.Add(time.Hour * 1).Before(time.Now()) {
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
