package persistant

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateGym_goers(t *testing.T) {

	testdb := Init()
	validGym_goers := models.Gym_goers{
		UserId:               uuid.New(),
		CreatedByPhoneNumber: utils.RandomePhoneNumber(),
		CreatedByFirstName:   utils.RandomUserName(),
		CreatedByLastName:    utils.RandomUserName(),
		StartDate:            time.Now(),
		EndDate:              time.Now().Add((time.Hour * 24) * 30),
		PaidBy:               "Bank Transfer",
	}

	testCase := []struct {
		name      string
		gym_goers models.Gym_goers
		check     func(t *testing.T, gym_goers models.Gym_goers, err error)
	}{
		{
			name:      "ok",
			gym_goers: validGym_goers,
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.NoError(t, err)
				require.Equal(t, gym_goers.CreatedByFirstName, gym_goers.CreatedByFirstName)
				require.Equal(t, gym_goers.CreatedByLastName, gym_goers.CreatedByLastName)

				require.Equal(t, gym_goers.CreatedByPhoneNumber, gym_goers.CreatedByPhoneNumber)

			},
		}, {
			name: "empty UserID",
			gym_goers: models.Gym_goers{

				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
				CreatedByFirstName:   utils.RandomUserName(),
				CreatedByLastName:    utils.RandomUserName(),
				StartDate:            time.Now(),
				EndDate:              time.Now().Add((time.Hour * 24) * 30),
				PaidBy:               "Bank Transfer",
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("user id can not be empty").Error())

			},
		}, {
			name: "empty Cread by Phone Number",
			gym_goers: models.Gym_goers{
				UserId:             uuid.New(),
				PaidBy:             "Bank Transfer",
				CreatedByFirstName: utils.RandomUserName(),
				CreatedByLastName:  utils.RandomUserName(),
				StartDate:          time.Now(),
				EndDate:            time.Now().Add((time.Hour * 24) * 30),
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("created_by_phonenumber: cannot be blank.").Error())

			},
		}, {
			name: "empty Cread by first name",
			gym_goers: models.Gym_goers{
				UserId:               uuid.New(),
				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
				PaidBy:               "Bank Transfer",
				CreatedByLastName:    utils.RandomUserName(),
				StartDate:            time.Now(),
				EndDate:              time.Now().Add((time.Hour * 24) * 30),
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("created_by_firstname: cannot be blank.").Error())

			},
		}, {
			name: "empty Cread by last name",
			gym_goers: models.Gym_goers{
				UserId:               uuid.New(),
				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
				PaidBy:               "Bank Transfer",
				CreatedByFirstName:   utils.RandomUserName(),
				StartDate:            time.Now(),
				EndDate:              time.Now().Add((time.Hour * 24) * 30),
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("created_by_lastname: cannot be blank.").Error())

			},
		}, {
			name: "empty Cread by start date",
			gym_goers: models.Gym_goers{
				UserId:               uuid.New(),
				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
				CreatedByFirstName:   utils.RandomUserName(),
				CreatedByLastName:    utils.RandomUserName(),
				PaidBy:               "Bank Transfer",
				EndDate:              time.Now().Add((time.Hour * 24) * 30),
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("start_date: cannot be blank.").Error())

			},
		}, {
			name: "empty Cread by end date",
			gym_goers: models.Gym_goers{
				UserId:               uuid.New(),
				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
				CreatedByFirstName:   utils.RandomUserName(),
				CreatedByLastName:    utils.RandomUserName(),
				PaidBy:               "Bank Transfer",
				StartDate:            time.Now().Add((time.Hour * 24) * 30),
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("end_date: cannot be blank.").Error())

			},
		}, {
			name: "invalid  end date",
			gym_goers: models.Gym_goers{
				UserId:               uuid.New(),
				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
				CreatedByFirstName:   utils.RandomUserName(),
				CreatedByLastName:    utils.RandomUserName(),
				StartDate:            time.Now().Add((time.Hour * 24) * 30),
				EndDate:              time.Now(),
				PaidBy:               "Bank Transfer",
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("end date can not be less than start date").Error())

			},
		}, {
			name: "invalid  start date",
			gym_goers: models.Gym_goers{
				UserId:               uuid.New(),
				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
				CreatedByFirstName:   utils.RandomUserName(),
				CreatedByLastName:    utils.RandomUserName(),
				EndDate:              time.Now().Add((time.Hour * 24) * 30),
				StartDate:            time.Now().Add(time.Hour * -2),
				PaidBy:               "Bank Transfer",
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("start date should be equal or after today").Error())

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gym_goer, err := testdb.CreateGymGoers(context.Background(), tc.gym_goers)
			tc.check(t, gym_goer, err)
		})
	}
}
func TestDeleteGymGoers(t *testing.T) {
	testdb := Init()
	validGym_goers := models.Gym_goers{
		UserId:               uuid.New(),
		CreatedByPhoneNumber: utils.RandomePhoneNumber(),
		CreatedByFirstName:   utils.RandomUserName(),
		CreatedByLastName:    utils.RandomUserName(),
		StartDate:            time.Now(),
		EndDate:              time.Now().Add((time.Hour * 24) * 30),
		PaidBy:               "Bank Transfer",
	}
	gym_goer, _ := testdb.CreateGymGoers(context.Background(), validGym_goers)
	testCase := []struct {
		name      string
		gym_goers models.Gym_goers
		check     func(t *testing.T, err error)
	}{
		{
			name:      "ok",
			gym_goers: models.Gym_goers{ID: gym_goer.ID},
			check: func(t *testing.T, err error) {
				require.NoError(t, err)

			},
		},
		{
			name:      "empty id",
			gym_goers: models.Gym_goers{},
			check: func(t *testing.T, err error) {
				require.Equal(t, err, errors.New("gym_goer's id can not be blank"))
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			err := testdb.DeleteGymGoers(context.Background(), tc.gym_goers)
			tc.check(t, err)
		})
	}

}

