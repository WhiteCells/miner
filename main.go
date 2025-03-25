package main

import (
	"flag"
	"miner/dao/redis"
	model "miner/model/migrate"
	"miner/route"
	"miner/settlement"
	"miner/utils"

	"github.com/gin-gonic/gin"
)

var configPath string

// go run main.go -f config.dev.yml
func init() {
	flag.StringVar(&configPath, "c", "./config.dev.yml", "path to config file")
	flag.Parse()

	utils.InitConfig(configPath, "yml")
	utils.InitJWT()
	utils.InitLogger()
	utils.InitRDB()
	redis.Init()
	utils.InitDB()
	model.Migrate()
	settlement.InitCronJob()
}

func main() {
	gin.SetMode(utils.Config.Server.Mode)
	ctx := gin.Default()
	route.Init(ctx)
	if err := ctx.Run(utils.GeneratePort()); err != nil {
		utils.Logger.Error("" + err.Error())
	}
}
