package rest

import (
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/pyment"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/user"
	routers "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/routes"
	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	// user Views
	GetLoginPage(ctx *gin.Context)
	GetRegistrationPage(ctx *gin.Context)
	GetDashBoard(ctx *gin.Context)
	GetRoles(ctx *gin.Context)
	GetGym_goers(ctx *gin.Context)
	GetPayment(ctx *gin.Context)
	GetSetting(ctx *gin.Context)
	GetGym_goers_detail(ctx *gin.Context)

	//router
	StartRoutes() []routers.Router

	//user related
	RegisterUser(ctx *gin.Context)
	LogOut(ctx *gin.Context)

	// pyments
	CreatePyment(ctx *gin.Context)
	GetAllPyments(ctx *gin.Context)
	UpdatePyment(ctx *gin.Context)
	DeletePyment(ctx *gin.Context)
}

type restHandler struct {
	appUser    user.UserService
	pymentUser pyment.PymentService
	Routers    []routers.Router
}

func Init(user user.UserService, pyment pyment.PymentService) RestHandler {
	return &restHandler{appUser: user, pymentUser: pyment}
}
