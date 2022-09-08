package user

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
)

type UserService interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
}
type userService struct {
	db ports.DBPort
}

func InitService(db ports.DBPort) UserService {
	return &userService{
		db: db,
	}
}
