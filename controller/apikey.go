package controller

import (
	"miner/common/rsp"
	"miner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BscApiKeyController struct {
	bscApiKeyService *service.BscApiKeyService
}

func NewBscApiController() *BscApiKeyController {
	return &BscApiKeyController{
		bscApiKeyService: service.NewBscApiKeyService(),
	}
}

func (c *BscApiKeyController) GetTokenBalance(ctx *gin.Context) {
	address := ctx.Query("address")
	balance, err := c.bscApiKeyService.GetTokenBalance(ctx, address)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed get address token balance", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get token balance success", balance)
}
