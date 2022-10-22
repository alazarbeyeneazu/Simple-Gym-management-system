package rest

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
)

func (uh *restHandler) CheckinUser(ctx *gin.Context) {

	var user models.Checkins
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request", "user": models.CheckinResponse{}})
		return
	}
	log.Println(user)
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.UserId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found", "user": models.CheckinResponse{}})
		return
	}

	gymgoer, err := uh.gymgoers.GetGymGoerByUserId(ctx, models.Gym_goers{UserId: usr.ID})
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "already checkedIn", "user": models.CheckinResponse{
			IsChackedIn: "Not Gym User",
		}})
		return

	}
	chcking, err := uh.checkin.CheckingUser(ctx, models.Checkins{UserId: usr.ID})
	if err != nil {
		if err.Error() == "already checkedIn" {
			yearLeft := gymgoer.EndDate.Year() - time.Now().Year()
			monthLeft := int(gymgoer.EndDate.Month()) - int(time.Now().Month())
			dayLeft := gymgoer.EndDate.Day() - time.Now().Day()
			if yearLeft >= 1 {

				if yearLeft > 1 {
					chcking.UserNumberOfDayLeft = fmt.Sprintf("You have %v years ", yearLeft)
				} else {
					chcking.UserNumberOfDayLeft = fmt.Sprintf("You have %v year ", yearLeft)
				}
			} else if monthLeft >= 1 {
				if monthLeft > 1 {
					chcking.UserNumberOfDayLeft = fmt.Sprintf("You have %v Months ", monthLeft)
				} else {
					chcking.UserNumberOfDayLeft = fmt.Sprintf("You have %v Month ", monthLeft)
				}

			} else if dayLeft >= 1 {
				if dayLeft > 1 {
					chcking.UserNumberOfDayLeft = fmt.Sprintf("You have  %v days", dayLeft)
				} else {
					chcking.UserNumberOfDayLeft = fmt.Sprintf("You have  %v day", dayLeft)
				}

			} else {
				chcking.UserNumberOfDayLeft = "expired user Id"
			}

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "already checkedIn", "user": chcking})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found", "user": models.CheckinResponse{}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "user not found", "user": models.CheckinResponse{

		UserFirstName: usr.FirstName,
		UserLastName:  usr.LastName,
		CheckedInDate: chcking.CheckedInDate,
	}})

}
