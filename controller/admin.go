package controller

import (
	"miner/service"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	admineService *service.AdmineService
}

func NewAdmineController() *AdminController {
	return &AdminController{
		admineService: service.NewAdmineService(),
	}
}

// AdminGetUser 获取所有用户
func (ar *AdminController) AdminGetUser(ctx *gin.Context) {

}

// AdminGetOperLog 获取所有用户操作日志
func (ar *AdminController) AdminGetOperLog(ctx *gin.Context) {

}

// AdmineSetPointsReward 调整邀请积分奖励
func (ar *AdminController) AdmineSetPointsReward(ctx *gin.Context) {

}

// AdmineSetRechargeReward 调整充值积分奖励
func (ar *AdminController) AdmineSetRechargeReward(ctx *gin.Context) {

}
