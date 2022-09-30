package rest

import (
	"log"
	"net/http"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (uh *userHanlder) GetRegistrationPage(ctx *gin.Context) {

	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "user.html", usr)
}

func (uh *userHanlder) GetLoginPage(ctx *gin.Context) {
	err, _ := ctx.Cookie("error")
	ctx.HTML(http.StatusOK, "index.html", gin.H{"error": err})

}
func (uh *userHanlder) GetDashBoard(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "dashboard.html", usr)

}
func (uh *userHanlder) GetRoles(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "roles.html", usr)

}
func (uh *userHanlder) GetGym_goers(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "gym-goers.html", usr)

}
func (uh *userHanlder) GetPayment(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "pyments.html", usr)

}
func (uh *userHanlder) GetSetting(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "setting.html", usr)

}
func (uh *userHanlder) GetGym_goers_detail(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "gym-goers-detail.html", gin.H{"error": ""})

}
