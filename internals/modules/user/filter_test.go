package user

import (
	"context"
	"errors"
	"testing"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	mockdb "github.com/alazarbeyeneazu/Simple-Gym-management-system/mocks/db"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetUserById(t *testing.T) {
	respondUser := models.User{
		ID:          uuid.New(),
		FirstName:   utils.RandomUserName(),
		LastName:    utils.RandomUserName(),
		PhoneNumber: utils.RandomePhoneNumber(),
		Password:    utils.RandomPassword(),
	}
	ctr := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctr)
	defer ctr.Finish()

	appuser := InitService(db)
	testCase := []struct {
		name    string
		user    models.User
		checker func(t *testing.T, user models.User, err error)
	}{
		{
			name: "ok",
			user: respondUser,
			checker: func(t *testing.T, user models.User, err error) {
				require.NoError(t, err)

			},
		}, {
			name: "not found",
			user: models.User{ID: uuid.New()},
			checker: func(t *testing.T, user models.User, err error) {
				require.EqualError(t, err, "user not found")

			},
		}, {
			name: "empty id",
			user: models.User{},
			checker: func(t *testing.T, user models.User, err error) {
				require.EqualError(t, err, "id can not be blank")

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				db.EXPECT().GetUserById(gomock.Any(), respondUser.ID).Return(respondUser, nil)
				user, err := appuser.GetUserById(context.Background(), tc.user)
				tc.checker(t, user, err)
			case "not found":
				db.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(models.User{}, errors.New("user not found"))
				user, err := appuser.GetUserById(context.Background(), tc.user)
				tc.checker(t, user, err)
			case "empty id":
				user, err := appuser.GetUserById(context.Background(), tc.user)
				tc.checker(t, user, err)

			}

		})
	}

}

//get users by first name test
func TestGetUsersByFirstName(t *testing.T) {
	firsname := utils.RandomUserName()
	var respondUser []models.User
	for i := 0; i < 10; i++ {
		resp := models.User{
			ID:          uuid.New(),
			FirstName:   firsname,
			LastName:    utils.RandomUserName(),
			PhoneNumber: utils.RandomePhoneNumber(),
			Password:    utils.RandomPassword(),
		}
		respondUser = append(respondUser, resp)
	}

	ctr := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctr)
	defer ctr.Finish()

	appuser := InitService(db)
	testCase := []struct {
		name    string
		user    models.User
		checker func(t *testing.T, user []models.User, err error)
	}{
		{
			name: "ok",
			user: models.User{FirstName: firsname},
			checker: func(t *testing.T, user []models.User, err error) {
				require.Equal(t, len(user), 10)
				require.NoError(t, err)

			},
		}, {
			name: "not found",
			user: models.User{FirstName: utils.RandomUserName()},
			checker: func(t *testing.T, user []models.User, err error) {
				require.EqualError(t, err, "user not found")

			},
		}, {
			name: "empty firstname",
			user: models.User{},
			checker: func(t *testing.T, user []models.User, err error) {
				require.EqualError(t, err, "first name cannot be blank")

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				db.EXPECT().GetUserByFirstName(gomock.Any(), gomock.Any()).Return(respondUser, nil)
				user, err := appuser.GetUsersByFirstName(context.Background(), tc.user)
				tc.checker(t, user, err)
			case "not found":
				db.EXPECT().GetUserByFirstName(gomock.Any(), gomock.Any()).Return([]models.User{}, errors.New("user not found"))
				user, err := appuser.GetUsersByFirstName(context.Background(), tc.user)
				tc.checker(t, user, err)
			case "empty firstname":
				user, err := appuser.GetUsersByFirstName(context.Background(), tc.user)
				tc.checker(t, user, err)

			}

		})
	}

}
