package controller

import "github.com/gin-gonic/gin"

type MinesoftController struct {
}

func NewMinesoftController() *MinerController {
	return &MinerController{}
}

func (c *MinerController) GetMinesoft(ctx *gin.Context) {

}
