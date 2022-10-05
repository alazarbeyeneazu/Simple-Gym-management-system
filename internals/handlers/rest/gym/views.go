package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (uh *restHandler) GetRegistrationPage(ctx *gin.Context) {

	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "user.html", usr)
}

func (uh *restHandler) GetLoginPage(ctx *gin.Context) {
	err, _ := ctx.Cookie("error")
	ctx.HTML(http.StatusOK, "index.html", gin.H{"error": err})

}
func (uh *restHandler) GetDashBoard(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "dashboard.html", usr)

}
func (uh *restHandler) GetRoles(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "roles.html", usr)

}
func (uh *restHandler) GetGym_goers(ctx *gin.Context) {
	var gymUsers []models.User
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}
	gymgoers, err := uh.gymgoers.GetAllGymGoers(ctx)
	if err != nil {
		return
	}
	for _, gmgoer := range gymgoers {
		guser := models.User{ID: gmgoer.UserId}
		gymuser, err := uh.appUser.GetUserById(ctx, guser)
		if err != nil {
			continue
		}
		gymUsers = append(gymUsers, gymuser)
	}
	payments, _ := uh.pymentUser.GetAllPyments(ctx)

	ctx.HTML(http.StatusOK, "gym-goers.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "gym_goers": gymUsers, "payments": payments})

}

func (uh *restHandler) GymGoersSimpleDetail(ctx *gin.Context) {
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

	var gymUsers []models.User
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	gymgoer := models.Gym_goers{UserId: uuids}

	gymusr, err := uh.gymgoers.GetGymGoerByUserId(ctx, gymgoer)
	if err != nil {
		log.Println(err)
		return
	}
	gymUserDetail, err := uh.appUser.GetUserById(ctx, models.User{ID: gymusr.UserId})
	if err != nil {
		log.Println(err)
		return
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}
	gymgoers, err := uh.gymgoers.GetAllGymGoers(ctx)
	if err != nil {
		return
	}
	for _, gmgoer := range gymgoers {
		guser := models.User{ID: gmgoer.UserId}
		gymuser, err := uh.appUser.GetUserById(ctx, guser)
		if err != nil {
			continue
		}
		gymUsers = append(gymUsers, gymuser)
	}
	payments, _ := uh.pymentUser.GetAllPyments(ctx)

	ctx.HTML(http.StatusOK, "gym-goers.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "gym_goers": gymUsers, "payments": payments, "gymgoerDetailFirstName": gymUserDetail.FirstName, "gymgoerDetailLastName": gymUserDetail.LastName, "qrid": gymusr.ID})

}

func (uh *restHandler) GetPayment(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
		return
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
		return
	}
	payments, err := uh.pymentUser.GetAllPyments(ctx)

	if err != nil {
		log.Println(err)
		return
	}

	ctx.HTML(http.StatusOK, "pyments.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "payments": payments, "Numberofdays": "", "Paymenttype": "", "Paymentfortype": "", "status": "Add", "method": "POST", "counter": len(payments)})

}
func (uh *restHandler) EditPayment(ctx *gin.Context) {
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
	pyment := models.PymentType{ID: uuids}
	payment, err := uh.pymentUser.GetPymentById(ctx, pyment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
		return
	}
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
		return
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
		return
	}
	payments, err := uh.pymentUser.GetAllPyments(ctx)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(payment)

	ctx.HTML(http.StatusOK, "pyments.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "payments": payments, "Numberofdays": payment.NumberOfDays, "Paymenttype": payment.PymentType, "Paymentfortype": payment.Payment, "status": "Update", "method": "PUT", "editting": uuids, "counter": len(payments)})

}
func (uh *restHandler) GetSetting(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}

	ctx.HTML(http.StatusOK, "setting.html", usr)

}
func (uh *restHandler) GetGym_goers_detail(ctx *gin.Context) {
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
	gymgoer, err := uh.gymgoers.GetGYmGorsById(ctx, models.Gym_goers{ID: uuids})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user, err := uh.appUser.GetUserById(ctx, models.User{ID: gymgoer.UserId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	startdate := fmt.Sprintf("%d/%d/%d", gymgoer.StartDate.Year(), gymgoer.StartDate.Month(), gymgoer.StartDate.Day())
	enddate := fmt.Sprintf("%d/%d/%d", gymgoer.EndDate.Year(), gymgoer.EndDate.Month(), gymgoer.EndDate.Day())
	ctx.HTML(http.StatusOK, "gym-goers-detail.html", gin.H{"error": "", "firstname": user.FirstName, "lastname": user.LastName, "createdAt": gymgoer.CreatedAt, "phonenumber": user.PhoneNumber, "startDate": startdate, "endDate": enddate, "creatorFirsName": gymgoer.CreatedByFirstName, "creatorLastName": gymgoer.CreatedByLastName, "creatorPhoneNumber": gymgoer.CreatedByPhoneNumber, "paidby": gymgoer.PaidBy})

}
