package main

import (
	"miner/middleware"
	"miner/route"
	"miner/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitConfig("./config.yml", "yml")
	utils.InitJWT()
	utils.InitLogger()
	utils.InitRDB()
	utils.InitDB()

	gin.SetMode(utils.Config.Server.Mode)
	ctx := gin.Default()
	ctx.Use(middleware.OperLog())
	userRoute := route.NewUserRoute()
	userRoute.InitUserRoute(ctx)
	farmRoute := route.NewFarmRoute()
	farmRoute.InitFarmRoute(ctx)
	minerRoute := route.NewMinerRoute()
	minerRoute.InitMinerRoute(ctx)
	flightsheetRoute := route.NewFlightsheetRoute()
	flightsheetRoute.InitFlightsheetRoute(ctx)
	walletRoute := route.NewWalletRoute()
	walletRoute.InitWalletRoute(ctx)
	adminRoute := route.NewAdminRoute()
	adminRoute.InitAdminRoute(ctx)

	ctx.Run(":8080")
}
