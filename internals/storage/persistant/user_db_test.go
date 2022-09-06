package persistant

import (
	"context"
	"testing"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	testdb := Init()
	randomeUser := models.User{
		FirstName:   utils.RandomUserName(),
		LastName:    utils.RandomUserName(),
		PhoneNumber: utils.RandomePhoneNumber(),
		Password:    utils.RandomPassword(),
	}

	testCase := []struct {
		name    string
		user    models.User
		checker func(t *testing.T, user models.User, err error)
	}{
		{
			name: "ok",
			user: randomeUser,
			checker: func(t *testing.T, user models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.FirstName, randomeUser.FirstName)
				require.Equal(t, user.LastName, randomeUser.LastName)
				require.Equal(t, user.PhoneNumber, randomeUser.PhoneNumber)

			},
		},
		{
			name: "empty first name",
			user: models.User{
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Password:    utils.RandomPassword(),
			},
			checker: func(t *testing.T, user models.User, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "first_name: cannot be blank.")
				require.Empty(t, user)

			},
		},
		{
			name: "empty last name",
			user: models.User{
				FirstName:   utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Password:    utils.RandomPassword(),
			},
			checker: func(t *testing.T, user models.User, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "last_name: cannot be blank.")
				require.Empty(t, user)

			},
		}, {
			name: "empty phone number",
			user: models.User{
				FirstName: utils.RandomUserName(),
				LastName:  utils.RandomUserName(),

				Password: utils.RandomPassword(),
			},
			checker: func(t *testing.T, user models.User, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "phone_number: cannot be blank.")
				require.Empty(t, user)

			},
		}, {
			name: "invalide phone number",
			user: models.User{
				FirstName:   utils.RandomUserName(),
				LastName:    utils.RandomUserName(),
				PhoneNumber: "0183",
				Password:    utils.RandomPassword(),
			},
			checker: func(t *testing.T, user models.User, err error) {
				require.Error(t, err)
				require.EqualError(t, err, "phone_number: the length must be exactly 13.")
				require.Empty(t, user)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			user, err := testdb.CreateUser(context.Background(), tc.user)
			tc.checker(t, user, err)
		})
	}

}

func TestGetUsers(t *testing.T) {
	testdb := Init()
	for i := 0; i < 10; i++ {
		randomeUser := models.User{
			FirstName:   utils.RandomUserName(),
			LastName:    utils.RandomUserName(),
			PhoneNumber: utils.RandomePhoneNumber(),
			Password:    utils.RandomPassword(),
		}
		testdb.CreateUser(context.Background(), randomeUser)
	}
	result, err := testdb.GetUsers(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 10)
}

func TestGetUserByFirstName(t *testing.T) {
	testdb := Init()
	firstname := utils.RandomUserName()
	for i := 0; i < 10; i++ {
		randomeUser := models.User{
			FirstName:   firstname,
			LastName:    utils.RandomUserName(),
			PhoneNumber: utils.RandomePhoneNumber(),
			Password:    utils.RandomPassword(),
		}
		testdb.CreateUser(context.Background(), randomeUser)
	}
	testCase := []struct {
		name      string
		firstname string
		checker   func(t *testing.T, users []models.User, err error)
	}{
		{
			name:      "ok",
			firstname: firstname,
			checker: func(t *testing.T, users []models.User, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(users), 10)
			},
		},
		{
			name:      "not found ",
			firstname: utils.RandomUserName(),
			checker: func(t *testing.T, users []models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, len(users), 0)
			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			users, err := testdb.GetUserByFirstName(context.Background(), tc.firstname)
			tc.checker(t, users, err)
		})
	}
}

func TestGetUserByLastName(t *testing.T) {
	testdb := Init()
	lastname := utils.RandomUserName()
	for i := 0; i < 10; i++ {
		randomeUser := models.User{
			LastName:    lastname,
			FirstName:   utils.RandomUserName(),
			PhoneNumber: utils.RandomePhoneNumber(),
			Password:    utils.RandomPassword(),
		}
		testdb.CreateUser(context.Background(), randomeUser)
	}
	testCase := []struct {
		name     string
		lastname string
		checker  func(t *testing.T, users []models.User, err error)
	}{
		{
			name:     "ok",
			lastname: lastname,
			checker: func(t *testing.T, users []models.User, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(users), 10)
			},
		},
		{
			name:     "not found ",
			lastname: utils.RandomUserName(),
			checker: func(t *testing.T, users []models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, len(users), 0)
			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			users, err := testdb.GetUserByLastName(context.Background(), tc.lastname)
			tc.checker(t, users, err)
		})
	}
}
