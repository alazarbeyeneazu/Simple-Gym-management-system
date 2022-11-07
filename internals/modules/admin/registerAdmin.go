package admin

import (
	"context"
	"errors"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	encription "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/encryption"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func (a *adminService) RegisterAdmin(ctx context.Context, admin models.CreateAdminRequest) (models.AdminUsers, error) {
	adminUser := models.AdminUsers{}
	if len(admin.PhoneNumber) == 10 {
		admin.PhoneNumber = "+251" + admin.PhoneNumber[1:]
	}

	err := validation.ValidateStruct(&admin,
		validation.Field(&admin.FirstName, validation.Required.Error("first name required")),
		validation.Field(&admin.LastName, validation.Required.Error("last name required")),
		validation.Field(&admin.PhoneNumber, validation.Required, validation.Length(13, 13).Error("invalid phone number")),
		validation.Field(&admin.Role, validation.Required.Error("role required")),
	)
	if err != nil {
		return models.AdminUsers{}, err
	}
	user, _ := a.db.GetUseByPhoneNumber(ctx, admin.PhoneNumber)

	if user.ID == uuid.Nil {
		passwordG, err := encription.GenerateHashedPassword(admin.Password)

		if err != nil {
			return models.AdminUsers{}, err
		}
		result, err := a.db.CreateUser(ctx, models.User{
			FirstName:   admin.FirstName,
			LastName:    admin.LastName,
			PhoneNumber: admin.PhoneNumber,
			Password:    passwordG,
		})
		if err != nil {
			return models.AdminUsers{}, err
		}
		adminUser.UserId = result.ID

	} else {
		adminUser.UserId = user.ID
	}
	adminResult, _ := a.db.GetAdminByUserId(ctx, models.AdminUsers{UserId: adminUser.UserId})
	if adminResult.ID != uuid.Nil {
		return models.AdminUsers{}, errors.New("user already exist")
	}

	adminReturn, err := a.db.CreateAdmin(ctx, adminUser)
	if adminResult.ID != uuid.Nil {
		return models.AdminUsers{}, err
	}
	a.db.AssignRole(ctx, models.UserRole{UserId: adminUser.UserId, RoleName: admin.Role})
	return adminReturn, nil
}
