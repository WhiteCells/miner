package route

import (
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

func Init(ctx *gin.Engine) {
	NewSwaggerRoute().InitSwaggerRoute(ctx)

	NewHiveosRoute().InitHiveosRoute(ctx)

	ctx.Use(middleware.OperLog())

	NewUserRoute().InitUserRoute(ctx)
	NewFarmRoute().InitFarmRoute(ctx)
	NewMinerRoute().InitMinerRoute(ctx)
	NewFsRoute().InitFsRoute(ctx)
	NewWalletRoute().InitWalletRoute(ctx)
	NewAdminRoute().InitAdminRoute(ctx)
}