func TestGetAllGymGoers(t *testing.T) {
	testdb := Init()
	for i := 0; i < 10; i++ {
		validGym_goers := models.Gym_goers{
			UserId:               uuid.New(),
			CreatedByPhoneNumber: utils.RandomePhoneNumber(),
			CreatedByFirstName:   utils.RandomUserName(),
			CreatedByLastName:    utils.RandomUserName(),
			StartDate:            time.Now(),
			EndDate:              time.Now().Add((time.Hour * 24) * 30),
			PaidBy:               "Bank Transfer",
		}

		testdb.CreateGymGoers(context.Background(), validGym_goers)

	}
	testCase := []struct {
		name  string
		check func(t *testing.T, gym_goers []models.Gym_goers, err error)
	}{
		{
			name: "ok",

			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(gym_goers), 10)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gym_goers, err := testdb.GetAllGymGoers(context.Background())
			tc.check(t, gym_goers, err)
		})
	}

}

func TestGetGymGoersById(t *testing.T) {
	testdb := Init()
	validGym_goers := models.Gym_goers{
		UserId:               uuid.New(),
		CreatedByPhoneNumber: utils.RandomePhoneNumber(),
		CreatedByFirstName:   utils.RandomUserName(),
		CreatedByLastName:    utils.RandomUserName(),
		StartDate:            time.Now(),
		EndDate:              time.Now().Add((time.Hour * 24) * 30),
		PaidBy:               "Bank Transfer",
	}

	gymgoers, _ := testdb.CreateGymGoers(context.Background(), validGym_goers)
	testCase := []struct {
		name    string
		gymgoer models.Gym_goers
		check   func(t *testing.T, gym_goers models.Gym_goers, err error)
	}{
		{
			name:    "ok",
			gymgoer: models.Gym_goers{ID: gymgoers.ID},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.NoError(t, err)
				require.Equal(t, gym_goers.ID, gymgoers.ID)
			},
		}, {
			name:    "empty Id",
			gymgoer: models.Gym_goers{},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Equal(t, err, errors.New("empy gym_goer's id"))
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gym_goers, err := testdb.GetGYmGorsById(context.Background(), tc.gymgoer)
			tc.check(t, gym_goers, err)
		})
	}

}

func TestGetGymGoerByUserId(t *testing.T) {
	testdb := Init()
	validGym_goers := models.Gym_goers{
		UserId:               uuid.New(),
		CreatedByPhoneNumber: utils.RandomePhoneNumber(),
		CreatedByFirstName:   utils.RandomUserName(),
		CreatedByLastName:    utils.RandomUserName(),
		StartDate:            time.Now(),
		EndDate:              time.Now().Add((time.Hour * 24) * 30),
		PaidBy:               "Bank Transfer",
	}

	gymgoers, _ := testdb.CreateGymGoers(context.Background(), validGym_goers)
	testCase := []struct {
		name    string
		gymgoer models.Gym_goers
		check   func(t *testing.T, gym_goers models.Gym_goers, err error)
	}{
		{
			name:    "ok",
			gymgoer: models.Gym_goers{UserId: gymgoers.UserId},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.NoError(t, err)
				require.Equal(t, gym_goers.UserId, gymgoers.UserId)
				require.Equal(t, gym_goers.ID, gymgoers.ID)
			},
		}, {
			name:    "empty Id",
			gymgoer: models.Gym_goers{},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Equal(t, err, errors.New("empy user id"))
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gym_goers, err := testdb.GetGymGoerByUserId(context.Background(), tc.gymgoer)
			tc.check(t, gym_goers, err)
		})
	}

}

