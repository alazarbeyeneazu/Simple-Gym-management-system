package rest

import (
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/user"
	routers "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/routes"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetLoginPage(ctx *gin.Context)
	StartRoutes() []routers.Router
	RegisterUser(ctx *gin.Context)
	GetRegistrationPage(ctx *gin.Context)
	GetDashBoard(ctx *gin.Context)
	GetRoles(ctx *gin.Context)
	GetGym_goers(ctx *gin.Context)
	GetPayment(ctx *gin.Context)
	GetSetting(ctx *gin.Context)
	GetGym_goers_detail(ctx *gin.Context)
	LogOut(ctx *gin.Context)
}

type userHanlder struct {
	appUser user.UserService
	Routers []routers.Router
}

func Init(user user.UserService) UserHandler {
	return &userHanlder{appUser: user}
}
