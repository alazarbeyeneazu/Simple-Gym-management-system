package authn

import (
	"net/http"

	token "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/middlewares/token"
	"github.com/gin-gonic/gin"
)

// var authorizationType = "bearer"

func AuthenticatRequest() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		maker, err := token.NewPastoMaker("01234567890123456789012345678912")
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:8282/view/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		}
		authHeader, err := ctx.Cookie("Athorization")
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:8282/view/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		}
		if authHeader == "" {
			ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:8282/view/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		}

		pload, err := maker.VerifyToken(authHeader)

		if err != nil {
			ctx.SetCookie("error", "login session expired", 2, "/", "localhost", false, false)
			ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:8282/view/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		}
		ctx.Set("userid", pload.UserId)
		ctx.Next()

	}
}
