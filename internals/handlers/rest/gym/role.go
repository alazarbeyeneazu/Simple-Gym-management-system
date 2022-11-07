package rest

import (
	"log"
	"net/http"
	"strings"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (gh *restHandler) CreateRole(ctx *gin.Context) {
	// roles := models.CreateRoleRequest{}
	ctx.Request.ParseForm()

	// if err := ctx.ShouldBind(&roles); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "role": models.Role{}})
	// 	log.Println(err)
	// 	return
	// }
	response := strings.Split(ctx.Request.Form.Encode(), "role_name=")

	permissions := strings.Split(strings.Join(strings.Split(response[0], "&"), "permissions="), "permissions=")

	for _, permissionId := range permissions {
		if permissionId == "" {
			continue
		}
		log.Println(permissionId)
		permsision := uuid.MustParse(permissionId)
		_, err := gh.auth.CreateRole(ctx, models.Role{
			RoleName:     response[1],
			PermissionID: permsision,
		})
		if err != nil {
			log.Println(err)
			continue
		}

	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/view/roles")
}
func (gh *restHandler) UpdateRole(ctx *gin.Context) {
	// roles := models.CreateRoleRequest{}
	ctx.Request.ParseForm()

	response := strings.Split(ctx.Request.Form.Encode(), "role_name=")

	permissions := strings.Split(strings.Join(strings.Split(response[0], "&"), "permissions="), "permissions=")
	gh.auth.DeleteRole(ctx, models.Role{RoleName: response[1]})

	for _, permissionId := range permissions {
		if permissionId == "" {
			continue
		}

		permsision := uuid.MustParse(permissionId)
		_, err := gh.auth.CreateRole(ctx, models.Role{
			RoleName:     response[1],
			PermissionID: permsision,
		})
		if err != nil {
			log.Println(err)
			continue
		}

	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/view/roles")
}
func (gh *restHandler) DeleteRole(ctx *gin.Context) {

	roleId := ctx.Param("roleId")

	if err := gh.auth.DeleteRole(ctx, models.Role{RoleName: roleId}); err != nil {
		log.Println(err)
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/view/roles")

}
