package middleware

import (
	"miner/common/rsp"
	"miner/common/status"
	"miner/model"
	"miner/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 检查注册开关
func CheckSwitchRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var system model.System
		if err := utils.DB.First(&system).Error; err != nil {
			rsp.Error(ctx, http.StatusInternalServerError, "'system' db error", nil)
			ctx.Abort()
			return
		}

		if system.SwitchRegister != status.RegisterOn {
			rsp.Error(ctx, http.StatusBadRequest, "register closed", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
