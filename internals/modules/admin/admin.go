package admin

import (
	"context"
	"log"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	encription "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/encryption"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
)

type AdminService interface {
	RegisterAdmin(ctx context.Context, admin models.CreateAdminRequest) (models.AdminUsers, error)
	GetAllAdmins(ctx context.Context) ([]models.AdminUsers, error)
	DeleteAdmin(ctx context.Context, admin models.AdminUsers) error
	GetAdminByUserId(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error)
	GetAdminById(ctx context.Context, admin models.AdminUsers) (models.AdminUsers, error)
	InitializeSuperAdmin(ctx context.Context, phoneNumber, password string)
}

type adminService struct {
	db ports.DBPort
}

func InitService(db ports.DBPort) AdminService {

	return &adminService{
		db: db,
	}
}
func (ad *adminService) InitializeSuperAdmin(ctx context.Context, phoneNumber, password string) {
	//create supper admin
	Password, err := encription.GenerateHashedPassword(password)
	if err != nil {
		log.Println(err)
		return
	}
	user, err := ad.db.CreateUser(ctx, models.User{
		FirstName:   "SuperAdmin",
		LastName:    "SupperAdmin",
		PhoneNumber: phoneNumber,
		Password:    Password,
	})
	if err != nil {
		log.Println(err)
		return
	}

	//create supper addmin Role
	permissions, err := ad.db.GetAllPermissions(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	for _, permission := range permissions {
		ad.db.CreateRole(ctx, models.Role{
			RoleName:     "systemSupperadminUser",
			PermissionID: permission.ID,
		})
	}
	//assign role to the user
	_, err = ad.db.AssignRole(ctx, models.UserRole{
		UserId:   user.ID,
		RoleName: "systemSupperadminUser",
	})
	if err != nil {
		log.Println(err)
		return
	}
}
