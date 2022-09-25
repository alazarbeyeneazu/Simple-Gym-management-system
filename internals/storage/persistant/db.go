package persistant

import (
	"context"
	"fmt"
	"log"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/ports"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
func (a *dbAdapter) DeleteUser(ctx context.Context, user models.User) error {
	err := validation.Validate(&user.ID, validation.Required, is.UUID)
	if err != nil {
		return fmt.Errorf("id %s ", err.Error())
	}

	result := a.db.Where("id = ?", user.ID).Delete(&user)
	if result.Error != nil {
		return err

	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (a *dbAdapter) UpdateUser(ctx context.Context, newUser, user models.User) (models.User, error) {
	err := validation.Validate(&user.ID, validation.Required, is.UUID)
	var updatedUser models.User
	result := a.db.Where("id = ?", user.ID).Find(&updatedUser)
	if result.RowsAffected == 0 {
		return models.User{}, fmt.Errorf("can not find the user with is id %v", user.ID)
	}
	if err != nil {
		return user, fmt.Errorf("user id %s", err.Error())
	}
	if len(newUser.FirstName) > 0 {
		err := validation.Validate(&user.FirstName, validation.Length(2, 100))
		if err != nil {
			return user, fmt.Errorf("first_name %s", err.Error())
		}
		a.db.Exec("UPDATE users set first_name = ?", newUser.FirstName)

	}
	if len(newUser.LastName) > 0 {
		err := validation.Validate(&user.LastName, validation.Length(2, 100))
		if err != nil {
			return user, fmt.Errorf("last_name %s", err.Error())
		}
		a.db.Exec("UPDATE users set last_name = ?", newUser.LastName)
	}
	if len(newUser.PhoneNumber) > 0 {
		if len(newUser.PhoneNumber) == 10 {
			newUser.PhoneNumber = "+251" + newUser.PhoneNumber[1:]
		}
		err := validation.Validate(&user.PhoneNumber, validation.Length(13, 13))
		if err != nil {
			return user, fmt.Errorf("phone_number %s", err.Error())
		}
		a.db.Exec("UPDATE users set phone_number = ?", newUser.PhoneNumber)

	}
	if len(newUser.Password) > 0 {

		err := validation.Validate(&user.Password, validation.Length(8, 100))
		if err != nil {
			return user, fmt.Errorf("phone_number %s", err.Error())
		}
		a.db.Exec("UPDATE users set password = ?", newUser.Password)

	}
	result = a.db.Where("id = ?", user.ID).First(&updatedUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return updatedUser, nil
}

func (a *dbAdapter) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	result := a.db.Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}
func (a *dbAdapter) GetUserByFirstName(ctx context.Context, firstname string) ([]models.User, error) {
	var users []models.User
	err := validation.Validate(&firstname, validation.Required)
	if err != nil {
		return []models.User{}, err
	}
	result := a.db.Where("first_name = ?", firstname).Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return users, nil
}
func (a *dbAdapter) GetUserByLastName(ctx context.Context, lastname string) ([]models.User, error) {

	var users []models.User
	err := validation.Validate(&lastname, validation.Required)
	if err != nil {
		return []models.User{}, err
	}
	result := a.db.Where("last_name = ?", lastname).Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return users, nil
}
func (a *dbAdapter) GetUseByPhoneNumber(ctx context.Context, phonenumber string) (models.User, error) {
	var user models.User
	if len(phonenumber) == 10 {
		phonenumber = "+251" + phonenumber[1:]
	}
	log.Println(phonenumber)
	err := validation.Validate(&phonenumber, validation.Required, validation.Length(13, 13))
	if err != nil {
		return models.User{}, fmt.Errorf("phone number %s", err.Error())
	}

	result := a.db.Where("phone_number = ?", phonenumber).First(&user)
	if result.Error != nil {
		return models.User{}, err
	}
	return user, nil
}

func (a *dbAdapter) GetUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	var user models.User
	err := validation.Validate(id, validation.Required, is.UUID)
	if err != nil {
		return models.User{}, fmt.Errorf("phone number %s", err.Error())
	}

	result := a.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, err
	}
	return user, nil
}
