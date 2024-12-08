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
		"farm_id": farmID,
		"message": "create farm success",
	})
}

func (c *FarmController) GetUserAllFarm(ctx *gin.Context) {
	farms, err := c.farmService.GetUserAllFarmInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"farms": farms,
	})
}

func (c *FarmController) DeleteFarm(ctx *gin.Context) {
	var req dto.DeleteFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.farmService.DeleteFarm(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

func (c *FarmController) UpdateFarm(ctx *gin.Context) {
	var req dto.UpdateFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.farmService.UpdateFarm(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}
