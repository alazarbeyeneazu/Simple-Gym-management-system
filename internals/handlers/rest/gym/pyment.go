package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (uh *restHandler) GetAllPyments(ctx *gin.Context) {

	pyments, err := uh.pymentUser.GetAllPyments(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "payments": pyments})
}

func (uh *restHandler) CreatePyment(ctx *gin.Context) {

	// user info
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

	// end of user info ale should be refactore

	var pymentRequst models.PymentTypeRequest
	if err := ctx.ShouldBind(&pymentRequst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pyment := models.PymentType{}
	numberofdays, err := strconv.Atoi(pymentRequst.NumberOfDays)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid number of days"})
		return
	}
	pyment.NumberOfDays = int64(numberofdays)
	pyment.PymentType = pymentRequst.PymentType
	pyment.Payment = pymentRequst.Payment
	pyment.CreatedByFirstName = usr.FirstName
	pyment.CreatedByLastName = usr.LastName
	pmnt, err := uh.pymentUser.CreatePyment(ctx, pyment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "payment": pmnt})

}

func (uh *restHandler) UpdatePyment(ctx *gin.Context) {

	// user info
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

	// end of user info ale should be refactore

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
	var pymentRequst models.PymentTypeRequest
	if err := ctx.ShouldBind(&pymentRequst); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nuberofdays, err := strconv.Atoi(pymentRequst.NumberOfDays)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Number of days"})
		return
	}
	pyment := models.PymentType{}

	pyment.ID = uuids
	pyment.CreatedByFirstName = usr.FirstName
	pyment.CreatedByLastName = usr.LastName
	pyment.NumberOfDays = int64(nuberofdays)
	pyment.PymentType = pymentRequst.PymentType
	pyment.Payment = pymentRequst.Payment
	pmnt, err := uh.pymentUser.UpdatePyment(ctx, pyment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "payment": pmnt})
}

func (uh *restHandler) DeletePyment(ctx *gin.Context) {

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

	err = uh.pymentUser.DeletePyment(ctx, pyment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/view/payment")

}
func (uh *restHandler) GetPymentById(ctx *gin.Context) {
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
	result, err := uh.pymentUser.GetPymentById(ctx, pyment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "payment": result})
}
