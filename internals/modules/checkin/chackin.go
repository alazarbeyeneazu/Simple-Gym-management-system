package checkin

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type CheckingService interface {
	CheckingUser(ctx context.Context, user models.Checkins) (models.CheckinResponse, error)
	GetCheckedInByUserId(ctx context.Context, user models.Checkins) ([]models.Checkins, error)
	GetAllCheckIns(ctx context.Context) ([]models.Checkins, error)
}
type checkingService struct {
	db ports.DBPort
}

func InitService(db ports.DBPort) CheckingService {
	return &checkingService{
		db: db,
	}
}

func (az *checkingService) CheckingUser(ctx context.Context, user models.Checkins) (models.CheckinResponse, error) {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.UserId, validation.NotIn(uuid.Nil).Error("User Id Required")),
	)
	if err != nil {
		return models.CheckinResponse{}, err
	}
	user.CheckedInDate = time.Now()
	checkins, err := az.db.GetCheckedInByUserId(ctx, models.Checkins{UserId: user.UserId})
	if err == nil {
		for _, checkin := range checkins {
			if checkin.CheckedInDate.Add(time.Hour * 6).After(time.Now()) {

				return models.CheckinResponse{
					IsChackedIn: fmt.Sprintf("You are aready checked before %v Hour", time.Now().Hour()-checkin.CheckedInDate.Hour()),
				}, errors.New("already checkedIn")
			}
		}
	}
	checked, err := az.db.CheckInUser(ctx, user)
	if err != nil {
		return models.CheckinResponse{}, err
	}
	return models.CheckinResponse{CheckedInDate: checked.CheckedInDate}, nil
}

func (az *checkingService) GetCheckedInByUserId(ctx context.Context, user models.Checkins) ([]models.Checkins, error) {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.UserId, validation.NotIn(uuid.Nil).Error("User Id Required")),
	)
	if err != nil {
		return []models.Checkins{}, err
	}

	checked, err := az.db.GetCheckedInByUserId(ctx, models.Checkins{UserId: user.UserId})
	if err != nil {
		return []models.Checkins{}, err
	}
	return checked, nil
}

func (az *checkingService) GetAllCheckIns(ctx context.Context) ([]models.Checkins, error) {
	return az.db.GetAllCheckIns(ctx)
}
