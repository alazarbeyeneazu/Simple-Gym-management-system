package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (uh *restHandler) LogOut(ctx *gin.Context) {

	ctx.SetCookie("Athorization", "", int(time.Duration(time.Hour*5)), "/", "http://localhost:8282", false, false)
	ctx.Redirect(http.StatusTemporaryRedirect, "/view/login")

}
