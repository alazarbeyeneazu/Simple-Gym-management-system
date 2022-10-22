package gymgoers

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
// 	mockdb "github.com/alazarbeyeneazu/Simple-Gym-management-system/mocks/db"
// 	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
// 	"github.com/golang/mock/gomock"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// )

// func TestRegisterGymGoers(t *testing.T) {
// 	validGym_goers := models.Gym_goers{
// 		UserId:               uuid.New(),
// 		CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 		CreatedByFirstName:   utils.RandomUserName(),
// 		CreatedByLastName:    utils.RandomUserName(),
// 		StartDate:            time.Now(),
// 		EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 		PaidBy:               "Bank Transfer",
// 	}
// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().CreateGymGoers(gomock.Any(), gomock.Any()).Return(validGym_goers, nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		geygoer models.Gym_goers
// 		checker func(t *testing.T, pyment models.Gym_goers, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			geygoer: validGym_goers,
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.NoError(t, err)
// 				require.NotEmpty(t, gymgoer)
// 				require.Equal(t, gymgoer.ID, validGym_goers.ID)
// 			},
// 		}, {
// 			name: "empty user id ",
// 			geygoer: models.Gym_goers{
// 				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 				CreatedByFirstName:   utils.RandomUserName(),
// 				CreatedByLastName:    utils.RandomUserName(),
// 				StartDate:            time.Now(),
// 				EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 				PaidBy:               "Bank Transfer",
// 			},
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "user id can not be empty")
// 			},
// 		},
// 		{
// 			name: "empty Created By PhoneNumber",
// 			geygoer: models.Gym_goers{
// 				UserId:             uuid.New(),
// 				CreatedByFirstName: utils.RandomUserName(),
// 				CreatedByLastName:  utils.RandomUserName(),
// 				StartDate:          time.Now(),
// 				EndDate:            time.Now().Add((time.Hour * 24) * 30),
// 				PaidBy:             "Bank Transfer",
// 			},
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "created_by_phonenumber: cannot be blank.")
// 			},
// 		}, {
// 			name: "empty Created By FirstName",
// 			geygoer: models.Gym_goers{
// 				UserId:               uuid.New(),
// 				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 				CreatedByLastName:    utils.RandomUserName(),
// 				StartDate:            time.Now(),
// 				EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 				PaidBy:               "Bank Transfer",
// 			},
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "created_by_firstname: cannot be blank.")
// 			},
// 		}, {
// 			name: "empty Created By LastName",
// 			geygoer: models.Gym_goers{
// 				UserId:               uuid.New(),
// 				CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 				CreatedByFirstName:   utils.RandomUserName(),
// 				StartDate:            time.Now(),
// 				EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 				PaidBy:               "Bank Transfer",
// 			},
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "created_by_lastname: cannot be blank.")
// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			pyment, err := appuser.RegisterGymGoer(context.Background(), tc.geygoer)
// 			tc.checker(t, pyment, err)
// 		})
// 	}
// }

// func TestDeleteGymGoer(t *testing.T) {
// 	validGym_goers := models.Gym_goers{
// 		ID:                   uuid.New(),
// 		UserId:               uuid.New(),
// 		CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 		CreatedByFirstName:   utils.RandomUserName(),
// 		CreatedByLastName:    utils.RandomUserName(),
// 		StartDate:            time.Now(),
// 		EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 		PaidBy:               "Bank Transfer",
// 	}
// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().DeleteGymGoers(gomock.Any(), gomock.Any()).Return(nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		gymgoer models.Gym_goers
// 		checker func(t *testing.T, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			gymgoer: validGym_goers,
// 			checker: func(t *testing.T, err error) {
// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name:    "empty ID",
// 			gymgoer: models.Gym_goers{},
// 			checker: func(t *testing.T, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "gym_goer's id can not be blank")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err := appuser.DeleteGymGoers(context.Background(), tc.gymgoer)
// 			tc.checker(t, err)
// 		})
// 	}
// }

// func TestGetAllGymGoers(t *testing.T) {
// 	var gymgoers []models.Gym_goers
// 	for i := 0; i < 10; i++ {
// 		validGym_goers := models.Gym_goers{
// 			ID:                   uuid.New(),
// 			UserId:               uuid.New(),
// 			CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 			CreatedByFirstName:   utils.RandomUserName(),
// 			CreatedByLastName:    utils.RandomUserName(),
// 			StartDate:            time.Now(),
// 			EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 			PaidBy:               "Bank Transfer",
// 		}
// 		gymgoers = append(gymgoers, validGym_goers)
// 	}

// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetAllGymGoers(gomock.Any()).Return(gymgoers, nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name string

// 		checker func(t *testing.T, gymgoers []models.Gym_goers, err error)
// 	}{
// 		{
// 			name: "ok",

// 			checker: func(t *testing.T, gymgoers []models.Gym_goers, err error) {
// 				require.NoError(t, err)
// 				require.GreaterOrEqual(t, len(gymgoers), 10)

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			gymgoers, err := appuser.GetAllGymGoers(context.Background())
// 			tc.checker(t, gymgoers, err)
// 		})
// 	}
// }

// // Test Get All Pyments
// func TestGetGymgoerById(t *testing.T) {
// 	validGym_goers := models.Gym_goers{
// 		ID:                   uuid.New(),
// 		UserId:               uuid.New(),
// 		CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 		CreatedByFirstName:   utils.RandomUserName(),
// 		CreatedByLastName:    utils.RandomUserName(),
// 		StartDate:            time.Now(),
// 		EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 		PaidBy:               "Bank Transfer",
// 	}

// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetGYmGorsById(gomock.Any(), gomock.Any()).Return(validGym_goers, nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		gymgoer models.Gym_goers
// 		checker func(t *testing.T, gymgoer models.Gym_goers, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			gymgoer: models.Gym_goers{ID: validGym_goers.ID},
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name:    "Empty Id",
// 			gymgoer: models.Gym_goers{},
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "empy gym_goer's id")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			gymgoer, err := appuser.GetGYmGorsById(context.Background(), tc.gymgoer)
// 			tc.checker(t, gymgoer, err)
// 		})
// 	}
// }

// func TestGetGymGoerByUserId(t *testing.T) {
// 	validGym_goers := models.Gym_goers{
// 		ID:                   uuid.New(),
// 		UserId:               uuid.New(),
// 		CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 		CreatedByFirstName:   utils.RandomUserName(),
// 		CreatedByLastName:    utils.RandomUserName(),
// 		StartDate:            time.Now(),
// 		EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 		PaidBy:               "Bank Transfer",
// 	}

// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetGymGoerByUserId(gomock.Any(), gomock.Any()).Return(validGym_goers, nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		gymgoer models.Gym_goers
// 		checker func(t *testing.T, gymgoer models.Gym_goers, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			gymgoer: models.Gym_goers{UserId: validGym_goers.UserId},
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name:    "Empty Id",
// 			gymgoer: models.Gym_goers{},
// 			checker: func(t *testing.T, gymgoer models.Gym_goers, err error) {
// 				require.Error(t, err)
// 				require.Equal(t, err.Error(), "empy user id")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			gymgoer, err := appuser.GetGymGoerByUserId(context.Background(), tc.gymgoer)
// 			tc.checker(t, gymgoer, err)
// 		})
// 	}
// }

// func TestGetGymGoersByCreatorFirstName(t *testing.T) {
// 	var gymgoers []models.Gym_goers
// 	createroFirstName := utils.RandomUserName()
// 	for i := 0; i < 10; i++ {
// 		validGym_goers := models.Gym_goers{
// 			UserId:               uuid.New(),
// 			CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 			CreatedByFirstName:   createroFirstName,
// 			CreatedByLastName:    utils.RandomUserName(),
// 			StartDate:            time.Now(),
// 			EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 			PaidBy:               "Bank Transfer",
// 		}

// 		gymgoers = append(gymgoers, validGym_goers)

// 	}
// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetGymGoerByCreatedByFirstName(gomock.Any(), gomock.Any()).Return(gymgoers, nil)
// 	appuser := InitService(db)

// 	testCase := []struct {
// 		name    string
// 		gymgoer models.User
// 		check   func(t *testing.T, gym_goers []models.Gym_goers, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			gymgoer: models.User{FirstName: createroFirstName},
// 			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
// 				require.NoError(t, err)
// 				require.GreaterOrEqual(t, len(gym_goers), 10)

// 			},
// 		},
// 	}

// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			gym_goers, err := appuser.GetGymGoerByCreatedByFirstName(context.Background(), tc.gymgoer)
// 			tc.check(t, gym_goers, err)
// 		})
// 	}
// }

