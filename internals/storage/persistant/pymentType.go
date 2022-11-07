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

func (a *dbAdapter) GetAllPyments(ctx context.Context) ([]models.PymentType, error) {
	var pyments []models.PymentType
	result := a.db.Find(&pyments)
	if result.Error != nil {
		return []models.PymentType{}, result.Error
	}
	return pyments, nil
}
func (a *dbAdapter) GetPymentById(ctx context.Context, pyment models.PymentType) (models.PymentType, error) {
	var pymt models.PymentType
	err := validation.Validate(&pyment.ID, validation.Required)
	if err != nil {
		return models.PymentType{}, err
	}
	if pyment.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.PymentType{}, errors.New("empy pyment id")
	}
	result := a.db.Where("id = ?", pyment.ID).Find(&pymt)
	if result.Error != nil {
		return models.PymentType{}, result.Error
	}
	return pymt, nil

}

func (a *dbAdapter) UpdatePyment(ctx context.Context, pyment models.PymentType) (models.PymentType, error) {

	err := validation.Validate(&pyment.ID, validation.Required)
	if err != nil {
		return models.PymentType{}, err
	}
	if pyment.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return models.PymentType{}, errors.New("empy pyment id")
	}

	if pyment.NumberOfDays > 0 {

		a.db.Model(models.PymentType{}).Where("id = ?", pyment.ID).Update("number_of_days", pyment.NumberOfDays)
	}
	if pyment.Payment != "" {

		a.db.Model(models.PymentType{}).Where("id = ?", pyment.ID).Update("payment", pyment.Payment)
	}
	if pyment.PymentType != "" {
		a.db.Model(models.PymentType{}).Where("id = ?", pyment.ID).Update("pyment_type", pyment.PymentType)

	}

	pym, err := a.GetPymentById(ctx, pyment)
	if err != nil {
		return models.PymentType{}, err
	}
	return pym, nil
}
