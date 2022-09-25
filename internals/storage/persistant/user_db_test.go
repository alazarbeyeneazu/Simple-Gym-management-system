package persistant

import (
	"context"
	"errors"
	"testing"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
	"github.com/google/uuid"
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

func TestGetUserByPhoneNumber(t *testing.T) {
	testdb := Init()
	phonenum := utils.RandomePhoneNumber()
	randomeUser := models.User{
		LastName:    utils.RandomUserName(),
		FirstName:   utils.RandomUserName(),
		PhoneNumber: phonenum,
		Password:    utils.RandomPassword(),
	}
	testdb.CreateUser(context.Background(), randomeUser)
	testCase := []struct {
		name        string
		phoneNumber string
		checker     func(t *testing.T, user models.User, err error)
	}{
		{
			name:        "ok",
			phoneNumber: phonenum,
			checker: func(t *testing.T, user models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.PhoneNumber, randomeUser.PhoneNumber)

			},
		}, {
			name:        "not found ",
			phoneNumber: utils.RandomePhoneNumber(),
			checker: func(t *testing.T, user models.User, err error) {
				require.NoError(t, err)
				require.Empty(t, user)

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			user, err := testdb.GetUseByPhoneNumber(context.Background(), tc.phoneNumber)
			tc.checker(t, user, err)
		})
	}
}

func TestGetUserById(t *testing.T) {
	testdb := Init()
	randomeUser := models.User{
		LastName:    utils.RandomUserName(),
		FirstName:   utils.RandomUserName(),
		PhoneNumber: utils.RandomePhoneNumber(),
		Password:    utils.RandomPassword(),
	}
	result, _ := testdb.CreateUser(context.Background(), randomeUser)
	testCase := []struct {
		name    string
		id      uuid.UUID
		checker func(t *testing.T, user models.User, err error)
	}{
		{
			name: "ok",
			id:   result.ID,
			checker: func(t *testing.T, user models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.ID, result.ID)

			},
		}, {
			name: "not found ",
			id:   uuid.New(),
			checker: func(t *testing.T, user models.User, err error) {
				require.NoError(t, err)
				require.Empty(t, user)

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			user, err := testdb.GetUserById(context.Background(), tc.id)
			tc.checker(t, user, err)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testdb := Init()
	phonenum := utils.RandomePhoneNumber()
	randomeUser := models.User{
		LastName:    utils.RandomUserName(),
		FirstName:   utils.RandomUserName(),
		PhoneNumber: phonenum,
		Password:    utils.RandomPassword(),
	}
	testdb.CreateUser(context.Background(), randomeUser)
	testCase := []struct {
		name        string
		phoneNumber string
		checker     func(t *testing.T, err error)
	}{
		{
			name:        "ok",
			phoneNumber: phonenum,
			checker: func(t *testing.T, err error) {
				require.NoError(t, err)

			},
		}, {
			name:        "not found ",
			phoneNumber: utils.RandomePhoneNumber(),
			checker: func(t *testing.T, err error) {
				require.Equal(t, err, errors.New("user not found"))

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			user, _ := testdb.GetUseByPhoneNumber(context.Background(), tc.phoneNumber)
			err := testdb.DeleteUser(context.Background(), user)
			tc.checker(t, err)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	testdb := Init()
	randomUser := models.User{
		FirstName:   utils.RandomUserName(),
		LastName:    utils.RandomUserName(),
		PhoneNumber: utils.RandomePhoneNumber(),
		Password:    utils.RandomPassword(),
	}
	registeredUser, err := testdb.CreateUser(context.Background(), randomUser)
	require.NoError(t, err)

	testCase := []struct {
		name     string
		update   string
		newValue string
		check    func(t *testing.T, newvalue string, user models.User, err error)
	}{
		{
			name:     "update first name",
			update:   "firstName",
			newValue: "helloWorld",
			check: func(t *testing.T, newvalue string, user models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.FirstName, newvalue)

			},
		},
		{
			name:     "update Last name",
			update:   "lastName",
			newValue: "helloWorld",
			check: func(t *testing.T, newvalue string, user models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.LastName, newvalue)

			},
		},
		{
			name:     "update phone number",
			update:   "phoneNumber",
			newValue: "+251975146165",
			check: func(t *testing.T, newvalue string, user models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.PhoneNumber, newvalue)

			},
		},
		{
			name:     "update first name",
			update:   "password",
			newValue: "hello password",
			check: func(t *testing.T, newvalue string, user models.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.Password, newvalue)

			},
		},
	}

	for _, tc := range testCase {
		switch tc.update {
		case "firstName":
			user, err := testdb.UpdateUser(context.Background(), models.User{FirstName: tc.newValue}, registeredUser)
			tc.check(t, tc.newValue, user, err)
		case "lastName":
			user, err := testdb.UpdateUser(context.Background(), models.User{LastName: tc.newValue}, registeredUser)
			tc.check(t, tc.newValue, user, err)
		case "phoneNumber":
			user, err := testdb.UpdateUser(context.Background(), models.User{PhoneNumber: tc.newValue}, registeredUser)
			tc.check(t, tc.newValue, user, err)
		case "password":
			user, err := testdb.UpdateUser(context.Background(), models.User{Password: tc.newValue}, registeredUser)
			tc.check(t, tc.newValue, user, err)
		}
	}
}
