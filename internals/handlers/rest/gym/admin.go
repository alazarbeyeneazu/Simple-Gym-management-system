package rest

import (
	"log"
	"net/http"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	encription "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/encryption"
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
	adusr, err := gh.admin.GetAdminById(ctx, models.AdminUsers{ID: adminId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}
	err = gh.appUser.DeleteUser(ctx, models.User{ID: adusr.UserId})
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
	var user models.CreateAdminRequest
	adminId, err := uuid.Parse(ctx.Param("adminId"))
	if err != nil {
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
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "admin": models.AdminUsers{}})
		log.Println(err)
		return
	}
	if user.Password != "" {

		user.Password, _ = encription.GenerateHashedPassword(user.Password)
		gh.appUser.UpdateUser(ctx, models.User{ID: admin.UserId, FirstName: user.FirstName, LastName: user.LastName, Password: user.Password})
	} else {
		gh.appUser.UpdateUser(ctx, models.User{FirstName: user.FirstName, LastName: user.LastName, Password: user.Password})
	}

	ctx.Redirect(http.StatusTemporaryRedirect, "/view/users")
	return
}
