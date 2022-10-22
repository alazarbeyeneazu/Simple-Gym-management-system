package user

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
// 	mockdb "github.com/alazarbeyeneazu/Simple-Gym-management-system/mocks/db"
// 	"github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/utils"
// 	"github.com/golang/mock/gomock"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// )

// func TestGetUserById(t *testing.T) {
// 	respondUser := models.User{
// 		ID:          uuid.New(),
// 		FirstName:   utils.RandomUserName(),
// 		LastName:    utils.RandomUserName(),
// 		PhoneNumber: utils.RandomePhoneNumber(),
// 		Password:    utils.RandomPassword(),
// 	}
// 	ctr := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctr)
// 	defer ctr.Finish()

// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		user    models.User
// 		checker func(t *testing.T, user models.User, err error)
// 	}{
// 		{
// 			name: "ok",
// 			user: respondUser,
// 			checker: func(t *testing.T, user models.User, err error) {
// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name: "not found",
// 			user: models.User{ID: uuid.New()},
// 			checker: func(t *testing.T, user models.User, err error) {
// 				require.EqualError(t, err, "user not found")

// 			},
// 		}, {
// 			name: "empty id",
// 			user: models.User{},
// 			checker: func(t *testing.T, user models.User, err error) {
// 				require.EqualError(t, err, "id can not be blank")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			switch tc.name {
// 			case "ok":
// 				db.EXPECT().GetUserById(gomock.Any(), respondUser.ID).Return(respondUser, nil)
// 				user, err := appuser.GetUserById(context.Background(), tc.user)
// 				tc.checker(t, user, err)
// 			case "not found":
// 				db.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(models.User{}, errors.New("user not found"))
// 				user, err := appuser.GetUserById(context.Background(), tc.user)
// 				tc.checker(t, user, err)
// 			case "empty id":
// 				user, err := appuser.GetUserById(context.Background(), tc.user)
// 				tc.checker(t, user, err)

// 			}

// 		})
// 	}

// }

// //get users by first name test
// func TestGetUsersByFirstName(t *testing.T) {
// 	firsname := utils.RandomUserName()
// 	var respondUser []models.User
// 	for i := 0; i < 10; i++ {
// 		resp := models.User{
// 			ID:          uuid.New(),
// 			FirstName:   firsname,
// 			LastName:    utils.RandomUserName(),
// 			PhoneNumber: utils.RandomePhoneNumber(),
// 			Password:    utils.RandomPassword(),
// 		}
// 		respondUser = append(respondUser, resp)
// 	}

// 	ctr := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctr)
// 	defer ctr.Finish()

// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		user    models.User
// 		checker func(t *testing.T, user []models.User, err error)
// 	}{
// 		{
// 			name: "ok",
// 			user: models.User{FirstName: firsname},
// 			checker: func(t *testing.T, user []models.User, err error) {
// 				require.Equal(t, len(user), 10)
// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name: "not found",
// 			user: models.User{FirstName: utils.RandomUserName()},
// 			checker: func(t *testing.T, user []models.User, err error) {
// 				require.EqualError(t, err, "user not found")

// 			},
// 		}, {
// 			name: "empty firstname",
// 			user: models.User{},
// 			checker: func(t *testing.T, user []models.User, err error) {
// 				require.EqualError(t, err, "first name cannot be blank")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			switch tc.name {
// 			case "ok":
// 				db.EXPECT().GetUserByFirstName(gomock.Any(), gomock.Any()).Return(respondUser, nil)
// 				user, err := appuser.GetUsersByFirstName(context.Background(), tc.user)
// 				tc.checker(t, user, err)
// 			case "not found":
// 				db.EXPECT().GetUserByFirstName(gomock.Any(), gomock.Any()).Return([]models.User{}, errors.New("user not found"))
// 				user, err := appuser.GetUsersByFirstName(context.Background(), tc.user)
// 				tc.checker(t, user, err)
// 			case "empty firstname":
// 				user, err := appuser.GetUsersByFirstName(context.Background(), tc.user)
// 				tc.checker(t, user, err)

// 			}

// 		})
// 	}

// }

// //get users by last name test
// func TestGetUsersByLastName(t *testing.T) {
// 	lastname := utils.RandomUserName()
// 	var respondUser []models.User
// 	for i := 0; i < 10; i++ {
// 		resp := models.User{
// 			ID:          uuid.New(),
// 			FirstName:   utils.RandomUserName(),
// 			LastName:    lastname,
// 			PhoneNumber: utils.RandomePhoneNumber(),
// 			Password:    utils.RandomPassword(),
// 		}
// 		respondUser = append(respondUser, resp)
// 	}

// 	ctr := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctr)
// 	defer ctr.Finish()

// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		user    models.User
// 		checker func(t *testing.T, user []models.User, err error)
// 	}{
// 		{
// 			name: "ok",
// 			user: models.User{LastName: lastname},
// 			checker: func(t *testing.T, user []models.User, err error) {
// 				require.Equal(t, len(user), 10)
// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name: "not found",
// 			user: models.User{LastName: utils.RandomUserName()},
// 			checker: func(t *testing.T, user []models.User, err error) {
// 				require.EqualError(t, err, "user not found")

// 			},
// 		}, {
// 			name: "empty lastname",
// 			user: models.User{},
// 			checker: func(t *testing.T, user []models.User, err error) {
// 				require.EqualError(t, err, "last name cannot be blank")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			switch tc.name {
// 			case "ok":
// 				db.EXPECT().GetUserByLastName(gomock.Any(), gomock.Any()).Return(respondUser, nil)
// 				user, err := appuser.GetUserByLastName(context.Background(), tc.user)
// 				tc.checker(t, user, err)
// 			case "not found":
// 				db.EXPECT().GetUserByLastName(gomock.Any(), gomock.Any()).Return([]models.User{}, errors.New("user not found"))
// 				user, err := appuser.GetUserByLastName(context.Background(), tc.user)
// 				tc.checker(t, user, err)
// 			case "empty lastname":
// 				user, err := appuser.GetUserByLastName(context.Background(), tc.user)
// 				tc.checker(t, user, err)

// 			}

// 		})
// 	}

// }

// //get users by last name test
// func TestGetUserByPhoneNumber(t *testing.T) {

// 	respondUser := models.User{
// 		ID:          uuid.New(),
// 		FirstName:   utils.RandomUserName(),
// 		LastName:    utils.RandomUserName(),
// 		PhoneNumber: utils.RandomePhoneNumber(),
// 		Password:    utils.RandomPassword(),
// 	}

// 	ctr := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctr)
// 	defer ctr.Finish()

// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		user    models.User
// 		checker func(t *testing.T, user models.User, err error)
// 	}{
// 		{
// 			name: "ok",
// 			user: models.User{PhoneNumber: respondUser.PhoneNumber},
// 			checker: func(t *testing.T, user models.User, err error) {

// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name: "not found",
// 			user: models.User{PhoneNumber: utils.RandomePhoneNumber()},
// 			checker: func(t *testing.T, user models.User, err error) {
// 				require.EqualError(t, err, "user not found")

// 			},
// 		}, {
// 			name: "empty lastname",
// 			user: models.User{},
// 			checker: func(t *testing.T, user models.User, err error) {
// 				require.EqualError(t, err, "phone number cannot be blank")

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			switch tc.name {
// 			case "ok":
// 				db.EXPECT().GetUseByPhoneNumber(gomock.Any(), gomock.Any()).Return(respondUser, nil)
// 				user, err := appuser.GetUserByPhoneNumber(context.Background(), tc.user)
// 				tc.checker(t, user, err)
// 			case "not found":
// 				db.EXPECT().GetUseByPhoneNumber(gomock.Any(), gomock.Any()).Return(models.User{}, errors.New("user not found"))
// 				user, err := appuser.GetUserByPhoneNumber(context.Background(), tc.user)
// 				tc.checker(t, user, err)
// 			case "empty lastname":
// 				user, err := appuser.GetUserByPhoneNumber(context.Background(), tc.user)
// 				tc.checker(t, user, err)

// 			}

// 		})
// 	}

// }

// //get all users
// func TestGetAllUsers(t *testing.T) {

// 	var respondUser []models.User
// 	for i := 0; i < 10; i++ {
// 		resp := models.User{
// 			ID:          uuid.New(),
// 			FirstName:   utils.RandomUserName(),
// 			LastName:    utils.RandomUserName(),
// 			PhoneNumber: utils.RandomePhoneNumber(),
// 			Password:    utils.RandomPassword(),
// 		}
// 		respondUser = append(respondUser, resp)
// 	}

// 	ctr := gomock.NewController(t)
// 	db := mockdb.NewMockDBPort(ctr)
// 	defer ctr.Finish()

// 	appuser := InitService(db)
// 	testCase := []struct {
// 		name    string
// 		checker func(t *testing.T, user []models.User, err error)
// 	}{
// 		{
// 			name: "ok",

// 			checker: func(t *testing.T, user []models.User, err error) {
// 				require.Equal(t, len(user), 10)
// 				require.NoError(t, err)

// 			},
// 		}, {
// 			name: "empty users",
// 			checker: func(t *testing.T, user []models.User, err error) {
// 				require.NoError(t, err)
// 				require.Empty(t, user)

// 			},
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			switch tc.name {
// 			case "ok":
// 				db.EXPECT().GetUsers(gomock.Any()).Return(respondUser, nil)
// 				user, err := appuser.GetAllUsers(context.Background())
// 				tc.checker(t, user, err)
// 			case "empty users":
// 				db.EXPECT().GetUsers(gomock.Any()).Return([]models.User{}, nil)
// 				user, err := appuser.GetAllUsers(context.Background())
// 				tc.checker(t, user, err)

// 			}

// 		})
// 	}
// }
