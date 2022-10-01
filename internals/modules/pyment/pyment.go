package pyment

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
)

type PymentService interface {
	CreatePyment(ctx context.Context, pyment models.PymentType) (models.PymentType, error)
	DeletePyment(ctx context.Context, pyment models.PymentType) error
	GetAllPyments(ctx context.Context) ([]models.PymentType, error)
	GetPymentById(ctx context.Context, pyment models.PymentType) (models.PymentType, error)
	UpdatePyment(ctx context.Context, pyment models.PymentType) (models.PymentType, error)
}
type pymentService struct {
	db ports.DBPort
}

func InitService(db ports.DBPort) PymentService {
	return &pymentService{
		db: db,
	}
}
