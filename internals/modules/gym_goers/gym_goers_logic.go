package gymgoers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

func (gm *gymGoersService) RegisterGymGoer(ctx context.Context, gymgoer models.Gym_goers) (models.Gym_goers, error) {
	err := validation.ValidateStruct(
		&gymgoer,
		validation.Field(&gymgoer.CreatedByFirstName, validation.Required),
		validation.Field(&gymgoer.CreatedByLastName, validation.Required),
		validation.Field(&gymgoer.CreatedByPhoneNumber, validation.Required),
		validation.Field(&gymgoer.StartDate, validation.Required),
		validation.Field(&gymgoer.EndDate, validation.Required),
		validation.Field(&gymgoer.PaidBy, validation.Required),
	)
	if gymgoer.UserId == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.Gym_goers{}, errors.New("user id can not be empty")
	}
	if gymgoer.StartDate.AddDate(0, 0, 1).Before(time.Now()) {
		log.Println("requested Time", gymgoer.StartDate)
		return models.Gym_goers{}, errors.New("start date should be equal or after today")
	}
	if gymgoer.EndDate.Before(gymgoer.StartDate) {
		return models.Gym_goers{}, errors.New("end date can not be less than start date")
	}
	if err != nil {
		return models.Gym_goers{}, err
	}
	pymnt, err := gm.db.CreateGymGoers(ctx, gymgoer)
	if err != err {
		return models.Gym_goers{}, err
	}
	return pymnt, nil

}
func (gm *gymGoersService) DeleteGymGoers(ctx context.Context, gymgoer models.Gym_goers) error {
	err := validation.Validate(&gymgoer.ID, validation.Required)
	if err != nil {
		return err
	}
	if gymgoer.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return errors.New("gym_goer's id can not be blank")
	}
	result := gm.db.DeleteGymGoers(ctx, gymgoer)
	if result != nil {
		return result
	}
	return nil

}

func (gm *gymGoersService) GetAllGymGoers(ctx context.Context) ([]models.Gym_goers, error) {
	gymgoers, err := gm.db.GetAllGymGoers(ctx)
	if err != nil {
		return []models.Gym_goers{}, err
	}
	return gymgoers, nil
}
func (gm *gymGoersService) GetGYmGorsById(cxt context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error) {

	err := validation.Validate(&gym_goers.ID, validation.Required)
	if err != nil {
		return models.Gym_goers{}, err
	}
	if gym_goers.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.Gym_goers{}, errors.New("empy gym_goer's id")
	}
	result, err := gm.db.GetGYmGorsById(cxt, gym_goers)
	if err != nil {
		return models.Gym_goers{}, err
	}
	return result, nil
}
func (gm *gymGoersService) GetGymGoerByUserId(ctx context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error) {

	err := validation.Validate(&gym_goers.UserId, validation.Required)
	if err != nil {
		return models.Gym_goers{}, err
	}
	if gym_goers.UserId == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.Gym_goers{}, errors.New("empy user id")
	}
	result, err := gm.db.GetGymGoerByUserId(ctx, gym_goers)
	if err != nil {
		return models.Gym_goers{}, err
	}
	return result, nil
}

func (gm *gymGoersService) GetGymGoerByCreatedByFirstName(ctx context.Context, creator models.User) ([]models.Gym_goers, error) {

	err := validation.Validate(&creator.FirstName, validation.Required)
	if err != nil {
		return []models.Gym_goers{}, err
	}

	result, err := gm.db.GetGymGoerByCreatedByFirstName(ctx, creator)
	if err != nil {
		return []models.Gym_goers{}, err
	}
	return result, nil
}
func (gm *gymGoersService) GetGymGoerByCreatedByLastName(ctx context.Context, creator models.User) ([]models.Gym_goers, error) {

	err := validation.Validate(&creator.LastName, validation.Required)
	if err != nil {
		return []models.Gym_goers{}, err
	}

	result, err := gm.db.GetGymGoerByCreatedByLastName(ctx, creator)
	if err != nil {
		return []models.Gym_goers{}, err
	}
	return result, nil
}

func (gm *gymGoersService) GetGymGoerByCreatedByPhoneNumber(ctx context.Context, creator models.User) ([]models.Gym_goers, error) {

	err := validation.Validate(&creator.PhoneNumber, validation.Required)
	if err != nil {
		return []models.Gym_goers{}, err
	}

	result, err := gm.db.GetGymGoerByCreatedByPhoneNumber(ctx, creator)
	if err != nil {
		return []models.Gym_goers{}, err
	}
	return result, nil
}

func (gm *gymGoersService) GetGymGoerByPaidBy(ctx context.Context, gym_goers models.Gym_goers) ([]models.Gym_goers, error) {

	err := validation.Validate(&gym_goers.PaidBy, validation.Required)
	if err != nil {
		return []models.Gym_goers{}, err
	}

	result, err := gm.db.GetGymGoerByPaidBy(ctx, gym_goers)
	if err != nil {
		return []models.Gym_goers{}, err
	}
	return result, nil
}
