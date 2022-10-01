package persistant

import (
	"context"
	"errors"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func (a *dbAdapter) CreatePymentType(ctx context.Context, pyment models.PymentType) (models.PymentType, error) {
	err := validation.ValidateStruct(&pyment,
		validation.Field(&pyment.CreatedByFirstName, validation.Required),
		validation.Field(&pyment.CreatedByLastName, validation.Required),
		validation.Field(&pyment.NumberOfDays, validation.Required, validation.Min(1)),
		validation.Field(&pyment.Payment, validation.Required),
		validation.Field(&pyment.PaidBy),
		validation.Field(&pyment.PymentType, validation.Required),
	)
	if err != nil {
		return models.PymentType{}, err
	}
	pyment.ID = uuid.New()
	result := a.db.Create(&pyment)
	if result.Error != nil {
		return models.PymentType{}, result.Error
	}
	return pyment, nil

}

func (a *dbAdapter) DeletePyment(ctx context.Context, pyment models.PymentType) error {
	err := validation.Validate(&pyment.ID, validation.Required)
	if err != nil {
		return err
	}
	if pyment.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return errors.New("empy pyment id")
	}
	result := a.db.Where("id = ?", pyment.ID).Delete(&pyment)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
