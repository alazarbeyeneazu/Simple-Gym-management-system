package user

import (
	"context"
	"testing"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	mockdb "github.com/alazarbeyeneazu/Simple-Gym-management-system/mocks/db"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {

	respondUser := models.User{
		ID:          uuid.New(),
		FirstName:   utils.RandomUserName(),
		LastName:    utils.RandomUserName(),
		PhoneNumber: utils.RandomePhoneNumber(),
		Password:    utils.RandomPassword(),
	}
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	defer ctl.Finish()
	db.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(respondUser, nil)
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
				require.NotEmpty(t, user)
				require.NotEmpty(t, user.ID)
			},
		},
		{
			name: "empty first_name",
			user: models.User{
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Password:    utils.RandomPassword(),
			},
			checker: func(t *testing.T, user models.User, err error) {
				require.Error(t, err)
				require.Empty(t, user)
				require.EqualError(t, err, "first_name: cannot be blank.")
			},
		},
		{
			name: "empty Last_name",
			user: models.User{
				FirstName:   utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Password:    utils.RandomPassword(),
			},
			checker: func(t *testing.T, user models.User, err error) {
				require.Error(t, err)
				require.Empty(t, user)
				require.EqualError(t, err, "last_name: cannot be blank.")
			},
		},
		{
			name: "empty phone number",
			user: models.User{
				FirstName: utils.RandomUserName(),
				LastName:  utils.RandomUserName(),
				Password:  utils.RandomPassword(),
			},
			checker: func(t *testing.T, user models.User, err error) {
				require.Error(t, err)
				require.Empty(t, user)
				require.EqualError(t, err, "phone_number: cannot be blank.")
			},
		},
		{
			name: "empty password",
			user: models.User{
				FirstName:   utils.RandomUserName(),
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
			},
			checker: func(t *testing.T, user models.User, err error) {
				require.Error(t, err)
				require.Empty(t, user)
				require.EqualError(t, err, "cannot be blank")
			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			user, err := appuser.RegisterUser(context.Background(), tc.user)
			tc.checker(t, user, err)
		})
	}
}
