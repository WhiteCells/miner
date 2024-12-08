package controller

import (
	"miner/common/dto"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MinerController struct {
	minerService *service.MinerService
}

func NewMinerController() *MinerController {
	return &MinerController{
		minerService: service.NewMinerService(),
	}
}

func (c *MinerController) CreateMiner(ctx *gin.Context) {
	var req dto.CreateMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.minerService.CreateMiner(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

func (c *MinerController) DeleteMiner(ctx *gin.Context) {
	var req dto.DeleteMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.minerService.DeleteMiner(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete miner success",
	})
}

func (c *MinerController) UpdateMiner(ctx *gin.Context) {
	var req dto.UpdateMinerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.minerService.UpdateMiner(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update miner success",
	})
}

func (c *MinerController) GetUserAllMinerInFarm(ctx *gin.Context) {
	var req dto.GetUserAllMinerInFarmReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	miners, err := c.minerService.GetUserAllMinerInFarm(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"miners": miners,
	})
}
