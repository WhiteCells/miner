package main

import (
	"miner/middleware"
	"miner/route"
	"miner/utils"
	"strconv"

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

	route.NewUserRoute().InitUserRoute(ctx)
	route.NewFarmRoute().InitFarmRoute(ctx)
	route.NewMinerRoute().InitMinerRoute(ctx)
	route.NewFlightsheetRoute().InitFlightsheetRoute(ctx)
	route.NewWalletRoute().InitWalletRoute(ctx)
	route.NewAdminRoute().InitAdminRoute(ctx)
	route.NewHiveosRoute().InitHiveosRoute(ctx)

	port := strconv.Itoa(utils.Config.Server.Port)
	ctx.Run(":" + port)
}
