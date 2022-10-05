package gymgoers

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
)

type GymGoersService interface {
	RegisterGymGoer(ctx context.Context, gymgoer models.Gym_goers) (models.Gym_goers, error)
	DeleteGymGoers(ctx context.Context, gymgoer models.Gym_goers) error
	GetAllGymGoers(ctx context.Context) ([]models.Gym_goers, error)
	GetGYmGorsById(cxt context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error)
	GetGymGoerByUserId(ctx context.Context, gym_goers models.Gym_goers) (models.Gym_goers, error)
	GetGymGoerByCreatedByFirstName(ctx context.Context, creator models.User) ([]models.Gym_goers, error)
	GetGymGoerByCreatedByLastName(ctx context.Context, creator models.User) ([]models.Gym_goers, error)
	GetGymGoerByCreatedByPhoneNumber(ctx context.Context, creator models.User) ([]models.Gym_goers, error)
	GetGymGoerByPaidBy(ctx context.Context, gym_goers models.Gym_goers) ([]models.Gym_goers, error)
}
type gymGoersService struct {
	db ports.DBPort
}

func InitService(db ports.DBPort) GymGoersService {
	return &gymGoersService{
		db: db,
	}
}
