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
		route.GET("", mr.minerController.GetFarmAllMiner)
		route.GET("/:farm_id/:miner_id", mr.minerController.GetMinerByID)
		route.GET("/info", mr.minerController.GetMinerByID)
		route.PUT("/apply_fs", mr.minerController.ApplyFs)
		route.PUT("/transfer", mr.minerController.Transfer)
		route.GET("/rig_conf", mr.minerController.GetRigConf)
	}
}
