package user

import (
	"context"
	"fmt"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (us *userService) GetUserById(ctx context.Context, user models.User) (models.User, error) {
	err := validation.Validate(&user.ID, validation.Required, is.UUID)
	if err != nil {
		return models.User{}, fmt.Errorf("id %s", err.Error())
	}
	if user.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return models.User{}, fmt.Errorf("id can not be blank")
	}
	returneduser, err := us.db.GetUserById(ctx, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return returneduser, nil
}

//get user by first name
func (us *userService) GetUsersByFirstName(ctx context.Context, user models.User) ([]models.User, error) {
	err := validation.Validate(user.FirstName, validation.Required)
	if err != nil {
		return []models.User{}, fmt.Errorf("first name %s", err.Error())
	}
	users, err := us.db.GetUserByFirstName(ctx, user.FirstName)
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}
