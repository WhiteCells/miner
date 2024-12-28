package controller

import (
	"miner/service"

	"github.com/gin-gonic/gin"
)

type HiveOsController struct {
	hiveOsService *service.HiveOsService
}

func NewHiveOsController() *HiveOsController {
	return &HiveOsController{
		hiveOsService: service.NewHiveOsService(),
	}
}

// 交互
func (c *HiveOsController) Interact(ctx *gin.Context) {
	// 解析传入的参数

	c.hiveOsService.Interact(ctx)
}