func TestGetGymgoersByCreatedByFirstName(t *testing.T) {
	testdb := Init()
	createroFirstName := utils.RandomUserName()
	for i := 0; i < 10; i++ {
		validGym_goers := models.Gym_goers{
			UserId:               uuid.New(),
			CreatedByPhoneNumber: utils.RandomePhoneNumber(),
			CreatedByFirstName:   createroFirstName,
			CreatedByLastName:    utils.RandomUserName(),
			StartDate:            time.Now(),
			EndDate:              time.Now().Add((time.Hour * 24) * 30),
			PaidBy:               "Bank Transfer",
		}

		testdb.CreateGymGoers(context.Background(), validGym_goers)

	}
	testCase := []struct {
		name    string
		gymgoer models.User
		check   func(t *testing.T, gym_goers []models.Gym_goers, err error)
	}{
		{
			name:    "ok",
			gymgoer: models.User{FirstName: createroFirstName},
			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(gym_goers), 10)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gym_goers, err := testdb.GetGymGoerByCreatedByFirstName(context.Background(), tc.gymgoer)
			tc.check(t, gym_goers, err)
		})
	}

}

func TestGetGymgoersByCreatedByLastName(t *testing.T) {
	testdb := Init()
	creatorLastName := utils.RandomUserName()
	for i := 0; i < 10; i++ {
		validGym_goers := models.Gym_goers{
			UserId:               uuid.New(),
			CreatedByPhoneNumber: utils.RandomePhoneNumber(),
			CreatedByFirstName:   utils.RandomUserName(),
			CreatedByLastName:    creatorLastName,
			StartDate:            time.Now(),
			EndDate:              time.Now().Add((time.Hour * 24) * 30),
			PaidBy:               "Bank Transfer",
		}

		testdb.CreateGymGoers(context.Background(), validGym_goers)

	}
	testCase := []struct {
		name    string
		gymgoer models.User
		check   func(t *testing.T, gym_goers []models.Gym_goers, err error)
	}{
		{
			name:    "ok",
			gymgoer: models.User{LastName: creatorLastName},
			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(gym_goers), 10)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gym_goers, err := testdb.GetGymGoerByCreatedByLastName(context.Background(), tc.gymgoer)
			tc.check(t, gym_goers, err)
		})
	}

}

func TestGetGymgoersByCreatedByPhoneNumber(t *testing.T) {
	testdb := Init()
	creatorPhoneNumber := utils.RandomePhoneNumber()
	for i := 0; i < 10; i++ {
		validGym_goers := models.Gym_goers{
			UserId:               uuid.New(),
			CreatedByPhoneNumber: creatorPhoneNumber,
			CreatedByFirstName:   utils.RandomUserName(),
			CreatedByLastName:    utils.RandomUserName(),
			StartDate:            time.Now(),
			EndDate:              time.Now().Add((time.Hour * 24) * 30),
			PaidBy:               "Bank Transfer",
		}

		testdb.CreateGymGoers(context.Background(), validGym_goers)

	}
	testCase := []struct {
		name    string
		gymgoer models.User
		check   func(t *testing.T, gym_goers []models.Gym_goers, err error)
	}{
		{
			name:    "ok",
			gymgoer: models.User{PhoneNumber: creatorPhoneNumber},
			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(gym_goers), 10)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gym_goers, err := testdb.GetGymGoerByCreatedByPhoneNumber(context.Background(), tc.gymgoer)
			tc.check(t, gym_goers, err)
		})
	}

}

func TestGetGymgoersByPaidBy(t *testing.T) {
	testdb := Init()
	paidBy := utils.RandomUserName()
	for i := 0; i < 10; i++ {
		validGym_goers := models.Gym_goers{
			UserId:               uuid.New(),
			CreatedByPhoneNumber: utils.RandomUserName(),
			CreatedByFirstName:   utils.RandomUserName(),
			CreatedByLastName:    utils.RandomUserName(),
			StartDate:            time.Now(),
			EndDate:              time.Now().Add((time.Hour * 24) * 30),
			PaidBy:               paidBy,
		}

		testdb.CreateGymGoers(context.Background(), validGym_goers)

	}
	testCase := []struct {
		name    string
		gymgoer models.Gym_goers
		check   func(t *testing.T, gym_goers []models.Gym_goers, err error)
	}{
		{
			name:    "ok",
			gymgoer: models.Gym_goers{PaidBy: paidBy},
			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(gym_goers), 10)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gym_goers, err := testdb.GetGymGoerByPaidBy(context.Background(), tc.gymgoer)
			tc.check(t, gym_goers, err)
		})
	}

}
