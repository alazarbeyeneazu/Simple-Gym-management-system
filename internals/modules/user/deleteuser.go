package user

import (
	"context"
	"fmt"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (us *userService) DeleteUser(ctx context.Context, user models.User) error {

	err := validation.Validate(&user.ID, validation.Required, is.UUID)
	if err != nil {
		return fmt.Errorf("id %s", err.Error())
	}
	if user.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return fmt.Errorf("id can not be blank")
	}
	err = us.db.DeleteUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
