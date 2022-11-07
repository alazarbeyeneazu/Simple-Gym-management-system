package rest

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type adminUser struct {
	AdminId     uuid.UUID `json:"admin_id"`
	PhoneNumber string    `json:"phone_number"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Role        string    `json:"role"`
}

func (uh *restHandler) GetRegistrationPage(ctx *gin.Context) {

	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}
	// users, err := uh.appUser.GetAllUsers(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	admins, err := uh.admin.GetAllAdmins(ctx)
	if err != nil {
		log.Println(err)
	}
	adminUserProfile := []adminUser{}
	for _, admin := range admins {
		usr, err := uh.appUser.GetUserById(ctx, models.User{ID: admin.UserId})
		if err != nil {
			log.Println(err)
			continue
		}
		role, _ := uh.auth.GetAssignRoleByUserId(ctx, models.UserRole{UserId: usr.ID})
		log.Println(role)
		adminUserProfile = append(adminUserProfile, adminUser{
			AdminId:     admin.ID,
			PhoneNumber: usr.PhoneNumber,
			FirstName:   usr.FirstName,
			LastName:    usr.LastName,
			Role:        role.RoleName,
		})
	}
	roles, err := uh.auth.GetAllRoles(ctx)
	if err != nil {
		log.Print(err)
		return
	}
	clearRole := []models.Role{}
	for _, role := range roles {
		if role.RoleName != "systemSupperadminUser" {
			clearRole = append(clearRole, role)
		}
	}
	ctx.HTML(http.StatusOK, "user.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "admins": adminUserProfile, "method": "POST", "sendTo": "", "actionType": "Register", "adminFirstName": "", "adminLastName": "", "adminPhoneNumber": "", "adminPassword": "", "roles": clearRole, "list": true})

}
func (uh *restHandler) EditAdmin(ctx *gin.Context) {

	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}
	// users, err := uh.appUser.GetAllUsers(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	admins, err := uh.admin.GetAllAdmins(ctx)
	if err != nil {
		log.Println(err)
	}
	adminUserProfile := []adminUser{}
	for _, admin := range admins {
		usr, err := uh.appUser.GetUserById(ctx, models.User{ID: usr.ID})
		if err != nil {
			log.Println(err)
			continue
		}
		role, _ := uh.auth.GetAssignRoleByUserId(ctx, models.UserRole{UserId: admin.UserId})

		adminUserProfile = append(adminUserProfile, adminUser{
			AdminId:     admin.ID,
			PhoneNumber: usr.PhoneNumber,
			FirstName:   usr.FirstName,
			LastName:    usr.LastName,
			Role:        role.RoleName,
		})
	}
	id, err := uuid.Parse(ctx.Param("adminId"))
	if err != nil {
		log.Println(er)
		return
	}
	adm, err := uh.admin.GetAdminById(ctx, models.AdminUsers{ID: id})
	if err != nil {
		log.Println(err)
		return
	}
	usrResponse, err := uh.appUser.GetUserById(ctx, models.User{ID: adm.UserId})
	if err != nil {
		log.Println(err)
		return
	}
	roles, err := uh.auth.GetAllRoles(ctx)
	if err != nil {
		log.Print(err)
		return
	}
	clearRole := []models.Role{}
	for _, role := range roles {
		if role.RoleName != "systemSupperadminUser" {
			clearRole = append(clearRole, role)
		}
	}
	ctx.HTML(http.StatusOK, "user.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "admins": adminUserProfile, "method": "PUT", "sendTo": fmt.Sprintf("/%v", id), "actionType": "Update", "adminFirstName": usrResponse.FirstName, "adminLastName": usrResponse.LastName, "adminPhoneNumber": usrResponse.PhoneNumber, "adminPassword": "", "roles": clearRole, "list": false})

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
	gmgoers, _ := uh.gymgoers.GetAllGymGoers(ctx)
	users, _ := uh.appUser.GetAllUsers(ctx)
	checkIn, _ := uh.checkin.GetAllCheckIns(ctx)
	checkins := []models.Checkins{}
	for _, checkin := range checkIn {
		if checkin.CheckedInDate.Year() == time.Now().Year() && checkin.CheckedInDate.Month() == time.Now().Month() && checkin.CheckedInDate.Day() == time.Now().Day() {
			checkins = append(checkins, checkin)
		}
	}
	usersToday := []models.User{}
	for _, usr := range users {
		for _, gym := range gmgoers {
			if usr.ID == gym.UserId {
				createdAt := gym.CreatedAt
				if createdAt.Year() == time.Now().Year() && createdAt.Month() == time.Now().Month() && createdAt.Day() == time.Now().Day() {
					usersToday = append(usersToday, usr)
				}

			}
		}
	}

	statics := make(map[string]int)
	for _, dates := range checkIn {
		statics[dates.CheckedInDate.Weekday().String()] = statics[dates.CheckedInDate.Weekday().String()] + 1
	}

	ctx.HTML(http.StatusOK, "dashboard.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "TotalGymGoeras": len(gmgoers), "totalUers": len(users), "today": len(checkins), "usersToday": len(usersToday), "monday": statics["Monday"], "Sunday": statics["Sunday"], "Tuesday": statics["Tuesday"], "Wednesday": statics["Wednesday"], "Thursday": statics["Thursday"], "Friday": statics["Friday"], "Saturday": statics["Saturday"]})

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
	permissions, err := uh.auth.GetAllPermission(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	roles, err := uh.auth.GetAllRoles(ctx)
	if err != nil {
		log.Print(err)
		return
	}
	clearRole := []models.Role{}
	for _, role := range roles {
		if role.RoleName != "systemSupperadminUser" {
			clearRole = append(clearRole, role)
		}
	}
	ctx.HTML(http.StatusOK, "roles.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "permission": permissions, "roles": clearRole})

}

func (uh *restHandler) EditRole(ctx *gin.Context) {
	user, er := ctx.Get("userid")
	if !er {
		log.Println("user id not found")
	}
	usr, err := uh.appUser.GetUserById(ctx, models.User{ID: user.(uuid.UUID)})
	if err != nil {
		log.Println(err)
	}
	roleName := ctx.Param("role")
	roles, err := uh.auth.GetRolesByName(ctx, models.Role{RoleName: roleName})
	if err != nil {
		log.Println(err)
		return
	}
	assignedPermissions := []models.Permission{}
	for _, role := range roles {

		prm, err := uh.auth.GetPermissionById(ctx, models.Permission{
			ID: role.PermissionID,
		})
		if err != nil {
			continue
		}
		assignedPermissions = append(assignedPermissions, prm)
	}

	if err != nil {
		log.Println(err)
		return
	}

	clearedAssingedPermissions := make(map[uuid.UUID]models.Permission)

	for _, clear := range assignedPermissions {
		clearedAssingedPermissions[clear.ID] = clear
	}

	ctx.HTML(http.StatusOK, "editRole.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "assignedPermission": clearedAssingedPermissions, "role": roles[0]})

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

	ctx.HTML(http.StatusOK, "gym-goers.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "gym_goers": gymUsers, "payments": payments, "actionName": "Register", "disabled": "none", "method": "POST", "url": "/v1/api/gymgoers"})

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

	ctx.HTML(http.StatusOK, "gym-goers.html", gin.H{"FirstName": usr.FirstName, "LastName": usr.LastName, "PhoneNumber": usr.PhoneNumber, "gym_goers": gymUsers, "payments": payments, "gymgoerDetailFirstName": gymUserDetail.FirstName, "gymgoerDetailLastName": gymUserDetail.LastName, "qrid": gymusr.ID, "gymgoerPhoneNumber": gymUserDetail.PhoneNumber, "actionName": "Update", "method": "PUT", "url": fmt.Sprintf("/v1/api/gymgoers/%s", gymUserDetail.ID), "startdateDisp": "none", "qrDisplayId": gymgoer.UserId})

}

func (uh *restHandler) GymGoersSimpleDetailByPhoneNumber(ctx *gin.Context) {
	phoneNumer := ctx.Request.URL.Query()["phone_number"][0]
	if len(phoneNumer) == 10 {
		phoneNumer = "+251" + phoneNumer[1:]
	}

	err := validation.Validate(&phoneNumer, validation.Required, validation.Length(13, 13))

	if phoneNumer == "" || err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/view/gym-goers")
		return
	}
	userd, err := uh.appUser.GetUserByPhoneNumber(ctx, models.User{PhoneNumber: phoneNumer})
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/view/gym-goers")
		return
	}
	gymgoer, err := uh.gymgoers.GetGymGoerByUserId(ctx, models.Gym_goers{UserId: userd.ID})
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/view/gym-goers")
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/view/gym-goers-detail/%s", gymgoer.ID))
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
	log.Println(user)
	ctx.HTML(http.StatusOK, "setting.html", usr)

}

// gymgoers detail
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
	checkedInAt := []models.DateResponse{}
	gymgoerCheckeIns, _ := uh.checkin.GetCheckedInByUserId(ctx, models.Checkins{UserId: user.ID})
	for _, checkAt := range gymgoerCheckeIns {
		checkedInAt = append(checkedInAt, models.DateResponse{

			Month:        checkAt.CheckedInDate.Month().String(),
			DayMonthYear: fmt.Sprintf("%d-%d-%d ", checkAt.CheckedInDate.Day(), int(checkAt.CheckedInDate.Month()), checkAt.CheckedInDate.Year()),
			Hour:         checkAt.CheckedInDate.Format(time.Kitchen),
		})

	}

	startdate := fmt.Sprintf("%d/%d/%d", gymgoer.StartDate.Year(), gymgoer.StartDate.Month(), gymgoer.StartDate.Day())
	enddate := fmt.Sprintf("%d/%d/%d", gymgoer.EndDate.Year(), gymgoer.EndDate.Month(), gymgoer.EndDate.Day())
	expired := time.Now().Before(gymgoer.EndDate)
	ctx.HTML(http.StatusOK, "gym-goers-detail.html", gin.H{"error": "", "firstname": user.FirstName, "lastname": user.LastName, "createdAt": gymgoer.CreatedAt, "phonenumber": user.PhoneNumber, "startDate": startdate, "endDate": enddate, "creatorFirsName": gymgoer.CreatedByFirstName, "creatorLastName": gymgoer.CreatedByLastName, "creatorPhoneNumber": gymgoer.CreatedByPhoneNumber, "paidby": gymgoer.PaidBy, "qrid": user.ID, "checkins": checkedInAt, "expired": expired})

}

func (uh *restHandler) Report(ctx *gin.Context) {
	reports := []models.HttpReportResponse{}
	gmgoers, _ := uh.reports.GetAllReports(ctx)
	for _, rep := range gmgoers {
		rp := models.HttpReportResponse{
			FirstName:  rep.FirstName,
			LastName:   rep.LastName,
			StartDate:  rep.StartDate.Format(time.ANSIC),
			EndDate:    fmt.Sprintf("%s/%s/%s/%s", rep.EndDate.Year(), rep.EndDate.Month(), rep.EndDate.Day(), rep.EndDate.Hour()),
			CreatedBy:  rep.CreatedBy,
			PymentType: rep.PymentType,
			Amount:     rep.Amount,
			PaidBy:     rep.PaidBy,
			CreatedAt:  rep.CreatedAt.Format(time.Kitchen),
		}
		reports = append(reports, rp)
	}

	ctx.HTML(http.StatusOK, "report.html", gin.H{"today": reports})

}

func (uh *restHandler) ReportByDate(ctx *gin.Context) {
	sdate := ctx.Request.URL.Query()["start_date"][0]
	pdate, err := time.Parse("2006-01-02", sdate)
	if err != nil {
		ctx.HTML(http.StatusOK, "report.html", gin.H{"hello": "world"})
	}

	gmgoers, _ := uh.reports.GetAllReports(ctx)

	usersToday := []models.HttpReportResponse{}

	for _, gym := range gmgoers {

		if gym.CreatedAt.After(pdate) {
			rp := models.HttpReportResponse{
				FirstName:  gym.FirstName,
				LastName:   gym.LastName,
				StartDate:  gym.StartDate.Format(time.ANSIC),
				EndDate:    gym.EndDate.Format(time.ANSIC),
				CreatedBy:  gym.CreatedBy,
				PymentType: gym.PymentType,
				Amount:     gym.Amount,
				PaidBy:     gym.PaidBy,
				CreatedAt:  gym.CreatedAt.Format(time.Kitchen),
			}
			usersToday = append(usersToday, rp)
		}

	}
	ctx.HTML(http.StatusOK, "report.html", gin.H{"today": usersToday})

}

func (uh *restHandler) Scanner(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "scanner.html", gin.H{"today": ""})

}
