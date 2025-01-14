package controller

import "github.com/gin-gonic/gin"

type CoinController struct {
}

func NewCoinController() *CoinController {
	return &CoinController{}
}

func (c *CoinController) GetCoin(ctx *gin.Context) {

}
