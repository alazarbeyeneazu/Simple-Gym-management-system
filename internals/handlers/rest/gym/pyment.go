package rest

import (
	"net/http"

	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/constants/models"
	"github.com/gin-gonic/gin"
)

func (uh *restHandler) GetAllPyments(ctx *gin.Context) {
	pyments, err := uh.pymentUser.GetAllPyments(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "payments": pyments})
}

func (uh *restHandler) CreatePyment(ctx *gin.Context) {
	var pyment models.PymentType
	if err := ctx.ShouldBind(&pyment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	pmnt, err := uh.pymentUser.CreatePyment(ctx, pyment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "payment": pmnt})

}

func (uh *restHandler) UpdatePyment(ctx *gin.Context) {
	var pyment models.PymentType
	if err := ctx.ShouldBind(&pyment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	pmnt, err := uh.pymentUser.UpdatePyment(ctx, pyment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "", "payment": pmnt})
}

func (uh *restHandler) DeletePyment(ctx *gin.Context) {
	var pyment models.PymentType
	if err := ctx.ShouldBind(&pyment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err := uh.pymentUser.DeletePyment(ctx, pyment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": ""})
}
