package rest

import (
	"log"
	"net/http"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (gh *restHandler) RegisterAdmin(ctx *gin.Context) {

	var admin models.CreateAdminRequest

	if err := ctx.ShouldBind(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}
	adminresult, err := gh.admin.RegisterAdmin(ctx, admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "", "admin": adminresult})

}
func (gh *restHandler) DeleteAdmin(ctx *gin.Context) {

	adminId, err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}

	err = gh.admin.DeleteAdmin(ctx, models.AdminUsers{ID: adminId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/view/users")

}
func (gh *restHandler) UpdateAdmin(ctx *gin.Context) {
	var user models.User
	adminId, err := uuid.Parse(ctx.Param("adminId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}
	admin, err := gh.admin.GetAdminById(ctx, models.AdminUsers{ID: adminId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}
	user.ID = admin.UserId
	gh.appUser.UpdateUser(ctx, user)

}
