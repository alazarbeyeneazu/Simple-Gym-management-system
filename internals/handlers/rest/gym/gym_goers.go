package rest

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (gh *restHandler) RegisterGymGoer(ctx *gin.Context) {

	// user info
	users, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
		return
	}
	usr, err := gh.appUser.GetUserById(ctx, models.User{ID: users.(uuid.UUID)})
	if err != nil {
		log.Println(err)
		return
	}

	// end of user info ale should be refactore

	var gymgoer models.Gym_goerRequest
	if err := ctx.ShouldBind(&gymgoer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "user": models.Gym_goers{}})
		log.Println(err)
		return
	}
	if gymgoer.Start_date == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Select Start Date", "user": models.User{}})
		return
	}

	ur, _ := gh.appUser.GetUserByPhoneNumber(ctx, models.User{PhoneNumber: gymgoer.PhoneNumber})
	if ur.PhoneNumber != "" {
		registergymgoer, err := gh.gymgoers.GetGymGoerByUserId(ctx, models.Gym_goers{UserId: ur.ID})
		if err != nil {
			if err.Error() == "record not found" {
				log.Println("adding user ...")
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("can read user with this %s phone number", ur.PhoneNumber), "user": models.Gym_goers{}})
				return
			}
		}
		if registergymgoer.CreatedByFirstName != "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("user exist with this %s phone number", ur.PhoneNumber), "user": models.Gym_goers{}})
			log.Println("user exist")
			return
		}

	}
	if gymgoer.PaymentType == uuid.Nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Empty Payment Type", "user": models.User{}})
		return
	}

	pyment, err := gh.pymentUser.GetPymentById(ctx, models.PymentType{ID: gymgoer.PaymentType})

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "user": models.User{}})
		return
	}
	user := models.User{}
	if ur.PhoneNumber == "" {
		user, err = gh.appUser.RegisterUser(ctx, models.User{FirstName: gymgoer.FirstName, LastName: gymgoer.LastName, PhoneNumber: gymgoer.PhoneNumber, Password: "0000000011111111"})

	} else {
		user.ID = ur.ID
	}
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "user": models.User{}})
		return
	}

	newGymGoer := models.Gym_goers{}
	newGymGoer.UserId = user.ID
	newGymGoer.CreatedByFirstName = usr.FirstName
	newGymGoer.CreatedByLastName = usr.LastName
	newGymGoer.CreatedByPhoneNumber = usr.PhoneNumber
	newGymGoer.PaidAmount = pyment.Payment
	dateformate := strings.Split(gymgoer.Start_date, "-")

	year := dateformate[0]
	month := dateformate[1]

	day := strings.Split(dateformate[2], "T")[0]
	layout := "2006-01-02T15:04:05.000Z"
	str := year + "-" + month + "-" + day + "T11:00:26.371Z"
	startDate, err := time.Parse(layout, str)

	if err != nil {
		gh.appUser.DeleteUser(ctx, user)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "user": models.Gym_goers{}})
		return
	}

	endDate := startDate.Add(time.Hour * 24 * time.Duration(pyment.NumberOfDays))

	newGymGoer.EndDate = endDate
	newGymGoer.PaidBy = gymgoer.PaidBy
	newGymGoer.StartDate = startDate

	gymgoerResult, err := gh.gymgoers.RegisterGymGoer(ctx, newGymGoer)
	if err != nil {

		gh.appUser.DeleteUser(ctx, user)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "user": models.Gym_goers{}})
		return
	}
	gh.reports.CreateReport(ctx, models.ReportResponse{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		CreatedBy:  newGymGoer.CreatedByFirstName + " " + newGymGoer.CreatedByLastName,
		StartDate:  newGymGoer.StartDate,
		Amount:     pyment.Payment,
		PymentType: pyment.PymentType,
		EndDate:    newGymGoer.EndDate,
		PaidBy:     newGymGoer.PaidBy,
		CreatedAt:  time.Now(),
	})
	ctx.JSON(http.StatusOK, gin.H{"error": "", "user": gymgoerResult})

}

