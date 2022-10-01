package persistant

import (
	"context"
	"errors"
	"testing"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
	"github.com/stretchr/testify/require"
)

func TestCreatePymentType(t *testing.T) {

	testdb := Init()
	validPyment := models.PymentType{
		PymentType:         utils.RandomUserName(),
		CreatedByFirstName: utils.RandomUserName(),
		CreatedByLastName:  utils.RandomUserName(),
		Payment:            "1500 ETB",
		NumberOfDays:       30,
	}

	testCase := []struct {
		name   string
		pyment models.PymentType
		check  func(t *testing.T, pyment models.PymentType, err error)
	}{
		{
			name:   "ok",
			pyment: validPyment,
			check: func(t *testing.T, pyment models.PymentType, err error) {
				require.NoError(t, err)
				require.Equal(t, validPyment.CreatedByFirstName, pyment.CreatedByFirstName)
				require.Equal(t, validPyment.CreatedByLastName, validPyment.CreatedByLastName)

				require.Equal(t, validPyment.NumberOfDays, pyment.NumberOfDays)
				require.Equal(t, validPyment.PymentType, pyment.PymentType)

			},
		}, {
			name: "empty creator first name",
			pyment: models.PymentType{
				PymentType:        utils.RandomUserName(),
				CreatedByLastName: utils.RandomUserName(),
				Payment:           "1500 ETB",
				NumberOfDays:      30,
			},
			check: func(t *testing.T, pyment models.PymentType, err error) {
				require.Error(t, err)
				require.EqualError(t, errors.New("created_by_firstname: cannot be blank."), err.Error())
				require.Empty(t, pyment)

			},
		}, {
			name: "empty creator last name",
			pyment: models.PymentType{
				PymentType:         utils.RandomUserName(),
				CreatedByFirstName: utils.RandomUserName(),
				Payment:            "1500 ETB",
				NumberOfDays:       30,
			},
			check: func(t *testing.T, pyment models.PymentType, err error) {
				require.Error(t, err)
				require.EqualError(t, errors.New("created_by_lastname: cannot be blank."), err.Error())
				require.Empty(t, pyment)

			},
		},
		{
			name: "empty pyment",
			pyment: models.PymentType{
				PymentType:         utils.RandomUserName(),
				CreatedByFirstName: utils.RandomUserName(),
				CreatedByLastName:  utils.RandomUserName(),
				NumberOfDays:       30,
			},
			check: func(t *testing.T, pyment models.PymentType, err error) {
				require.Error(t, err)
				require.EqualError(t, errors.New("pyment: cannot be blank."), err.Error())
				require.Empty(t, pyment)

			},
		}, {
			name: "empty numberof days",
			pyment: models.PymentType{
				PymentType:         utils.RandomUserName(),
				CreatedByFirstName: utils.RandomUserName(),
				CreatedByLastName:  utils.RandomUserName(),
				Payment:            "1500 ETB",
			},
			check: func(t *testing.T, pyment models.PymentType, err error) {
				require.Error(t, err)
				require.EqualError(t, errors.New("number_of_days: cannot be blank."), err.Error())
				require.Empty(t, pyment)

			},
		}, {
			name: "negative numberof days",
			pyment: models.PymentType{
				PymentType:         utils.RandomUserName(),
				CreatedByFirstName: utils.RandomUserName(),
				CreatedByLastName:  utils.RandomUserName(),
				NumberOfDays:       -30,
				Payment:            "1500 ETB",
			},
			check: func(t *testing.T, pyment models.PymentType, err error) {
				require.Error(t, err)
				require.EqualError(t, errors.New("number_of_days: must be no less than 1."), err.Error())
				require.Empty(t, pyment)

			},
		}, {
			name: "empty pyment type ",
			pyment: models.PymentType{

				CreatedByFirstName: utils.RandomUserName(),
				CreatedByLastName:  utils.RandomUserName(),
				Payment:            "1500 ETB",
				NumberOfDays:       30,
			},
			check: func(t *testing.T, pyment models.PymentType, err error) {
				require.Error(t, err)
				require.EqualError(t, errors.New("pyment_type: cannot be blank."), err.Error())

				require.Empty(t, pyment)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			pyment, err := testdb.CreatePymentType(context.Background(), tc.pyment)
			tc.check(t, pyment, err)
		})
	}
}

func TestDeletePyment(t *testing.T) {

	testdb := Init()
	validPyment := models.PymentType{
		PymentType:         utils.RandomUserName(),
		CreatedByFirstName: utils.RandomUserName(),
		CreatedByLastName:  utils.RandomUserName(),
		Payment:            "1500 ETB",
		NumberOfDays:       30,
	}
	pyment, _ := testdb.CreatePymentType(context.Background(), validPyment)

	testCase := []struct {
		name   string
		pyment models.PymentType
		check  func(t *testing.T, err error)
	}{
		{
			name:   "ok",
			pyment: models.PymentType{ID: pyment.ID},
			check: func(t *testing.T, err error) {
				require.NoError(t, err)

			},
		}, {
			name:   "empty id",
			pyment: models.PymentType{},
			check: func(t *testing.T, err error) {
				require.Error(t, err)
				require.EqualError(t, errors.New("empy pyment id"), err.Error())

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			err := testdb.DeletePyment(context.Background(), tc.pyment)
			tc.check(t, err)
		})
	}
}

func TestGetAllPymentTypes(t *testing.T) {

	testdb := Init()
	for i := 0; i < 10; i++ {
		validPyment := models.PymentType{
			PymentType:         utils.RandomUserName(),
			CreatedByFirstName: utils.RandomUserName(),
			CreatedByLastName:  utils.RandomUserName(),
			Payment:            "1500 ETB",
			NumberOfDays:       30,
		}
		testdb.CreatePymentType(context.Background(), validPyment)

	}

	testCase := []struct {
		name string

		check func(t *testing.T, pyments []models.PymentType, err error)
	}{
		{
			name: "ok",
			check: func(t *testing.T, pyments []models.PymentType, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(pyments), 10)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			pyments, err := testdb.GetAllPyments(context.Background())
			tc.check(t, pyments, err)
		})
	}
}

func TestGetPymentPyID(t *testing.T) {

	testdb := Init()

	validPyment := models.PymentType{
		PymentType:         utils.RandomUserName(),
		CreatedByFirstName: utils.RandomUserName(),
		CreatedByLastName:  utils.RandomUserName(),
		Payment:            "1500 ETB",
		NumberOfDays:       30,
	}
	py, _ := testdb.CreatePymentType(context.Background(), validPyment)

	testCase := []struct {
		name   string
		pyment models.PymentType
		check  func(t *testing.T, pyments models.PymentType, err error)
	}{
		{
			name:   "ok",
			pyment: py,
			check: func(t *testing.T, pyments models.PymentType, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, pyments)
				require.Equal(t, py.ID, pyments.ID)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			pyments, err := testdb.GetPymentById(context.Background(), tc.pyment)
			tc.check(t, pyments, err)
		})
	}
}
