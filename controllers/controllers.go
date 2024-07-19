package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/koshigi/grocery-tracker-backend/models"
	"gorm.io/gorm"
)

type ListController struct {
	DB *gorm.DB
}

func NewListController(DB *gorm.DB) ListController {
	return ListController{DB}
}

func (uc *ListController) GetListById(ctx *gin.Context) {
	list_id := ctx.Param("list_id")
	result := uc.DB.First(&models.List{}, list_id)

	//TODO: fix this error Error #01: json: unsupported type: func() time.Time
	// unable to marshal the time data type, maybe specify formatting in models?
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"list": result}})
}
