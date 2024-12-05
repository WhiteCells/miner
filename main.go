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
	middleware.InitSession(ctx)
	userRoute := route.NewUserRoute()
	userRoute.InitUserRoute(ctx)
	farmRoute := route.NewFarmRoute()
	farmRoute.InitFarmRoute(ctx)
	ctx.Run(":8080")
}
