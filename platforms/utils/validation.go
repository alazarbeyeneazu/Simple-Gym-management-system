package utils

import (
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValideteUser(user models.User) error {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.FirstName, validation.Required, validation.Length(2, 100)),
		validation.Field(&user.LastName, validation.Required, validation.Length(2, 100)),
		validation.Field(&user.PhoneNumber, validation.Length(13, 13), validation.Required),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 100)),
	)

	return err
}
