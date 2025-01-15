package controller

import "github.com/gin-gonic/gin"

type CoinController struct {
}

func NewCoinController() *CoinController {
	return &CoinController{}
}

func (c *CoinController) SetCoin(ctx *gin.Context) {

}

func (c *CoinController) BatchSetCoin(ctx *gin.Context) {

}

func (c *CoinController) DelCoin(ctx *gin.Context) {

}

func (c *CoinController) GetCoin(ctx *gin.Context) {

}
