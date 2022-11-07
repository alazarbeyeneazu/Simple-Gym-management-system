package rest

import (
	"net/http"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
)

func (uh *restHandler) RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"user":   models.User{},
			"status": "not registered",
		})
		return
	}
	respuser, err := uh.appUser.RegisterUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":  err.Error(),
			"user":   models.User{},
			"status": "not registered",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":  "",
		"user":   respuser,
		"status": "registered",
	})
}

func (uh *restHandler) GetAllUsers(ctx *gin.Context) {

	users, err := uh.appUser.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err, "users": []models.User{}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "users": users})

}
