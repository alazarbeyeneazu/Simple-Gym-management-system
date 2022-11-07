package persistant

import (
	"context"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
)

func (a *dbAdapter) CreateReport(ctx context.Context, report models.ReportResponse) error {
	result := a.db.Create(&report)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
func (a *dbAdapter) GetAllReports(ctx context.Context) ([]models.ReportResponse, error) {

	var response []models.ReportResponse
	result := a.db.Find(&response)
	if result.Error != nil {
		return []models.ReportResponse{}, result.Error
	}
	return response, nil
}
