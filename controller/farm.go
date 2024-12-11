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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	farm, err := c.farmService.CreateFarm(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": farm,
		"msg":  "create farm success",
	})
}

func (c *FarmController) DeleteFarm(ctx *gin.Context) {
	var req dto.DeleteFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := c.farmService.DeleteFarm(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "delete success",
	})
}

func (c *FarmController) UpdateFarm(ctx *gin.Context) {
	var req dto.UpdateFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := c.farmService.UpdateFarm(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}

func (c *FarmController) GetUserAllFarm(ctx *gin.Context) {
	farms, err := c.farmService.GetUserAllFarmInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": farms,
	})
}

func (c *FarmController) ApplyFlightSheet(ctx *gin.Context) {
	var req dto.ApplyFarmFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := c.farmService.ApplyFlightSheet(ctx, &req); err != nil {
		return
	}
}
