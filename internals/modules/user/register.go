package user

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	encription "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/encryption"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
)

func (us *userService) RegisterUser(ctx context.Context, user models.User) (models.User, error) {
	var err error
	if len(user.PhoneNumber) == 10 {
		user.PhoneNumber = "+251" + user.PhoneNumber[1:]
	}

	user.Password, err = encription.GenerateHashedPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	err = utils.ValideteUser(user)
	if err != nil {
		return models.User{}, err
	}
	newuser, err := us.db.CreateUser(ctx, user)
	if err != nil {
		return user, err
	}
	return newuser, nil

}
