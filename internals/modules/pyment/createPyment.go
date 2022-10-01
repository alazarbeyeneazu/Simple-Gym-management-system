package pyment

import (
	"context"
	"errors"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

func (py *pymentService) CreatePyment(ctx context.Context, pyment models.PymentType) (models.PymentType, error) {
	err := validation.ValidateStruct(&pyment,
		validation.Field(&pyment.CreatedByFirstName, validation.Required),
		validation.Field(&pyment.CreatedByLastName, validation.Required),
		validation.Field(&pyment.NumberOfDays, validation.Required, validation.Min(1)),
		validation.Field(&pyment.Payment, validation.Required),
		validation.Field(&pyment.PymentType, validation.Required),
	)
	if err != nil {
		return models.PymentType{}, err
	}
	pymnt, err := py.db.CreatePymentType(ctx, pyment)
	if err != err {
		return models.PymentType{}, err
	}
	return pymnt, nil

}
func (py *pymentService) DeletePyment(ctx context.Context, pyment models.PymentType) error {
	err := validation.ValidateStruct(&pyment,
		validation.Field(&pyment.ID, validation.Required),
	)
	if err != nil {
		return err
	}
	if pyment.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return errors.New("pyment id can not be blank")
	}
	err = py.db.DeletePyment(ctx, pyment)
	if err != err {
		return err
	}
	return nil

}

func (py *pymentService) GetAllPyments(ctx context.Context) ([]models.PymentType, error) {

	pyments, err := py.db.GetAllPyments(ctx)
	if err != err {
		return []models.PymentType{}, err
	}
	return pyments, nil

}

func (py *pymentService) GetPymentById(ctx context.Context, pyment models.PymentType) (models.PymentType, error) {
	err := validation.ValidateStruct(&pyment,
		validation.Field(&pyment.ID, validation.Required),
	)
	if err != nil {
		return models.PymentType{}, err
	}
	if pyment.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.PymentType{}, errors.New("pyment id can not be blank")
	}
	pyment, err = py.db.GetPymentById(ctx, pyment)
	if err != err {
		return models.PymentType{}, err
	}
	return pyment, nil

}
func (py *pymentService) UpdatePyment(ctx context.Context, pyment models.PymentType) (models.PymentType, error) {
	err := validation.ValidateStruct(&pyment,
		validation.Field(&pyment.ID, validation.Required),
	)
	if err != nil {
		return models.PymentType{}, err
	}
	if pyment.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.PymentType{}, errors.New("pyment id can not be blank")
	}
	pyment, err = py.db.UpdatePyment(ctx, pyment)
	if err != err {
		return models.PymentType{}, err
	}
	return pyment, nil

}
