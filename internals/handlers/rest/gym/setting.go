package rest

import (
	"log"
	"net/http"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	encription "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/encryption"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SettingRequest struct {
	OldPassword   string `json:"old_password"`
	NewPassword   string `json:"new_password"`
	ReNewPassword string `json:"re_new_password"`
}

func (rh *restHandler) UpdateSetting(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := rh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
		return
	}

	var settingRequest SettingRequest
	if err := ctx.ShouldBind(&settingRequest); err != nil {

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		return
	}
	if settingRequest.OldPassword != "" {
		if settingRequest.NewPassword != settingRequest.ReNewPassword {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "new password and confirmed password not equal"})
			return

		} else if len(settingRequest.NewPassword) < 8 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "short password"})
			return
		} else {
			password, err := encription.GenerateHashedPassword(settingRequest.NewPassword)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to hash password"})
				return
			} else {

				if err := encription.CheckPassword(settingRequest.OldPassword, usr.Password); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Old Password"})
					return
				}
				_, err = rh.appUser.UpdateUser(ctx, models.User{ID: usr.ID, Password: password})
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to update password"})
					return
				}
			}
		}

	}
	ctx.JSON(http.StatusOK, gin.H{"error": ""})

}
