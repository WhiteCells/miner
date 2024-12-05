package controller

import (
	"miner/common/dto"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FarmController struct {
	farmService *service.FarmService
}

func NewFarmController() *FarmController {
	return &FarmController{
		farmService: service.NewFarmService(),
	}
}

func (c *FarmController) CreateFarm(ctx *gin.Context) {
	var req dto.CreateFarmReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	farmID, err := c.farmService.CreateFarm(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create farm success",
		"farm_id": farmID,
	})
}
