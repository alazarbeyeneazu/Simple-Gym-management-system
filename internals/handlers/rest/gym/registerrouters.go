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
			Path:        "/view/payment",
			Handler:     uh.GetPayment,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/setting",
			Handler:     uh.GetSetting,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/gym-goers-detail",
			Handler:     uh.GetGym_goers_detail,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:      http.MethodGet,
			Path:        "/view/logout",
			Handler:     uh.LogOut,
			MiddleWares: []gin.HandlerFunc{authn.AuthenticatRequest()},
		}, {
			Method:  http.MethodPut,
			Path:    "/v1/api/user",
			Handler: uh.RegisterUser,
		}, {
			Method:  http.MethodPost,
			Path:    "/v1/api/user",
			Handler: uh.LoginUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/v1/api/pyments",
			Handler: uh.GetAllPyments,
		}, {
			Method:  http.MethodPost,
			Path:    "/v1/api/pyments",
			Handler: uh.CreatePyment,
		}, {
			Method:  http.MethodPut,
			Path:    "/v1/api/pyments",
			Handler: uh.UpdatePyment,
		}, {
			Method:  http.MethodDelete,
			Path:    "/v1/api/pyments",
			Handler: uh.DeletePyment,
		},
	}

	uh.Routers = registerUser

	return uh.Routers
}
