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
			},
			check: func(t *testing.T, gym_goers models.Gym_goers, err error) {
				require.Error(t, err)
				require.Empty(t, gym_goers)
				require.Equal(t, err.Error(), errors.New("user id can not be empty").Error())

			},
		}, {
			name: "empty Cread by Phone Number",
			gym_goers: models.Gym_goers{
				UserId: uuid.New(),

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

				CreatedByLastName: utils.RandomUserName(),
				StartDate:         time.Now(),
				EndDate:           time.Now().Add((time.Hour * 24) * 30),
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

				CreatedByFirstName: utils.RandomUserName(),
				StartDate:          time.Now(),
				EndDate:            time.Now().Add((time.Hour * 24) * 30),
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

				EndDate: time.Now().Add((time.Hour * 24) * 30),
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

				StartDate: time.Now().Add((time.Hour * 24) * 30),
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
