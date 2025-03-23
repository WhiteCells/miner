package route

import (
	"miner/common/role"
	"miner/controller"
	"miner/middleware"

	"github.com/gin-gonic/gin"
)

type MinerRoute struct {
	minerController *controller.MinerController
}

func NewMinerRoute() *MinerRoute {
	return &MinerRoute{
		minerController: controller.NewMinerController(),
	}
}

func (mr *MinerRoute) InitMinerRoute(r *gin.Engine) {
	route := r.Group("/miner")
	route.Use(middleware.JWTAuth())
	// route.Use(middleware.IPAuth())
	route.Use(middleware.RoleAuth(role.User, role.Admin))
	route.Use(middleware.StatusAuth())
	{
		route.POST("", mr.minerController.CreateMiner)
		route.DELETE("", mr.minerController.DeleteMiner)
		route.PUT("", mr.minerController.UpdateMiner)
		route.PUT("/watchdog", mr.minerController.UpdateMinerWatchdog)
		route.PUT("/options", mr.minerController.UpdateMinerOptions)
		route.PUT("/autofan", mr.minerController.UpdateMinerAutofan)
		route.PUT("/wallet", mr.minerController.UpdateMinerWallet)
		route.GET("", mr.minerController.GetFarmAllMiner)
		route.GET("/info", mr.minerController.GetMinerByMinerID)
		route.PUT("/apply_fs", mr.minerController.ApplyFs)
		route.PUT("/transfer", mr.minerController.Transfer)
		// route.GET("/rig_conf", mr.minerController.GetRigConf)
		// watchdog
		route.POST("/watchdog", mr.minerController.SetWatchdog)
		route.GET("/watchdog", mr.minerController.GetWatchdog)
		// autofan
		route.POST("/autofan", mr.minerController.SetAutoFan)
		route.GET("/autofan", mr.minerController.GetAutoFan)
		// options
		route.POST("/options", mr.minerController.SetOptions)
		route.GET("/options", mr.minerController.GetOptions)
	}
}
