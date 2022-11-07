package user

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
)

func (us *userService) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	var err error
	if len(user.PhoneNumber) == 10 {
		user.PhoneNumber = "+251" + user.PhoneNumber[1:]
	}

	newuser, err := us.db.UpdateUser(ctx, user, user)
	if err != nil {
		return user, err
	}
	return newuser, nil
}
