package rest

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (uh *restHandler) CheckinUser(ctx *gin.Context) {
	left := ""
	userid := ctx.Param("userid")

	uuid, err := uuid.Parse(userid)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Fake User ID", "user_first_name": "", "user_last_name": "", "left_days": "", "isChackedIn": ""})
		return
	}

	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: uuid})
	if err != nil {

		ctx.JSON(http.StatusNotFound, gin.H{"error": "Fake User ID", "user_first_name": "", "user_last_name": "", "left_days": "", "isChackedIn": ""})
		return
	}
	log.Print(usr.FirstName)
	gymgoer, err := uh.gymgoers.GetGymGoerByUserId(ctx, models.Gym_goers{UserId: usr.ID})
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found", "user_first_name": "", "user_last_name": "", "left_days": "", "isChackedIn": ""})
		return

	}

	yearLeft := gymgoer.EndDate.Year() - time.Now().Year()
	monthLeft := int(gymgoer.EndDate.Month()) - int(time.Now().Month())
	dayLeft := gymgoer.EndDate.Day() - time.Now().Day()

	if yearLeft >= 1 {

		if yearLeft > 1 {
			left = fmt.Sprintf("You have %v years ", yearLeft)
		} else {
			left = "You have this year "
		}
	} else if monthLeft >= 1 {
		if monthLeft > 1 {
			left = fmt.Sprintf("You have %v Months ", monthLeft)
		} else {
			left = fmt.Sprintf("You have %v Month ", monthLeft)
		}

	} else if dayLeft >= 1 {
		if dayLeft > 0 {
			left = fmt.Sprintf("You have  %v days", dayLeft)
		} else {
			left = fmt.Sprintf("You have  %v day", dayLeft)
		}

	} else {
		left = "expired user Id"
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "your service is expired", "user_first_name": "", "user_last_name": "", "left_days": "", "isChackedIn": ""})
		return
	}

	chcking, err := uh.checkin.CheckingUser(ctx, models.Checkins{UserId: usr.ID})
	if err != nil {
		if err.Error() == "already checkedIn" {
			chcking.UserFirstName = usr.FirstName
			chcking.UserLastName = usr.LastName
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "user_first_name": "", "left_days": "", "user_last_name": "", "isChackedIn": ""})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"user": models.CheckinResponse{}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "", "user_first_name": usr.FirstName, "left_days": left, "user_last_name": usr.LastName, "isChackedIn": chcking.IsChackedIn})

}
