package ports

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/google/uuid"
)

//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package mockdb -destination ../mocks/db/mockdb.go github.com/alazarbeyeneazu/Simple-Gym-management-system/ports DBPort
type DBPort interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, newUser, user models.User) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserByFirstName(ctx context.Context, firstname string) ([]models.User, error)
	GetUserByLastName(ctx context.Context, lastname string) ([]models.User, error)
	GetUseByPhoneNumber(ctx context.Context, phonenumber string) (models.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (models.User, error)
}
