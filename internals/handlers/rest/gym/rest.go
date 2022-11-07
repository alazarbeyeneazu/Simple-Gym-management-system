package rest

import (
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/admin"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/authz"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/checkin"
	gymgoers "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/gym_goers"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/pyment"
	repor "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/report"
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

	//admin related
	RegisterAdmin(ctx *gin.Context)
	DeleteAdmin(ctx *gin.Context)
	EditAdmin(ctx *gin.Context)
	UpdateAdmin(ctx *gin.Context)

	//role related
	CreateRole(ctx *gin.Context)
	EditRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)

	//setting
	UpdateSetting(ctx *gin.Context)

	//checking
	CheckinUser(ctx *gin.Context)

	//checking by phone
	GymGoersSimpleDetailByPhoneNumber(ctx *gin.Context)
	Report(ctx *gin.Context)
	ReportByDate(ctx *gin.Context)

	//scanner
	Scanner(ctx *gin.Context)
}

type restHandler struct {
	appUser    user.UserService
	pymentUser pyment.PymentService
	auth       authz.AuthService
	admin      admin.AdminService
	checkin    checkin.CheckingService
	gymgoers   gymgoers.GymGoersService
	reports    repor.ReportService
	Routers    []routers.Router
}

func Init(user user.UserService, pyment pyment.PymentService, gymgoer gymgoers.GymGoersService, admin admin.AdminService, auth authz.AuthService, checkin checkin.CheckingService, report repor.ReportService) RestHandler {
	return &restHandler{appUser: user, pymentUser: pyment, gymgoers: gymgoer, admin: admin, auth: auth, checkin: checkin, reports: report}
}
