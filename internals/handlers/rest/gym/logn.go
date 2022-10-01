package rest

import (
	"net/http"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/middlewares/token"
	encription "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/encryption"
	"github.com/gin-gonic/gin"
)

func (uh *restHandler) LoginUser(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	usr, err := uh.appUser.GetUserByPhoneNumber(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	if err := encription.CheckPassword(user.Password, usr.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect password"})
		return
	}
	paytomaker, err := token.NewPastoMaker("01234567890123456789012345678912")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "internal error can't initialize token maker"})
		return
	}
	tokenKey, _, err := paytomaker.CreateToken(usr.ID, time.Hour*4)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "internal error can't create token"})
		return
	}

	ctx.SetCookie("Athorization", tokenKey, int(time.Duration(time.Hour*5)), "/", "http://localhost:8282", false, false)

}
