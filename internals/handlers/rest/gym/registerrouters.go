package rest

import (
	"net/http"

	routers "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/routes"

	authn "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/middlewares/Authn"
	"github.com/gin-gonic/gin"
)

func (uh *restHandler) StartRoutes() []routers.Router {
	registerUser := []routers.Router{
		{
			Method:      http.MethodGet,
			Path:        "/view/login",
			Handler:     uh.GetLoginPage,
			MiddleWares: []gin.HandlerFunc{},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/users",
			Handler:     uh.GetRegistrationPage,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/users/:adminId",
			Handler:     uh.EditAdmin,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/dashboard",
			Handler:     uh.GetDashBoard,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/roles",
			Handler:     uh.GetRoles,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/gym-goers",
			Handler:     uh.GetGym_goers,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/gym-goers/:id",
			Handler:     uh.GymGoersSimpleDetail,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/v1/api/gymgoers/delete/:id",
			Handler:     uh.DeleteGymGoer,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/payment",
			Handler:     uh.GetPayment,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/payment/:id",
			Handler:     uh.EditPayment,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/setting",
			Handler:     uh.GetSetting,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/gym-goers-detail/:id",
			Handler:     uh.GetGym_goers_detail,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/logout",
			Handler:     uh.LogOut,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodPut,
			Path:        "/v1/api/users",
			Handler:     uh.RegisterUser,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodPut,
			Path:        "/v1/api/admins/:adminId",
			Handler:     uh.UpdateAdmin,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/v1/api/roles",
			Handler:     uh.CreateRole,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/v1/api/roles/:roleId",
			Handler:     uh.DeleteRole,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodPost,
			Path:        "/v1/api/settings",
			Handler:     uh.UpdateSetting,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		},
		{
			Method:      http.MethodGet,
			Path:        "/v1/api/updaterole",
			Handler:     uh.UpdateRole,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		},
		{
			Method:      http.MethodGet,
			Path:        "/view/roles/:role",
			Handler:     uh.EditRole,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		},
		{
			Method:  http.MethodPost,
			Path:    "/v1/api/login",
			Handler: uh.LoginUser,
		},
		{
			Method:      http.MethodGet,
			Path:        "/v1/api/users",
			Handler:     uh.GetAllUsers,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:  http.MethodPost,
			Path:    "/v1/api/admins",
			Handler: uh.RegisterAdmin,
			// MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		},
		{
			Method:      http.MethodGet,
			Path:        "/v1/api/admins/:user_id",
			Handler:     uh.DeleteAdmin,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		},
		{
			Method:      http.MethodGet,
			Path:        "/v1/api/pyments",
			Handler:     uh.GetAllPyments,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/v1/api/pyments/:id",
			Handler:     uh.GetPymentById,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodPost,
			Path:        "/v1/api/pyments",
			Handler:     uh.CreatePyment,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodPut,
			Path:        "/v1/api/pyments/:id",
			Handler:     uh.UpdatePyment,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/v1/api/pyments/delete/:id",
			Handler:     uh.DeletePyment,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodPost,
			Path:        "/v1/api/gymgoers",
			Handler:     uh.RegisterGymGoer,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/v1/api/gymgoers",
			Handler:     uh.GetAllGymGoers,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodPut,
			Path:        "/v1/api/gymgoers/:id",
			Handler:     uh.UpadateGymoer,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		},
		{
			Method:      http.MethodGet,
			Path:        "/v1/api/gymgoers/:id",
			Handler:     uh.GetGymGoerById,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:  http.MethodPost,
			Path:    "/v1/api/checkins",
			Handler: uh.CheckinUser,
		}, {
			Method:      http.MethodGet,
			Path:        "/view/gymgoers",
			Handler:     uh.GymGoersSimpleDetailByPhoneNumber,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/reports",
			Handler:     uh.Report,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/reportbydays",
			Handler:     uh.ReportByDate,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		},
	}

	uh.Routers = registerUser

	return uh.Routers
}
