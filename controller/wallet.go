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

// CreateWallet 创建钱包
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

// DeleteWallet 删除钱包
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

// UpdateWallet 更新钱包
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

// GetUserAllWallet 获取用户所有钱包
func (c *WalletController) GetAllWallet(ctx *gin.Context) {
	coin := ctx.Query("coin")
	wallets, err := c.walletService.GetAllWalletByCoin(ctx, coin)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get user all wallet success", wallets)
}

// GetUserWalletByID 通过钱包 ID 获取指定钱包
func (c *WalletController) GetUserWalletByID(ctx *gin.Context) {
	walletID := ctx.Param("wallet_id")
	wallet, err := c.walletService.GetUserWalletByID(ctx, walletID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "get user wallet by id success", wallet)
}

// GetAllCoin
func (c *WalletController) GetAllWalletAllCoin(ctx *gin.Context) {
	coins, err := c.walletService.GetAllWalletAllCoin(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "get user all wallet success", coins)
}
