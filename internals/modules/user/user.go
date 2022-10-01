package user

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
)

type UserService interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
	GetUserById(ctx context.Context, user models.User) (models.User, error)
	GetUsersByFirstName(ctx context.Context, user models.User) ([]models.User, error)
	GetUserByLastName(ctx context.Context, user models.User) ([]models.User, error)
	GetUserByPhoneNumber(ctx context.Context, user models.User) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	DeleteUser(ctx context.Context, user models.User) error
}
type userService struct {
	db ports.DBPort
}

func InitService(db ports.DBPort) UserService {
	return &userService{
		db: db,
	}
}
