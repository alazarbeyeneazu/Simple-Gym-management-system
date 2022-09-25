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

func TestDeleteUser(t *testing.T) {
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
		checker func(t *testing.T, err error)
	}{
		{
			name: "ok",
			user: respondUser,
			checker: func(t *testing.T, err error) {
				require.NoError(t, err)

			},
		}, {
			name: "not found",
			user: models.User{ID: uuid.New()},
			checker: func(t *testing.T, err error) {
				require.EqualError(t, err, "user not found")

			},
		}, {
			name: "empty id",
			user: models.User{},
			checker: func(t *testing.T, err error) {
				require.EqualError(t, err, "id can not be blank")

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				db.EXPECT().DeleteUser(gomock.Any(), respondUser).Return(nil)
				err := appuser.DeleteUser(context.Background(), tc.user)
				tc.checker(t, err)
			case "not found":
				db.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Return(errors.New("user not found"))
				err := appuser.DeleteUser(context.Background(), tc.user)
				tc.checker(t, err)
			case "empty id":
				err := appuser.DeleteUser(context.Background(), tc.user)
				tc.checker(t, err)

			}

		})
	}

}
