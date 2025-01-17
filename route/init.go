package route

import (
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

func Init(ctx *gin.Engine) {
	ctx.Use(middleware.OperLog())

	NewUserRoute().InitUserRoute(ctx)
	NewFarmRoute().InitFarmRoute(ctx)
	NewMinerRoute().InitMinerRoute(ctx)
	NewFlightsheetRoute().InitFlightsheetRoute(ctx)
	NewWalletRoute().InitWalletRoute(ctx)
	NewAdminRoute().InitAdminRoute(ctx)
	NewHiveosRoute().InitHiveosRoute(ctx)
}
