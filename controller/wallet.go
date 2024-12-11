package controller

import (
	"miner/common/dto"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	walletService *service.WalletService
}

func NewWalletController() *WalletController {
	return &WalletController{
		walletService: service.NewWalletService(),
	}
}

func (c *WalletController) CreateWallet(ctx *gin.Context) {
	var req dto.CreateWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	walletID, err := c.walletService.CreateWallet(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": walletID,
		"msg":  "create wallet succes",
	})
}

func (c *WalletController) DeleteWallet(ctx *gin.Context) {
	var req dto.DeleteWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := c.walletService.DeleteWallet(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "delete walelt success",
	})
}

func (c *WalletController) UpdateWallet(ctx *gin.Context) {
	var req dto.UpdateWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := c.walletService.UpdateWallet(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "update walelt success",
	})
}

func (c *WalletController) GetUserAllWallet(ctx *gin.Context) {
	wallets, err := c.walletService.GetUserAllWallet(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": wallets,
	})
}

func (c *WalletController) GetUserWalletByID(ctx *gin.Context) {
	var req dto.GetUserWalletByIDReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	wallet, err := c.walletService.GetUserWalletByID(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": wallet,
	})
}
