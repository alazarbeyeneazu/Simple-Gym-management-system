package repor

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
)

type ReportService interface {
	CreateReport(ctx context.Context, report models.ReportResponse) error
	GetAllReports(ctx context.Context) ([]models.ReportResponse, error)
}
type reportService struct {
	db ports.DBPort
}

func InitService(db ports.DBPort) ReportService {
	return &reportService{
		db: db,
	}
}

func (r *reportService) CreateReport(ctx context.Context, report models.ReportResponse) error {
	return r.db.CreateReport(ctx, report)

}

func (r *reportService) GetAllReports(ctx context.Context) ([]models.ReportResponse, error) {
	return r.db.GetAllReports(ctx)
}
