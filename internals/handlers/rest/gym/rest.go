package rest

import (
	gymgoers "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/gym_goers"
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
	GetPymentById(ctx *gin.Context)
	UpdatePyment(ctx *gin.Context)
	DeletePyment(ctx *gin.Context)

	//gym_goers related
	RegisterGymGoer(ctx *gin.Context)
	GetAllGymGoers(ctx *gin.Context)
	GetGymGoerById(ctx *gin.Context)
}

type restHandler struct {
	appUser    user.UserService
	pymentUser pyment.PymentService
	gymgoers   gymgoers.GymGoersService
	Routers    []routers.Router
}

func Init(user user.UserService, pyment pyment.PymentService, gymgoer gymgoers.GymGoersService) RestHandler {
	return &restHandler{appUser: user, pymentUser: pyment, gymgoers: gymgoer}
}
