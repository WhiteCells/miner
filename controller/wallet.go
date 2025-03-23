package controller

import (
	"miner/common/dto"
	"miner/common/params"
	"miner/common/rsp"
	"miner/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	walletService *services.WalletService
}

func NewWalletController() *WalletController {
	return &WalletController{
		walletService: services.NewWalletService(),
	}
}

// 创建钱包
func (c *WalletController) CreateWallet(ctx *gin.Context) {
	var req dto.CreateWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	err := c.walletService.CreateWallet(ctx, userID, req.CoinID, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "create wallet succes", nil)
}

// DeleteWallet 删除钱包
func (c *WalletController) DeleteWallet(ctx *gin.Context) {
	var req dto.DeleteWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.walletService.DelWallet(ctx, userID, req.WalletID); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "delete wallet success", nil)
}

// 更新钱包
func (c *WalletController) UpdateWallet(ctx *gin.Context) {
	var req dto.UpdateWalletReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userID := ctx.GetInt("user_id")

	if err := c.walletService.UpdateWallet(ctx, userID, req.WalletID, req.UpdateInfo); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "update wallet success", nil)
}

// 获取对应 coin 的钱包
func (m *WalletController) GetWalletByCoinID(ctx *gin.Context) {
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}
	userID := ctx.GetInt("user_id")
	wallets, total, err := m.walletService.GetWalletsByUserID(ctx, userID, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.DBQuerySuccess(ctx, http.StatusOK, "get user all wallet success", wallets, total)
}

// 获取用户所有钱包
func (c *WalletController) GetWalletByUserID(ctx *gin.Context) {
	var params params.PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  params.PageNum,
		"page_size": params.PageSize,
	}
	userID := ctx.GetInt("user_id")

	wallets, total, err := c.walletService.GetWalletsByUserID(ctx, userID, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "get user all wallet success", wallets, total)
}

// 通过钱包 ID 获取指定钱包
func (c *WalletController) GetWalletByWalletID(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("wallet_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	userID := ctx.GetInt("user_id")
	wallet, err := c.walletService.GetWalletByWalletID(ctx, userID, walletID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "get user wallet by id success", wallet)
}

// GetAllCoin
// func (c *WalletController) GetAllWalletAllCoin(ctx *gin.Context) {
// 	coins, err := c.walletService.GetWalletAllCoin(ctx)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}

// 	rsp.QuerySuccess(ctx, http.StatusOK, "get user all wallet success", coins)
// }
