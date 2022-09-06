package ports

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
)

type DBPort interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	Delete(ctx context.Context, user models.User) error
	Update(ctx context.Context, user models.User) (models.User, error)
	Getuser(ctx context.Context, user models.User) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}
