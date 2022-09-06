package persistant

import (
	"context"
	"log"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type dbAdapter struct {
	db *gorm.DB
}

func Init() ports.DBPort {

	db, err := gorm.Open(sqlite.Open("gym.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})

	return &dbAdapter{db: db}
}
func (a *dbAdapter) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	if len(user.PhoneNumber) == 10 {
		user.PhoneNumber = "+251" + user.PhoneNumber[1:]
	}
	err := utils.ValideteUser(user)
	if err != nil {
		return models.User{}, err
	}
	user.ID = uuid.New()
	result := a.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, result.Error

}
func (a *dbAdapter) Delete(ctx context.Context, user models.User) error {

	return nil
}

func (a *dbAdapter) Update(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (a *dbAdapter) Getuser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (a *dbAdapter) GetUsers(ctx context.Context) ([]models.User, error) {
	return []models.User{}, nil
}