// UPDATE GYMGOERS
func (gh *restHandler) UpadateGymoer(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		log.Println("empty id ")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
		return
	}
	uuids, err := uuid.Parse(id)
	if err != nil {
		log.Println("invalid id ", id)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid gymgoer id"})
		return
	}
	if uuids == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		log.Println("empty id ")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
		return
	}

	// user info
	users, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
		return
	}
	userR, err := gh.appUser.GetUserById(ctx, models.User{ID: users.(uuid.UUID)})
	if err != nil {
		log.Println(err)
		return
	}

	// end of user info ale should be refactore

	var gymgoer models.Gym_goerRequest
	if err := ctx.ShouldBind(&gymgoer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "user": models.Gym_goers{}})
		log.Println(err)
		return
	}

	ur, _ := gh.appUser.GetUserById(ctx, models.User{ID: uuids})

	if ur.PhoneNumber == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Not found ", "user": models.Gym_goers{}})
		return
	}
	gymUser := models.User{ID: ur.ID, FirstName: gymgoer.FirstName, LastName: gymgoer.LastName, PhoneNumber: gymgoer.PhoneNumber}

	_, err = gh.appUser.UpdateUser(ctx, gymUser)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "user": models.User{}})
		return
	}

	if gymgoer.PaymentType == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Check Payment Type", "user": models.Gym_goers{}})
		return

	}
	if gymgoer.PaidBy == "" {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Check Payment Method", "user": models.User{}})
		return
	}
	pyment, err := gh.pymentUser.GetPymentById(ctx, models.PymentType{ID: gymgoer.PaymentType})
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "user": models.User{}})
		return
	}
	if gymgoer.Start_date == "" {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Check Start Date", "user": models.Gym_goers{}})
		return
	}
	startDate, err := time.Parse("2006-01-02", gymgoer.Start_date)
	log.Print("start date will be => ", startDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "user": models.Gym_goers{}})
		return
	}

	if time.Now().After(startDate) {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Starting Date", "user": models.Gym_goers{}})
		return
	}

	newGymGoer := models.Gym_goers{}

	endDate := startDate.Add(time.Hour * 24 * time.Duration(pyment.NumberOfDays))

	newGymGoer.EndDate = endDate
	newGymGoer.UserId = uuids
	newGymGoer.PaidBy = gymgoer.PaidBy
	newGymGoer.StartDate = startDate

	gymgoerResult, err := gh.gymgoers.UpdateGymGoer(ctx, newGymGoer)
	if err != nil {
		log.Println("payment creation error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "user": models.Gym_goers{}})
		return
	}
	gh.reports.CreateReport(ctx, models.ReportResponse{
		FirstName:  ur.FirstName,
		LastName:   ur.LastName,
		CreatedBy:  userR.FirstName + " " + userR.LastName,
		StartDate:  newGymGoer.StartDate,
		Amount:     pyment.Payment,
		PymentType: pyment.PymentType,
		EndDate:    newGymGoer.EndDate,
		PaidBy:     newGymGoer.PaidBy,
		CreatedAt:  time.Now(),
	})
	ctx.JSON(http.StatusOK, gin.H{"error": "", "user": gymgoerResult})

}

func (gh *restHandler) GetAllGymGoers(ctx *gin.Context) {

	gymgoerResult, err := gh.gymgoers.GetAllGymGoers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "user": models.Gym_goers{}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "user": gymgoerResult})

}

func (gh *restHandler) GetGymGoerById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		log.Println("empty id ")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
		return
	}
	uuids, err := uuid.Parse(id)
	if err != nil {
		log.Println("invalid id ", id)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid gymgoer id"})
		return
	}
	if uuids == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		log.Println("empty id ")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
		return
	}

	gymgoer := models.Gym_goers{ID: uuids}
	gymgoeruser, err := gh.gymgoers.GetGymGoerByUserId(ctx, gymgoer)
	if err != nil {
		log.Println("gym goer not found ")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "gym_goer not found "})
		return
	}
	user, err := gh.appUser.GetUserById(ctx, models.User{ID: gymgoeruser.UserId})
	if err != nil {
		log.Println("user not found with id", gymgoeruser.UserId)
	} else {
		err := gh.appUser.DeleteUser(ctx, user)
		if err != nil {
			log.Println("can not delete user with id ", err)
		}
	}

	result, err := gh.gymgoers.GetGYmGorsById(ctx, gymgoer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Println("user =>", user)
	log.Println("gym_goer =>", gymgoer)
	ctx.JSON(http.StatusOK, gin.H{"error": "", "gymgoer": result})

}

func (uh *restHandler) DeleteGymGoer(ctx *gin.Context) {

	id := ctx.Params.ByName("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
		return
	}
	uuids, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid pyment id"})
		return
	}
	if uuids == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
		return
	}
	gymGoer := models.Gym_goers{UserId: uuids}

	gm, err := uh.gymgoers.GetGymGoerByUserId(ctx, gymGoer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := uh.gymgoers.DeleteGymGoers(ctx, gm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	_, err = uh.admin.GetAdminByUserId(ctx, models.AdminUsers{UserId: uuids})
	if err != nil {

		if err.Error() == "record not found" {
			err = uh.appUser.DeleteUser(ctx, models.User{ID: uuids})
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}
		}
	}

	ctx.Redirect(http.StatusTemporaryRedirect, "/view/gym-goers")

}