// func TestGetGymGoersByCreatorLastName(t *testing.T) {
// 	var gymgoers []models.Gym_goers
// 	creatorLastName := utils.RandomUserName()
// 	for i := 0; i < 10; i++ {
// 		validGym_goers := models.Gym_goers{
// 			UserId:               uuid.New(),
// 			CreatedByPhoneNumber: utils.RandomePhoneNumber(),
// 			CreatedByFirstName:   utils.RandomUserName(),
// 			CreatedByLastName:    creatorLastName,
// 			StartDate:            time.Now(),
// 			EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 			PaidBy:               "Bank Transfer",
// 		}

// 		gymgoers = append(gymgoers, validGym_goers)

// 	}
// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetGymGoerByCreatedByLastName(gomock.Any(), gomock.Any()).Return(gymgoers, nil)
// 	appuser := InitService(db)

// 	testCase := []struct {
// 		name    string
// 		gymgoer models.User
// 		check   func(t *testing.T, gym_goers []models.Gym_goers, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			gymgoer: models.User{LastName: creatorLastName},
// 			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
// 				require.NoError(t, err)
// 				require.GreaterOrEqual(t, len(gym_goers), 10)

// 			},
// 		},
// 	}

// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			gym_goers, err := appuser.GetGymGoerByCreatedByLastName(context.Background(), tc.gymgoer)
// 			tc.check(t, gym_goers, err)
// 		})
// 	}
// }

// func TestGetGymgoersByCreatedByPhoneNumber(t *testing.T) {
// 	var gymgoers []models.Gym_goers

// 	creatorPhoneNumber := utils.RandomePhoneNumber()
// 	for i := 0; i < 10; i++ {
// 		validGym_goers := models.Gym_goers{
// 			UserId:               uuid.New(),
// 			CreatedByPhoneNumber: creatorPhoneNumber,
// 			CreatedByFirstName:   utils.RandomUserName(),
// 			CreatedByLastName:    utils.RandomUserName(),
// 			StartDate:            time.Now(),
// 			EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 			PaidBy:               "Bank Transfer",
// 		}

// 		gymgoers = append(gymgoers, validGym_goers)

// 	}

// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetGymGoerByCreatedByPhoneNumber(gomock.Any(), gomock.Any()).Return(gymgoers, nil)
// 	appuser := InitService(db)

// 	testCase := []struct {
// 		name    string
// 		gymgoer models.User
// 		check   func(t *testing.T, gym_goers []models.Gym_goers, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			gymgoer: models.User{PhoneNumber: creatorPhoneNumber},
// 			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
// 				require.NoError(t, err)
// 				require.GreaterOrEqual(t, len(gym_goers), 10)

// 			},
// 		},
// 	}

// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			gym_goers, err := appuser.GetGymGoerByCreatedByPhoneNumber(context.Background(), tc.gymgoer)
// 			tc.check(t, gym_goers, err)
// 		})
// 	}

// }

// func TestGetGymgoersByPaidBy(t *testing.T) {
// 	var gymgoers []models.Gym_goers
// 	paidBy := utils.RandomUserName()
// 	for i := 0; i < 10; i++ {
// 		validGym_goers := models.Gym_goers{
// 			UserId:               uuid.New(),
// 			CreatedByPhoneNumber: utils.RandomUserName(),
// 			CreatedByFirstName:   utils.RandomUserName(),
// 			CreatedByLastName:    utils.RandomUserName(),
// 			StartDate:            time.Now(),
// 			EndDate:              time.Now().Add((time.Hour * 24) * 30),
// 			PaidBy:               paidBy,
// 		}

// 		gymgoers = append(gymgoers, validGym_goers)

// 	}

// 	ctl := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctl)
// 	defer ctl.Finish()
// 	db.EXPECT().GetGymGoerByPaidBy(gomock.Any(), gomock.Any()).Return(gymgoers, nil)
// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		gymgoer models.Gym_goers
// 		check   func(t *testing.T, gym_goers []models.Gym_goers, err error)
// 	}{
// 		{
// 			name:    "ok",
// 			gymgoer: models.Gym_goers{PaidBy: paidBy},
// 			check: func(t *testing.T, gym_goers []models.Gym_goers, err error) {
// 				require.NoError(t, err)
// 				require.GreaterOrEqual(t, len(gym_goers), 10)

// 			},
// 		},
// 	}

// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			gym_goers, err := appuser.GetGymGoerByPaidBy(context.Background(), tc.gymgoer)
// 			tc.check(t, gym_goers, err)
// 		})
// 	}

// }
