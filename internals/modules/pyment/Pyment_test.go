package pyment

// import (
// 	"context"
// 	"testing"

// 	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
// 	mockdb "github.com/alazarbeyeneazu/Simple-Gym-management-system/mocks/db"
// 	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
// 	"github.com/golang/mock/gomock"
// 	"github.com/google/uuid"

// 	"github.com/stretchr/testify/require"
// )

// func TestCreatePyment(t *testing.T) {
// 	validPyment := models.PymentType{
// 		ID:                 uuid.New(),
// 		PymentType:         utils.RandomUserName(),
// 		CreatedByFirstName: utils.RandomUserName(),
// 		CreatedByLastName:  utils.RandomUserName(),
// 		Payment:            "1500 ETB",
// 		NumberOfDays:       30,
// 	}
// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().CreatePymentType(gomock.Any(), gomock.Any()).Return(validPyment, nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		pyment  models.PymentType
// 		checker func(t *testing.T, pyment models.PymentType, err error)
// 	}{
// 		{
// 			name:   "ok",
// 			pyment: validPyment,
// 			checker: func(t *testing.T, pyment models.PymentType, err error) {
// 				require.NoError(t, err)
// 				require.NotEmpty(t, pyment)
// 				require.NotEmpty(t, pyment.ID)
// 			},
// 		}, {
// 			name: "empty PymentType",
// 			pyment: models.PymentType{

// 				CreatedByFirstName: utils.RandomUserName(),
// 				CreatedByLastName:  utils.RandomUserName(),
// 				Payment:            "1500 ETB",
// 				NumberOfDays:       30,
// 			},
// 			checker: func(t *testing.T, pyment models.PymentType, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "pyment_type: cannot be blank.")

// 			},
// 		}, {
// 			name: "empty CreatedByFirstName",
// 			pyment: models.PymentType{
// 				PymentType:        utils.RandomUserName(),
// 				CreatedByLastName: utils.RandomUserName(),
// 				Payment:           "1500 ETB",
// 				NumberOfDays:      30,
// 			},
// 			checker: func(t *testing.T, pyment models.PymentType, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "created_by_firstname: cannot be blank.")

// 			},
// 		},
// 		{
// 			name: "empty CreatedByLastName",
// 			pyment: models.PymentType{
// 				PymentType:         utils.RandomUserName(),
// 				CreatedByFirstName: utils.RandomUserName(),
// 				Payment:            "1500 ETB",
// 				NumberOfDays:       30,
// 			},
// 			checker: func(t *testing.T, pyment models.PymentType, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "created_by_lastname: cannot be blank.")

// 			},
// 		}, {
// 			name: "empty Pyment",
// 			pyment: models.PymentType{
// 				PymentType: utils.RandomUserName(),

// 				CreatedByFirstName: utils.RandomUserName(),
// 				CreatedByLastName:  utils.RandomUserName(),

// 				NumberOfDays: 30,
// 			},
// 			checker: func(t *testing.T, pyment models.PymentType, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "pyment: cannot be blank.")

// 			},
// 		}, {
// 			name: "empty Number Of Days",
// 			pyment: models.PymentType{
// 				PymentType: utils.RandomUserName(),

// 				CreatedByFirstName: utils.RandomUserName(),
// 				CreatedByLastName:  utils.RandomUserName(),
// 				Payment:            "1500 ETB",
// 			},
// 			checker: func(t *testing.T, pyment models.PymentType, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "number_of_days: cannot be blank.")

// 			},
// 		}, {
// 			name: "Invalid number of days",
// 			pyment: models.PymentType{
// 				PymentType: utils.RandomUserName(),

// 				CreatedByFirstName: utils.RandomUserName(),
// 				CreatedByLastName:  utils.RandomUserName(),
// 				Payment:            "1500 ETB",
// 				NumberOfDays:       -30,
// 			},
// 			checker: func(t *testing.T, pyment models.PymentType, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "number_of_days: must be no less than 1.")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			pyment, err := appuser.CreatePyment(context.Background(), tc.pyment)
// 			tc.checker(t, pyment, err)
// 		})
// 	}
// }

// // Test Delete Pyments
// func TestDeletePyement(t *testing.T) {
// 	validPyment := models.PymentType{
// 		ID:                 uuid.New(),
// 		PymentType:         utils.RandomUserName(),
// 		CreatedByFirstName: utils.RandomUserName(),
// 		CreatedByLastName:  utils.RandomUserName(),
// 		Payment:            "1500 ETB",
// 		NumberOfDays:       30,
// 	}
// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().DeletePyment(gomock.Any(), gomock.Any()).Return(nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		pyment  models.PymentType
// 		checker func(t *testing.T, err error)
// 	}{
// 		{
// 			name:   "ok",
// 			pyment: validPyment,
// 			checker: func(t *testing.T, err error) {
// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name:   "empty ID",
// 			pyment: models.PymentType{},
// 			checker: func(t *testing.T, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "pyment id can not be blank")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err := appuser.DeletePyment(context.Background(), tc.pyment)
// 			tc.checker(t, err)
// 		})
// 	}
// }

// // Test Get All Pyments
// func TestGetAllPyments(t *testing.T) {
// 	var pyments []models.PymentType
// 	for i := 0; i < 10; i++ {
// 		validPyment := models.PymentType{
// 			ID:                 uuid.New(),
// 			PymentType:         utils.RandomUserName(),
// 			CreatedByFirstName: utils.RandomUserName(),
// 			CreatedByLastName:  utils.RandomUserName(),
// 			Payment:            "1500 ETB",
// 			NumberOfDays:       30,
// 		}
// 		pyments = append(pyments, validPyment)
// 	}
// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetAllPyments(gomock.Any()).Return(pyments, nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		checker func(t *testing.T, pyments []models.PymentType, err error)
// 	}{
// 		{
// 			name: "ok",

// 			checker: func(t *testing.T, pyments []models.PymentType, err error) {
// 				require.NoError(t, err)
// 				require.GreaterOrEqual(t, len(pyments), 10)
// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			pyemnts, err := appuser.GetAllPyments(context.Background())
// 			tc.checker(t, pyemnts, err)
// 		})
// 	}
// }

// func TestGetPymentById(t *testing.T) {

// 	validPyment := models.PymentType{
// 		ID:                 uuid.New(),
// 		PymentType:         utils.RandomUserName(),
// 		CreatedByFirstName: utils.RandomUserName(),
// 		CreatedByLastName:  utils.RandomUserName(),
// 		Payment:            "1500 ETB",
// 		NumberOfDays:       30,
// 	}
// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetPymentById(gomock.Any(), gomock.Any()).Return(validPyment, nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		payment models.PymentType
// 		checker func(t *testing.T, pyments models.PymentType, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			payment: models.PymentType{ID: validPyment.ID},
// 			checker: func(t *testing.T, pyments models.PymentType, err error) {
// 				require.NoError(t, err)
// 				require.Equal(t, pyments.ID, validPyment.ID)

// 			},
// 		}, {
// 			name:    "empty id ",
// 			payment: models.PymentType{},
// 			checker: func(t *testing.T, pyments models.PymentType, err error) {
// 				require.Error(t, err)
// 				require.Empty(t, pyments.ID, validPyment.ID)

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			pyemnts, err := appuser.GetPymentById(context.Background(), tc.payment)
// 			tc.checker(t, pyemnts, err)
// 		})
// 	}
// }
