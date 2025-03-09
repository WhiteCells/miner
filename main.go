package main

import (
	"fmt"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/route"
	"miner/settlement"
	"miner/utils"

	"github.com/gin-gonic/gin"
)

func initialize() error {
	//if err := utils.InitConfig("./config.dev.yml", "yml"); err != nil {
	//	return err
	//}
	if err := utils.InitConfig("./config.yml", "yml"); err != nil {
		return err
	}
	utils.InitJWT()
	if err := utils.InitLogger(); err != nil {
		return err
	}
	if err := utils.InitRDB(); err != nil {
		return err
	}
	if err := redis.Init(); err != nil {
		return err
	}
	if err := utils.InitDB(); err != nil {
		return err
	}
	if err := mysql.Init(); err != nil {
		return err
	}
	settlement.InitCronJob()
	return nil
}

func main() {
	fmt.Println("哈哈哈哈哈哈哈")

	if err := initialize(); err != nil {
		utils.Logger.Error(err.Error())
		return
	}
	gin.SetMode(utils.Config.Server.Mode)
	ctx := gin.Default()
	route.Init(ctx)
	if err := ctx.Run(utils.GeneratePort()); err != nil {
		utils.Logger.Error("" + err.Error())
	}
}
