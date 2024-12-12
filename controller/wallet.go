package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
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
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	wallet, err := c.walletService.CreateWallet(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "create wallet succes", wallet)
}

func (c *WalletController) DeleteWallet(ctx *gin.Context) {
	var req dto.DeleteWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.walletService.DeleteWallet(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "delete wallet success", nil)
}

func (c *WalletController) UpdateWallet(ctx *gin.Context) {
	var req dto.UpdateWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := c.walletService.UpdateWallet(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update wallet success", nil)
}

func (c *WalletController) GetUserAllWallet(ctx *gin.Context) {
	wallets, err := c.walletService.GetUserAllWallet(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "get user all wallet success", wallets)
}

func (c *WalletController) GetUserWalletByID(ctx *gin.Context) {
	var req dto.GetUserWalletByIDReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	wallet, err := c.walletService.GetUserWalletByID(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "get user wallet by id success", wallet)
}
