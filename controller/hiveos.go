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

// 发送命令
func (c *HiveOsController) SendCmd(ctx *gin.Context) {

}

// 获取命令结果
func (c *HiveOsController) GetCmdRes(ctx *gin.Context) {

}

// 设置配置
func (c *HiveOsController) SetConfig(ctx *gin.Context) {

}

// 获取设置配置结果
func (c *HiveOsController) GetConfigRes(ctx *gin.Context) {

}
